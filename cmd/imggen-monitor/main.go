// Command imggen-monitor verifies that every minted Cosmic Signature NFT has
// its generated image and video artifacts available on the artifact server,
// and optionally asks the generator service to (re)create missing ones.
//
// Modes:
//
//	imggen-monitor                          # report presence of artifacts for all tokens
//	imggen-monitor -regenerate              # regenerate any missing artifacts
//	imggen-monitor -token 123 -seed 0xabc   # one-shot generation for a single token
//
// Configuration (environment):
//
//	IM_REQUEST_URL  POST endpoint of the artifact generator service
//	IM_IMAGE_URL    base URL where images are served  (<base><id>.png)
//	IM_VIDEO_URL    base URL where videos are served  (<base><id>.mp4)
//	PGSQL_*         PostgreSQL connection (scan mode only)
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

type artifactClient struct {
	requestURL string
	imageURL   string
	videoURL   string
}

func newArtifactClientFromEnv() (*artifactClient, error) {
	c := &artifactClient{
		requestURL: os.Getenv("IM_REQUEST_URL"),
		imageURL:   os.Getenv("IM_IMAGE_URL"),
		videoURL:   os.Getenv("IM_VIDEO_URL"),
	}
	if c.requestURL == "" || c.imageURL == "" || c.videoURL == "" {
		return nil, fmt.Errorf("IM_REQUEST_URL, IM_IMAGE_URL and IM_VIDEO_URL must all be set")
	}
	return c, nil
}

// generate asks the generator service to create image/video artifacts for a token.
func (c *artifactClient) generate(tokenID int64, seed string) error {
	payload, err := json.Marshal(map[string]interface{}{"token_id": tokenID, "seed": seed})
	if err != nil {
		return fmt.Errorf("encoding request: %w", err)
	}
	resp, err := http.Post(c.requestURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("submitting generation request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path
	if resp.StatusCode >= 300 {
		return fmt.Errorf("generator service returned %s", resp.Status)
	}
	return nil
}

// exists reports whether both the image and the video artifact are reachable.
func (c *artifactClient) exists(tokenID int64) (bool, error) {
	for _, url := range []string{
		fmt.Sprintf("%v%06d.png", c.imageURL, tokenID),
		fmt.Sprintf("%v%06d.mp4", c.videoURL, tokenID),
	} {
		resp, err := http.Head(url) //nolint:gosec // G107: URL bases come from operator config (IM_IMAGE_URL/IM_VIDEO_URL)
		if err != nil {
			return false, fmt.Errorf("HEAD %s: %w", url, err)
		}
		_ = resp.Body.Close() // HEAD responses carry no body
		if resp.StatusCode != http.StatusOK {
			return false, nil
		}
	}
	return true, nil
}

// waitUntilPresent polls until the token's artifacts appear or an error occurs.
func (c *artifactClient) waitUntilPresent(tokenID int64) error {
	for {
		exists, err := c.exists(tokenID)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}
		fmt.Print(".")
		time.Sleep(10 * time.Second)
	}
}

func scanAll(client *artifactClient, regenerate bool) error {
	info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	st, err := store.New(context.Background(), store.ConfigFromEnv())
	if err != nil {
		return fmt.Errorf("failed to connect to storage: %w", err)
	}
	defer st.Close()
	var storagew cgstore.SQLStorageWrapper
	storagew.S = store.NewSQLStorageFromDB(st.DB(), info)
	storagew.S.Db_set_schema_name("public")

	if regenerate {
		fmt.Println("Regenerating missing images/videos")
	} else {
		fmt.Println("Checking image/video presence")
	}

	tokens := storagew.Get_cosmic_signature_nft_list(0, 100000)
	for _, tok := range tokens {
		fmt.Printf("token id = %v    ", tok.TokenId)
		exists, err := client.exists(tok.TokenId)
		switch {
		case err != nil:
			fmt.Printf("error: %v\n", err)
		case exists:
			fmt.Println("image/video present")
		case !regenerate:
			fmt.Println("doesn't exist")
		default:
			fmt.Print(" regenerating ...")
			if err := client.generate(tok.TokenId, tok.Seed); err != nil {
				fmt.Printf(" request failed: %v\n", err)
				continue
			}
			if err := client.waitUntilPresent(tok.TokenId); err != nil {
				fmt.Printf(" aborting due to error: %v\n", err)
				continue
			}
			fmt.Println(" done.")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	regenerate := flag.Bool("regenerate", false, "regenerate missing artifacts while scanning")
	tokenID := flag.Int64("token", -1, "one-shot: generate artifacts for this token id (requires -seed)")
	seed := flag.String("seed", "", "one-shot: seed of the token passed with -token")
	flag.Parse()

	client, err := newArtifactClientFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "imggen-monitor: %v\n", err)
		os.Exit(1)
	}

	if *tokenID >= 0 {
		if *seed == "" {
			fmt.Fprintln(os.Stderr, "imggen-monitor: -token requires -seed")
			os.Exit(1)
		}
		if err := client.generate(*tokenID, *seed); err != nil {
			fmt.Fprintf(os.Stderr, "imggen-monitor: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := scanAll(client, *regenerate); err != nil {
		fmt.Fprintf(os.Stderr, "imggen-monitor: %v\n", err)
		os.Exit(1)
	}
}
