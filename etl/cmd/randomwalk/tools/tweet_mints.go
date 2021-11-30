package main


import (
	"fmt"
	"time"
	"os"
	"io"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os/signal"
	"syscall"
	"net/http"
	"log"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/tweets"
)
const (
)
var (
	TWITTER_KEYS_FILE = os.Getenv("TWITTER_KEYS_FILE")

	//Sample URL: https://randomwalknft.s3.us-east-2.amazonaws.com/003246_black.png
	IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	TMP_IMAGE_FILE		string = "randomwalk_tmp.png"
	DETAIL_URL			string = "https://randomwalknft.com/detail/"
	MAX_TIMEOUT_COUNTER		int = 1000

	market_order_id int64 = 0
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage
	rw_contracts RW_ContractAddresses
	market_addr common.Address
	rwalk_addr common.Address
	twkeys TwitterKeys
	twitter_nonce	uint64 = uint64(time.Now().UnixNano())
)
type TwitterKeys struct {
	ApiKey			string
	ApiSecret		string
	TokenKey		string
	TokenSecret		string
}
type MediaImage struct {
	Image_type		string		`json:"image_type"`
	W				int64		`json:"w"`
	H				int64		`json:"h"`
}
type ImageResponse struct {
	Media_id			int64		`json:"media_id"`
	Media_id_string		string		`json:"media_id_string"`
	Media_key			string		`json:"media_key"`
	Size				int64		`json:"size"`
	Expires_after_secs	int64		`json:"expires_after_secs"`
	Image				MediaImage	`json:"image"`
}
func read_twitter_keys() error {
	file_name := fmt.Sprintf("%v/configs/%v",os.Getenv("HOME"),TWITTER_KEYS_FILE)
	b, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Can't read configuration file with twitter account keys in %v: %v\n",file_name,err)
		os.Exit(1)
	}
	return json.Unmarshal(b, &twkeys)
}
func decode_response(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	} else {
		fmt.Printf("Body:\n")
	//	io.Copy(os.Stdout, resp.Body);
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
func tmp_img_filename() string {
	return fmt.Sprintf("%v/%v",os.TempDir(),TMP_IMAGE_FILE)
}
func fetch_image(url string) (int,error) {

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		Error.Printf("Can't fetch image %v : %v\n",url,err)
		Info.Printf("Can't fetch image %v : %v\n",url,err)
		return 0,err
	}
	if response.StatusCode == 200 {
		img_file_name := tmp_img_filename()
		os.Remove(img_file_name)
		file, err := os.Create(img_file_name)
		if err != nil {
			Error.Printf("Can't create temporal image file %v : %v\n",img_file_name,err)
			Info.Printf("Can't create temporal image file %v : %v\n",img_file_name,err)
			return 0,err
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			Error.Printf("Can't copy image data to tmp file: %v\n",err)
			Info.Printf("Can't copy image data to tmp file: %v\n",err)
			return 0,err
		}
		return response.StatusCode,nil
	} else {
		err_str := fmt.Sprintf("HTTP response was not 'Ok' : %v\n",response.StatusCode)
		Error.Printf("%v\n",err_str)
		Info.Printf("%v\n",err_str)
		return response.StatusCode,errors.New(err_str)
	}
}
func get_image_file_from_net_until_success(token_id int64) bool {

	time_out_counter := int(0)
	url := fmt.Sprintf("%v/%06d_black.png",IMAGES_URL,token_id)
	Info.Printf("Fetching image for token %v: %v\n",token_id,url)
	for {
		status,err := fetch_image(url)
		if status == 404 {	// image wasn't generated yet
			Info.Printf("Image for token %v is not yet ready (%v status), waiting...\n",token_id,status)
			time.Sleep(60 * time.Second)
		} else {
			if err != nil {
				Info.Printf("Aborting due to errors\n")
				return false
			} else {
				return true
			}
		}
		time_out_counter++
		if time_out_counter > MAX_TIMEOUT_COUNTER {
			Info.Printf("Aborted by timeout at %v iterations\n",time_out_counter)
			return false
		}
	}
	return false
}
func notify_twitter(rec *RW_NotificationEvent) {

	url_twitter := fmt.Sprintf(
		"http://%v",
		rec.TokenId,
		rec.Price,
	)
	_,err := http.Get(url_twitter)
	if err!= nil {
		Error.Printf("Error accesing Twitter :%v  (url = %v)\n",err,url_twitter)
	}
}
func monitor_events(exit_chan chan bool,addr common.Address) {

	rwalk_aid := storage.Lookup_address_id(addr.String())
	ts := storage.Get_server_timestamp()
	ts = ts-5*24*60*60 /// for testing only
	for {
		Info.Printf("Timestamp for query is %v\n",time.Unix(ts,0))
		records := storage.Get_all_events_for_notification(rwalk_aid,ts)
		for i:=0; i<len(records); i++ {
			rec := &records[i]
			success := get_image_file_from_net_until_success(rec.TokenId)
			if !success {
				Error.Printf("Couldn't get image file, aborting.")
				time.Sleep(10 * time.Second)
				break
			}
			image_filename := tmp_img_filename()
			image_data, err := os.ReadFile(image_filename)
			if err != nil {
				fmt.Printf("Can't read image at %v : %v\n",image_filename)
				os.Exit(1)
			}
			ts = rec.TimeStampMinted
			var tweet_msg string
			switch rec.EvtType {
				case 1:
					tweet_msg = fmt.Sprintf(
						"Token %v minted. Price %.4fΞ  %v ",
						rec.TokenId,
						rec.Price,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
				case 2:
					tweet_msg = fmt.Sprintf(
						"New offer for token %v . Price %.4fΞ ",
						rec.TokenId,
						rec.Price,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
				case 3:
					tweet_msg = fmt.Sprintf(
						"Token %v bought. Price %.4fΞ ",
						rec.TokenId,
						rec.Price,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
			}
			Info.Printf("evt type = %v, msg=%v, ts=%v\n",rec.EvtType,tweet_msg,ts)
			twitter_nonce++
			status_code,body,err := SendTweetWithImage(
				twkeys.ApiKey,
				twkeys.ApiSecret,
				twkeys.TokenKey,
				twkeys.TokenSecret,
				tweet_msg,
				twitter_nonce,
				image_data,
			)
			if err != nil {
				Info.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
			}

			Info.Printf("Notified mint of token id=%v to Twitter (price= %v)\n",rec.TokenId,rec.Price)
		}
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		Info.Printf("Loop ended\n")
		if len(records) == 0 {
			Info.Printf("Sleeping 5 sec\n")
			time.Sleep(5 * time.Second) // sleep only if there is no data
		}
	}
}

func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/tweet_notifs_db.log",log_dir)

	fname:=fmt.Sprintf("%v/tweet_notifs_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/tweet_notifs_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)


	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	rw_contracts = storage.Get_randomwalk_contract_addresses()
	rwalk_addr = common.HexToAddress(rw_contracts.RandomWalk)
	market_addr = common.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n",rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n",market_addr.String())

	err = read_twitter_keys()
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()
/*
			tweet_msg := fmt.Sprintf(
					"Token %v minted. Price %v. Seed %v.",
					1,1.45,"adfadsfas",
			)
			status_code,body,err := SendTweet(
				twkeys.ApiKey,
				twkeys.ApiSecret,
				twkeys.TokenKey,
				twkeys.TokenSecret,
				tweet_msg,
				twitter_nonce,
			)
			if err != nil {
				Info.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
			}
			os.Exit(1)
*/
	monitor_events(exit_chan,rwalk_addr)

}
