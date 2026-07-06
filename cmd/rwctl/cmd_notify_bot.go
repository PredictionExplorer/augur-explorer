package main

// The notify-bot and tweet-mints subcommands share this implementation: the
// legacy notif_bot.go and tweet_mints.go tools were identical copies of the
// same monitor that announces RandomWalk mint/offer/purchase events and floor
// price changes on Twitter.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
	rwprim "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Notification-bot constants preserved from the legacy tool.
const (
	// notifyImagesURL is the base URL of the generated token images.
	notifyImagesURL = "https://api.randomwalknft.com:1443/images/randomwalk"
	// notifyTmpImageFile is the temp-file name the fetched image is stored under.
	notifyTmpImageFile = "randomwalk_tmp.png"
	// notifyDetailURL is the base URL of the token detail page linked in tweets.
	notifyDetailURL = "https://randomwalknft.com/detail"
	// notifyMaxTimeoutCounter bounds the retry loops for images and RPC reads.
	notifyMaxTimeoutCounter = 1000
)

// readTwitterKeys loads the Twitter API credentials from
// $HOME/configs/$TWITTER_KEYS_FILE.
func readTwitterKeys() (tweets.TwitterKeys, error) {
	var keys tweets.TwitterKeys
	fileName := fmt.Sprintf("%v/configs/%v", os.Getenv("HOME"), os.Getenv("TWITTER_KEYS_FILE"))
	b, err := os.ReadFile(fileName)
	if err != nil {
		return keys, fmt.Errorf("can't read configuration file with twitter account keys in %v: %w", fileName, err)
	}
	if err := json.Unmarshal(b, &keys); err != nil {
		return keys, fmt.Errorf("can't parse twitter account keys in %v: %w", fileName, err)
	}
	return keys, nil
}

// mintNotifier holds the state of the running notification bot.
type mintNotifier struct {
	info          *log.Logger
	errlog        *log.Logger
	storagew      *rwstore.SQLStorageWrapper
	rwalkCtrct    *rwcontracts.RWalk
	keys          tweets.TwitterKeys
	nonce         uint64
	curFloorPrice float64
	rwalkAid      int64
	marketAid     int64
}

// tmpImgFilename returns the temp-file path the fetched token image is saved to.
func (n *mintNotifier) tmpImgFilename() string {
	return fmt.Sprintf("%v/%v", os.TempDir(), notifyTmpImageFile)
}

// fetchImage downloads url into the temp image file and returns the HTTP
// status code.
func (n *mintNotifier) fetchImage(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		n.errlog.Printf("Can't fetch image %v : %v\n", url, err)
		n.info.Printf("Can't fetch image %v : %v\n", url, err)
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		errStr := fmt.Sprintf("HTTP response was not 'Ok' : %v\n", response.StatusCode)
		n.errlog.Printf("%v\n", errStr)
		n.info.Printf("%v\n", errStr)
		return response.StatusCode, errors.New(errStr)
	}
	imgFileName := n.tmpImgFilename()
	os.Remove(imgFileName)
	file, err := os.Create(imgFileName)
	if err != nil {
		n.errlog.Printf("Can't create temporal image file %v : %v\n", imgFileName, err)
		n.info.Printf("Can't create temporal image file %v : %v\n", imgFileName, err)
		return 0, err
	}
	defer file.Close()
	if _, err = io.Copy(file, response.Body); err != nil {
		n.errlog.Printf("Can't copy image data to tmp file: %v\n", err)
		n.info.Printf("Can't copy image data to tmp file: %v\n", err)
		return 0, err
	}
	return response.StatusCode, nil
}

// getImageFileUntilSuccess keeps fetching a token's image until it has been
// generated (a 404 means "not ready yet"), bounded by notifyMaxTimeoutCounter.
func (n *mintNotifier) getImageFileUntilSuccess(tokenID int64) bool {
	timeOutCounter := 0
	url := fmt.Sprintf("%v/%06d_black.png", notifyImagesURL, tokenID)
	n.info.Printf("Fetching image for token %v: %v\n", tokenID, url)
	for {
		status, err := n.fetchImage(url)
		if status == 404 { // image wasn't generated yet
			n.info.Printf("Image for token %v is not yet ready (%v status), waiting...\n", tokenID, status)
			time.Sleep(60 * time.Second)
		} else {
			if err != nil {
				n.info.Printf("Aborting due to errors\n")
				return false
			}
			return true
		}
		timeOutCounter++
		if timeOutCounter > notifyMaxTimeoutCounter {
			n.info.Printf("get_image_file...: aborted by timeout at %v iterations\n", timeOutCounter)
			return false
		}
	}
}

// getWithdrawalAmount reads the contract's withdrawal amount (in ETH),
// retrying on RPC errors, bounded by notifyMaxTimeoutCounter.
func (n *mintNotifier) getWithdrawalAmount() (float64, bool) {
	timeOutCounter := 0
	for {
		var copts bind.CallOpts
		amount, err := n.rwalkCtrct.WithdrawalAmount(&copts)
		if err != nil {
			n.info.Printf("Error getting withdrawal amount: %v\n", err)
			n.errlog.Printf("Error getting withdrawal amount: %v\n", err)
			time.Sleep(1 * time.Second)
		} else {
			f := new(big.Float).SetInt(amount)
			f = f.Quo(f, big.NewFloat(1e+18))
			output, _ := f.Float64()
			return output, true
		}
		timeOutCounter++
		if timeOutCounter > notifyMaxTimeoutCounter {
			n.info.Printf("get_withdrawal_amount(): aborted by timeout at %v iterations\n", timeOutCounter)
			return 0.0, false
		}
	}
}

// sendTweetWithImage posts a tweet with the current temp image attached.
func (n *mintNotifier) sendTweetWithImage(msg string, imageData []byte) {
	n.nonce++
	statusCode, body, err := tweets.SendTweetWithImage(
		n.keys.ApiKey,
		n.keys.ApiSecret,
		n.keys.TokenKey,
		n.keys.TokenSecret,
		msg,
		n.nonce,
		imageData,
		"",
	)
	if err != nil {
		n.info.Printf("Error sending tweet: %v (status %v; body = %v)\n", err, statusCode, body)
	}
}

// checkFloorPriceChangeAndEmit tweets when the marketplace floor price moved.
func (n *mintNotifier) checkFloorPriceChangeAndEmit() {
	noOffers, dbFloorPrice, _, tokenID, err := n.storagew.Get_floor_price(n.rwalkAid, n.marketAid)
	if noOffers {
		return
	}
	if err != nil {
		n.errlog.Printf("Can't get floor price: %v\n", err)
		return
	}
	if dbFloorPrice == n.curFloorPrice {
		return
	}
	n.curFloorPrice = dbFloorPrice

	if !n.getImageFileUntilSuccess(tokenID) {
		n.errlog.Printf("Couldn't get image file in check_floor_price(), aborting.")
		return
	}
	imageData, err := os.ReadFile(n.tmpImgFilename())
	if err != nil {
		n.errlog.Printf("Can't read image at %v : %v\n", n.tmpImgFilename(), err)
		return
	}
	tweetMsg := fmt.Sprintf(
		"Floor price changed to %.4fΞ.\n\n%v",
		n.curFloorPrice,
		fmt.Sprintf("%v/%v", notifyDetailURL, tokenID),
	)
	n.sendTweetWithImage(tweetMsg, imageData)
}

// tweetMessageForEvent formats the tweet text for a notification record.
func tweetMessageForEvent(rec *rwprim.NotificationEvent, withdrawalAmount float64) string {
	detail := fmt.Sprintf("%v/%v", notifyDetailURL, rec.TokenId)
	switch rec.EvtType {
	case 1:
		return fmt.Sprintf(
			"#%v Minted for %.4fΞ.\nLast minter would get %.2fΞ if there is no other mint for 30 days.\n\n%v",
			rec.TokenId, rec.Price, withdrawalAmount, detail,
		)
	case 2:
		return fmt.Sprintf("#%v On sale for %.4fΞ\n\n%v", rec.TokenId, rec.Price, detail)
	case 3:
		return fmt.Sprintf("#%v Bought for %.4fΞ\n\n%v", rec.TokenId, rec.Price, detail)
	}
	return ""
}

// monitorEvents polls the database for new notification events and tweets
// them until an exit is requested.
func (n *mintNotifier) monitorEvents(exitChan chan bool, addr common.Address) error {
	rwalkAid, err := n.storagew.S.Lookup_address_id(addr.String())
	if err != nil {
		return fmt.Errorf("can't resolve RandomWalk contract address id: %w", err)
	}
	ts := n.storagew.Get_server_timestamp()
	for {
		select {
		case exitFlag := <-exitChan:
			if exitFlag {
				n.info.Println("Exiting by user request.")
				return nil
			}
		default:
		}
		n.checkFloorPriceChangeAndEmit()
		records := n.storagew.Get_all_events_for_notification(rwalkAid, ts)
		for i := 0; i < len(records); i++ {
			select {
			case exitFlag := <-exitChan:
				if exitFlag {
					n.info.Println("Exiting by user request.")
					return nil
				}
			default:
			}

			rec := &records[i]
			var withdrawalAmount float64
			var success bool
			if rec.EvtType == 1 {
				withdrawalAmount, success = n.getWithdrawalAmount()
				if !success {
					n.errlog.Printf("Couldn't get withdrawal amount, aborting.")
					break
				}
			}
			if !n.getImageFileUntilSuccess(rec.TokenId) {
				n.errlog.Printf("Couldn't get image file for token %v, aborting.", rec.TokenId)
				time.Sleep(10 * time.Second)
				break
			}
			imageData, err := os.ReadFile(n.tmpImgFilename())
			if err != nil {
				return fmt.Errorf("can't read image at %v: %w", n.tmpImgFilename(), err)
			}
			ts = rec.TimeStampMinted
			n.sendTweetWithImage(tweetMessageForEvent(rec, withdrawalAmount), imageData)
			n.info.Printf("Notified mint of token id=%v to Twitter (price= %v)\n", rec.TokenId, rec.Price)
		}
		if len(records) == 0 {
			time.Sleep(5 * time.Second) // sleep only if there is no data
		}
	}
}

// runNotifyBot wires logs, database, RPC and Twitter credentials together and
// runs the notification monitor until interrupted.
func runNotifyBot(cmd *cobra.Command, args []string) error {
	logDir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), primitives.DEFAULT_LOG_DIR)
	os.MkdirAll(logDir, os.ModePerm)
	dbLogFile := fmt.Sprintf("%v/tweet_notifs_db.log", logDir)

	fname := fmt.Sprintf("%v/tweet_notifs_info.log", logDir)
	infoFile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("can't start: %w", err)
	}
	info := log.New(infoFile, "INFO: ", log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/tweet_notifs_error.log", logDir)
	errFile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("can't start: %w", err)
	}
	errlog := log.New(errFile, "ERROR: ", log.Ltime|log.Lshortfile)

	rpcURL := os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("can't connect to ETH node at %v: %w", rpcURL, err)
	}
	info.Printf("Connected to ETH node: %v\n", rpcURL)

	storagew, err := connectRWStorage(info)
	if err != nil {
		return err
	}
	if err := storagew.S.Init_log(dbLogFile); err != nil {
		return fmt.Errorf("can't initialize DB log: %w", err)
	}
	storagew.S.Log_msg("Log initialized\n")

	rwContracts := storagew.Get_randomwalk_contract_addresses()
	rwalkAddr := common.HexToAddress(rwContracts.RandomWalk)
	marketAddr := common.HexToAddress(rwContracts.MarketPlace)
	info.Printf("RandomWalk contract %v\n", rwalkAddr.String())
	info.Printf("MarketPlace contract %v\n", marketAddr.String())

	keys, err := readTwitterKeys()
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	exitChan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		info.Printf("Got SIGINT signal, will exit after processing is over." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exitChan <- true
	}()

	rwalkCtrct, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
	if err != nil {
		info.Printf("Can't instantiate RandomWalk contract %v : %v\n", rwalkAddr.String(), err)
		errlog.Printf("Can't instantiate RandomWalk contract %v : %v\n", rwalkAddr.String(), err)
		return fmt.Errorf("can't instantiate RandomWalk contract %v: %w", rwalkAddr.String(), err)
	}

	rwalkAid, err := storagew.S.Lookup_address_id(rwalkAddr.String())
	if err != nil {
		return fmt.Errorf("can't resolve RandomWalk contract address id: %w", err)
	}
	marketAid, err := storagew.S.Lookup_address_id(marketAddr.String())
	if err != nil {
		return fmt.Errorf("can't resolve MarketPlace contract address id: %w", err)
	}

	n := &mintNotifier{
		info:          info,
		errlog:        errlog,
		storagew:      storagew,
		rwalkCtrct:    rwalkCtrct,
		keys:          keys,
		nonce:         uint64(time.Now().UnixNano()),
		curFloorPrice: 0.0,
		rwalkAid:      rwalkAid,
		marketAid:     marketAid,
	}
	return n.monitorEvents(exitChan, rwalkAddr)
}

// notifyBotEnvHelp documents the environment variables of the notification bot.
const notifyBotEnvHelp = "Environment variables:\n" +
	"  AUGUR_ETH_NODE_RPC_URL  Ethereum RPC endpoint (required)\n" +
	"  TWITTER_KEYS_FILE       name of the JSON credentials file under $HOME/configs (required)\n" +
	"  PGSQL_*                 PostgreSQL connection (required)"

// newNotifyBotCmd builds the notify-bot subcommand (legacy notif_bot tool).
func newNotifyBotCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "notify-bot",
		Short: "Monitor RandomWalk events and announce them on Twitter",
		Long: "Monitors new mint/offer/purchase events and floor price changes and announces them to the Twitter account.\n\n" +
			notifyBotEnvHelp,
		Args: cobra.NoArgs,
		RunE: runNotifyBot,
	}
}

func init() { register(newNotifyBotCmd()) }
