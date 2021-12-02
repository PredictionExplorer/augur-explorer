package main


import (
	"fmt"
	"time"
	"os"
	"io"
	"math/big"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os/signal"
	"syscall"
	"net/http"
	"log"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/tweets"
	contracts "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
)
var (
	TWITTER_KEYS_FILE = os.Getenv("TWITTER_KEYS_FILE")

	//Sample URL: https://randomwalknft.s3.us-east-2.amazonaws.com/003246_black.png
	IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	TMP_IMAGE_FILE		string = "randomwalk_tmp.png"
	DETAIL_URL			string = "https://randomwalknft.com/detail"
	MAX_TIMEOUT_COUNTER		int = 1000
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	market_order_id int64 = 0
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage
	rw_contracts RW_ContractAddresses
	market_addr common.Address
	rwalk_addr common.Address
	rwalk_ctrct *contracts.RWalk
	eclient *ethclient.Client
	rpcclient *rpc.Client
	twkeys TwitterKeys
	twitter_nonce	uint64 = uint64(time.Now().UnixNano())
	cur_floor_price			float64
	rwalk_ctrct_aid			int64
	market_ctrct_aid			int64
)
type TwitterKeys struct {
	ApiKey			string
	ApiSecret		string
	TokenKey		string
	TokenSecret		string
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
			Info.Printf("get_image_file...: aborted by timeout at %v iterations\n",time_out_counter)
			return false
		}
	}
	return false
}
func get_withdrawal_amount() (float64,bool) {

	time_out_counter := int(0)
	for {
		var copts bind.CallOpts
		amount,err := rwalk_ctrct.WithdrawalAmount(&copts)
		if err != nil {
			Info.Printf("Error getting withdrawal amount: %v\n",err)
			Error.Printf("Error getting withdrawal amount: %v\n",err)
			time.Sleep(1 * time.Second)
		} else {
			f := big.NewFloat(0.0)
			f.SetInt(amount)
			div := big.NewFloat(1e+18)
			f=f.Quo(f,div)
			output,_ := f.Float64()
			return output,true
		}
		time_out_counter++
		if time_out_counter > MAX_TIMEOUT_COUNTER {
			Info.Printf("get_withdrawal_amount(): aborted by timeout at %v iterations\n",time_out_counter)
			return 0.0,false
		}
	}
	return 0.0,false
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
func check_floor_price_change_and_emit() {

	no_offers,db_floor_price,offer_id,token_id,err := storage.Get_floor_price(rwalk_ctrct_aid,market_ctrct_aid)
	if no_offers {
		return
	}
	if err != nil {
		Error.Printf("Can't get floor price: %v\n",err)
		Error.Printf("Can't get floor price: %v\n",err)
		return
	}
	if db_floor_price == cur_floor_price {
		return
	}
	cur_floor_price = db_floor_price

	var success bool
	success = get_image_file_from_net_until_success(token_id)
	if !success {
		Error.Printf("Couldn't get image file in check_floor_price(), aborting.")
		return
	}
	image_filename := tmp_img_filename()
	image_data, err := os.ReadFile(image_filename)
	if err != nil {
		fmt.Printf("Can't read image at %v : %v\n",image_filename)
		os.Exit(1)
	}
	var tweet_msg string
	tweet_msg = fmt.Sprintf(
		"Floor price changed to %.4fΞ. Offer %v for token %v.\n\n%v",
		cur_floor_price,
		offer_id,
		token_id,
		fmt.Sprintf("%v/%v",DETAIL_URL,token_id),
	)
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
}
func monitor_events(exit_chan chan bool,addr common.Address) {

	rwalk_aid := storage.Lookup_address_id(addr.String())
	ts := storage.Get_server_timestamp()
	//ts = ts-12*60*60 /// for testing only
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		check_floor_price_change_and_emit()
		records := storage.Get_all_events_for_notification(rwalk_aid,ts)
		for i:=0; i<len(records); i++ {
			select {
				case exit_flag := <-exit_chan:
					if exit_flag {
						Info.Println("Exiting by user request.")
						os.Exit(0)
					}
				default:
			}

			rec := &records[i]
			var withdrawal_amount float64
			var success bool
			if rec.EvtType == 1 {
				withdrawal_amount,success = get_withdrawal_amount()
				if !success {
					Error.Printf("Couldn't get withdrawal amount, aborting.")
					break;
				}
			}
			success = get_image_file_from_net_until_success(rec.TokenId)
			if !success {
				Error.Printf("Couldn't get image file for token %v, aborting.",rec.TokenId)
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
						"#%v Minted for %.4fΞ.\nLast minter would get %.2fΞ if there is no other mint for 30 days.\n\n%v",
						rec.TokenId,
						rec.Price,
						withdrawal_amount,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
				case 2:
					tweet_msg = fmt.Sprintf(
						"#%v On sale for %.4fΞ\n\n%v",
						rec.TokenId,
						rec.Price,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
				case 3:
					tweet_msg = fmt.Sprintf(
						"#%v Bought for %.4fΞ\n\n%v",
						rec.TokenId,
						rec.Price,
						fmt.Sprintf("%v/%v",DETAIL_URL,rec.TokenId),
					)
			}
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
		if len(records) == 0 {
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

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

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


	rwalk_ctrct,err = contracts.NewRWalk(rwalk_addr,eclient)
	if err != nil {
		Info.Printf("Can't instantiate RandomWalk contract %v : %v\n",rwalk_addr.String(),err)
		Error.Printf("Can't instantiate RandomWalk contract %v : %v\n",rwalk_addr.String(),err)
		os.Exit(1)
	}
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
	rwalk_ctrct_aid=storage.Lookup_address_id(rwalk_addr.String())
	market_ctrct_aid=storage.Lookup_address_id(market_addr.String())
	_,cur_floor_price,_,_,err = storage.Get_floor_price(rwalk_ctrct_aid,market_ctrct_aid)
	cur_floor_price = 0.0;
	monitor_events(exit_chan,rwalk_addr)

}
