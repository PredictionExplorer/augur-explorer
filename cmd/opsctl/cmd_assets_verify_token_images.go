package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/spf13/cobra"
)

const (
	// rwalkNFTAddr is the RandomWalk NFT contract whose minted-token count
	// bounds the scan.
	rwalkNFTAddr = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"
	// rwalkTmpImageFile is the temp file each fetched image is written to.
	rwalkTmpImageFile = "randomwalk_tmp.png"
	// rwalkImagesURL is the base URL of the RandomWalk image server.
	rwalkImagesURL = "https://api.randomwalknft.com:1443/images/randomwalk"
)

// newAssetsVerifyTokenImagesCmd builds `opsctl assets verify-token-images`,
// the replacement for the notibot verify-token-imgs script.
func newAssetsVerifyTokenImagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-token-images",
		Short: "Check every RandomWalk token image URL for HTTP 403 responses",
		Long: `Fetches the image of every minted RandomWalk token from the public image
server and reports tokens for which the server answers HTTP 403 (a known
RandomWalk webserver bug for early token ids).

The database connection is built from the PGSQL_* environment variables.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVerifyTokenImages()
		},
	}
	return cmd
}

func init() { assetsCmd.AddCommand(newAssetsVerifyTokenImagesCmd()) }

func runVerifyTokenImages() error {
	info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	st, err := store.New(context.Background(), store.ConfigFromEnv())
	if err != nil {
		return fmt.Errorf("failed to connect to storage: %w", err)
	}
	defer st.Close()
	storagew := &rwstore.SQLStorageWrapper{S: store.NewSQLStorageFromDB(st.DB(), info)}
	rwalkAid, err := storagew.S.Lookup_address_id(rwalkNFTAddr)
	if err != nil {
		return fmt.Errorf("can't resolve RandomWalk contract address id: %w", err)
	}
	rwStats := storagew.Get_random_walk_stats(rwalkAid)
	numTokens := rwStats.TokensMinted
	info.Printf("num_tokens = %v\n", numTokens)

	failed := int64(0)
	for i := int64(0); i < numTokens; i++ {
		if webReturns403(info, i) {
			info.Printf("token %v: FAIL\n", i)
			failed++
		}
	}
	info.Printf("Process ended, failed tokens: %v\n", failed)
	return nil
}

// webReturns403 reports whether the image server answers HTTP 403 for the
// token. This detects a known bug at the RandomWalk webserver, which returns
// 403 for token ids 1 to 269.
func webReturns403(info *log.Logger, tokenID int64) bool {
	url := tokenImageURL(tokenID)
	status, _ := fetchImage(info, url)
	if status == http.StatusForbidden {
		info.Printf("Image server returns 403 code for token %v...\n", tokenID)
		return true
	}
	return false
}

// tokenImageURL returns the black-background image URL for a token.
func tokenImageURL(tokenID int64) string {
	return fmt.Sprintf("%v/%06d_black.png", rwalkImagesURL, tokenID)
}

// tmpImageFilename is the scratch path the fetched image is stored at.
func tmpImageFilename() string {
	return fmt.Sprintf("%v/%v", os.TempDir(), rwalkTmpImageFile)
}

// fetchImage downloads url into the temp image file and returns the HTTP
// status code.
func fetchImage(info *log.Logger, url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		info.Printf("Can't fetch image %v : %v\n", url, err)
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errStr := fmt.Sprintf("HTTP response was not 'Ok' : %v (url=%v)\n", response.StatusCode, url)
		info.Printf("%v\n", errStr)
		return response.StatusCode, errors.New(errStr)
	}

	imgFileName := tmpImageFilename()
	os.Remove(imgFileName)
	file, err := os.Create(imgFileName)
	if err != nil {
		info.Printf("Can't create temporal image file %v : %v\n", imgFileName, err)
		return 0, err
	}
	defer file.Close()
	if _, err = io.Copy(file, response.Body); err != nil {
		info.Printf("Can't copy image data to tmp file: %v\n", err)
		return 0, err
	}
	return response.StatusCode, nil
}
