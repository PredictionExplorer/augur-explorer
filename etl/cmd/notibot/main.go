package main
// Permission required for Discord Bot to update statistics channels:
//		Manage Channel
//		Connect
// Permissions required for other users to avoid joining statistical channels:
//		View Channel			Yes
//		Manage channel			No
//		Connect					No

import (
	"fmt"
	"time"
	"sync"
	"os"
	"io"
	"regexp"
	"flag"
	"strconv"
	"math/big"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os/signal"
	"bytes"
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

	"github.com/andersfylling/disgord"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)
const (
	NumMintsUnicodeChar			string = "🪙 "
	LastPriceUnicodeChar		string = "💲"
	EthSign						string = "Ξ"
	DEFAULT_LAST_MINTED_INTERVAL	time.Duration = 60*9	//Discord resource limit time is about 6 min
)
var (
	MintChannelID_Uint			uint64 = 0 //918642461314785290
	MintChannelID				= disgord.Snowflake(MintChannelID_Uint)
	PriceChannelID_Uint			uint64 = 0 //918643820734869525
	PriceChannelID				= disgord.Snowflake(PriceChannelID_Uint)
	LastDateChannelID_Uint		uint64 = 0 // 918653298813313044
	LastDateChannelID			= disgord.Snowflake(LastDateChannelID_Uint)
	LastRewardChannelID_Uint	uint64 = 0
	LastRewardChannelID			= disgord.Snowflake(LastRewardChannelID_Uint)


	TWITTER_KEYS_FILE = os.Getenv("TWITTER_KEYS_FILE")
	DISCORD_KEYS_FILE = os.Getenv("DISCORD_KEYS_FILE")

	//Sample URL: https://randomwalknft.s3.us-east-2.amazonaws.com/003246_black.png
	IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	VIDEOS_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	TMP_IMG_DATA_FILE	string = "randomwalk_img.png"	// could be an image or video
	TMP_VIDEO_DATA_FILE	string = "randomwalk_video.mp4"
	RESAMPLED_TMP_FILE	string = "randomwalk_resampled.mp4"	// could be an image or video
	DETAIL_URL			string = "https://randomwalknft.com/detail"
	MAX_TIMEOUT_COUNTER		int = 1000
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	DEV_MODE				bool = false // true for testing new features

	Error					*log.Logger
	Info					*log.Logger
	storage					*SQLStorage
	rw_contracts			RW_ContractAddresses
	market_addr				common.Address
	rwalk_addr common.Address
	rwalk_ctrct *contracts.RWalk
	eclient *ethclient.Client
	rpcclient *rpc.Client
	twitter_keys TwitterKeys
	twitter_nonce	uint64 = uint64(time.Now().UnixNano())
	cur_floor_price			float64
	rwalk_ctrct_aid			int64
	market_ctrct_aid			int64
	discord_keys			DiscordKeys
	disc_client				*disgord.Client
	flag_twitter			*bool
	flag_discord			*bool
	last_mint_ts			int64 = 0
)
type DiscordKeys struct {
	TokenKey				string
	ChannelId				uint64	// Notifications Channel
	MainChannelId			uint64	// Main chat Channel
	MintStatsChanId			uint64
	PriceStatsChanId		uint64
	DateStatsChanId			uint64
	RewardStatsChanId		uint64
}
func read_twitter_keys() error {
	file_name := fmt.Sprintf("%v/configs/%v",os.Getenv("HOME"),TWITTER_KEYS_FILE)
	b, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Can't read configuration file with Twitter account keys in %v: %v\n",file_name,err)
		os.Exit(1)
	}
	return json.Unmarshal(b, &twitter_keys)
}
func read_discord_keys() error {
	file_name := fmt.Sprintf("%v/configs/%v",os.Getenv("HOME"),DISCORD_KEYS_FILE)
	b, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Can't read configuration file with Discord account keys in %v: %v\n",file_name,err)
		os.Exit(1)
	}
	err = json.Unmarshal(b, &discord_keys)
	if err != nil {
		return err
	}
	Info.Printf("Main channel ID=%v\n",discord_keys.MainChannelId)

	MintChannelID_Uint = discord_keys.MintStatsChanId
	Info.Printf("Channel ID for Mint statistics: %v\n",MintChannelID_Uint)
	MintChannelID = disgord.Snowflake(MintChannelID_Uint)

	PriceChannelID_Uint	= discord_keys.PriceStatsChanId 
	Info.Printf("Channel ID for Price statistics: %v\n",PriceChannelID_Uint)
	PriceChannelID = disgord.Snowflake(PriceChannelID_Uint)

	LastDateChannelID_Uint = discord_keys.DateStatsChanId
	Info.Printf("Channel ID for Date statistics: %v\n",LastDateChannelID_Uint)
	LastDateChannelID = disgord.Snowflake(LastDateChannelID_Uint)

	LastRewardChannelID_Uint = discord_keys.RewardStatsChanId
	Info.Printf("Channel ID for Last Reward statistics: %v\n",LastRewardChannelID_Uint)
	LastRewardChannelID = disgord.Snowflake(LastRewardChannelID_Uint)

	return err
}
func decode_response(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	} else {
	//	fmt.Printf("Body:\n")
	//	io.Copy(os.Stdout, resp.Body);
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
func tmp_img_filename() string {
	return fmt.Sprintf("%v/%v",os.TempDir(),TMP_IMG_DATA_FILE)
}
func tmp_video_filename() string {
	return fmt.Sprintf("%v/%v",os.TempDir(),TMP_VIDEO_DATA_FILE)
}
func tmp_video_filename_resampled() string {
	return fmt.Sprintf("%v/%v",os.TempDir(),RESAMPLED_TMP_FILE)
}
func fetch_remote_file(url string,dst_file_name string) (int,error) {

	response, err := http.Get(url)
	if err != nil {
		Error.Printf("Can't fetch image %v : %v\n",url,err)
		Info.Printf("Can't fetch image %v : %v\n",url,err)
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		os.Remove(dst_file_name)
		file, err := os.Create(dst_file_name)
		if err != nil {
			Error.Printf("Can't create temporal image file %v : %v\n",dst_file_name,err)
			Info.Printf("Can't create temporal image file %v : %v\n",dst_file_name,err)
			return 0, err
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			Error.Printf("Can't copy file data to tmp file: %v\n",err)
			Info.Printf("Can't copy file data to tmp file: %v\n",err)
			return 0, err
		}
		return response.StatusCode,nil
	} else {
		err_str := fmt.Sprintf("HTTP response was not 'Ok' : %v\n",response.StatusCode)
		Error.Printf("%v\n",err_str)
		Info.Printf("%v\n",err_str)
		return response.StatusCode,errors.New(err_str)
	}
}
func fmt_url_addr_for_image(token_id int64) string {

	url := fmt.Sprintf("%v/%06d_black.png",IMAGES_URL,token_id)
	return url
}
func fmt_url_addr_for_video(token_id int64) string {
	//	https://randomwalknft.s3.us-east-2.amazonaws.com/003968_black_single.mp4
	url := fmt.Sprintf("%v/%06d_black_single.mp4",VIDEOS_URL,token_id)
	return url
}
func web_returns_403_code(token_id int64,url string,dst_file_name string) bool {
	//this is needed to recover from 403 code which the image server returns for 
	//token ids from 1 to 269 (a bug at the Rwalk webserver)

	status,_:= fetch_remote_file(url,dst_file_name)
	if status == 403 {
		Info.Printf("Image server returns 403 code for token %v...\n",token_id)
		return true
	}
	return false
}
func get_file_from_net_until_success(token_id int64,url string,dst_file string) bool {

	time_out_counter := int(0)
	Info.Printf("Fetching file for token %v: %v\n",token_id,url)
	Info.Printf("Destination file: %v\n",dst_file)
	for {
		status,err := fetch_remote_file(url,dst_file)
		if status == 404 {	// image wasn't generated yet
			Info.Printf("File for token %v is not yet ready (%v status), waiting...\n",token_id,status)
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
			Info.Printf("get_file_from_net...: aborted by timeout at %v iterations\n",time_out_counter)
			return false
		}
	}
	return false
}
func get_image(token_id int64) ([]byte,bool) {

	img_file_name := tmp_img_filename()
	success := get_file_from_net_until_success(token_id,fmt_url_addr_for_image(token_id),img_file_name)
	if !success {
		return nil,false
	}
	image_filename := tmp_img_filename()
	image_data, err := os.ReadFile(image_filename)
	if err != nil {
		fmt.Printf("Can't read image at %v : %v\n",image_filename)
		os.Exit(1)
	}
	return image_data,true
}
func get_video(token_id int64) ([]byte,bool) {

	in_video_file_name := tmp_video_filename()
	success := get_file_from_net_until_success(token_id,fmt_url_addr_for_video(token_id),in_video_file_name)
	if !success {
		return nil,false
	}
	// now we have to convert the file to 640x480 because Twitter doesn't accept big resolution
	out_filename := tmp_video_filename_resampled()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := ffmpeg.Input(in_video_file_name).
			Output(out_filename, ffmpeg.KwArgs{"s": "640x480"}).
			OverWriteOutput().ErrorToStdOut().Run()
		if err != nil {
			err_str := fmt.Sprintf("Can't convert the video file with ffmpeg tools: %v\n",err)
			Info.Printf("%v",err_str)
			Error.Printf("%v",err_str)
			Info.Printf("Exiting due to unrecoverable error\n")
			os.Exit(1)
		}
		Info.Printf("Video resampling finished\n")
		syscall.Sync()
		wg.Done()
	}()
	wg.Wait()
	video_data, err := os.ReadFile(out_filename)
	if err != nil {
		err_str := fmt.Sprintf("Can't read video at %v : %v\n",out_filename)
		Info.Printf("%v\n",err_str)
		Error.Printf("%v\n",err_str)
		os.Exit(1)
	}
	return video_data,true
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
			Info.Printf("got withdrawal amount %v (rounded %v)\n",amount.String(),output)
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
func format_url(token_id int64) string {
	return fmt.Sprintf("\n\n%v/%v",DETAIL_URL,token_id)
}
func format_notification_message(event_type int64,token_id int64,price,withdrawal_amount float64,include_url bool) string {
	var output string
	var url string
	if include_url {
		url = format_url(token_id)
	}
	switch event_type {
		case 1:
			output = fmt.Sprintf(
				"#%v Minted for %.4fΞ.\nLast minter would get %.2fΞ if there is no other mint for 30 days.%v",
				token_id,
				price,
				withdrawal_amount,
				url,
			)
		case 2:
			output = fmt.Sprintf(
				"#%v On sale for %.4fΞ.%v",
				token_id,
				price,
				url,
			)
		case 3:
			output = fmt.Sprintf(
				"#%v Bought for %.4fΞ.%v",
				token_id,
				price,
				url,
			)
		case 4:
			output = fmt.Sprintf(
				"Floor price changed to %.4fΞ.%v",
				cur_floor_price,
				url,
			)
		case 5:
			output = fmt.Sprintf(
				"Buy offer for %.4fΞ.%v",
				price,
				url,
			)
	}
	return output
}
func set_channel_name(new_name string,channel_id disgord.Snowflake) {
	if !*flag_discord { return }
	_,err := disc_client.Channel(channel_id).UpdateBuilder().SetName(new_name).Execute()
	if err != nil {
		Info.Printf("Couldn't change channel name to %v (channel id = %v) : %v\n",new_name,channel_id,err)
		Error.Printf("Couldn't change channel name to %v (channel id = %v) : %v\n",new_name,channel_id,err)
		pattern := "retry_after\":\\s+(\\d+)\\.\\d"
		re := regexp.MustCompile(pattern)
		matchall := re.FindStringSubmatch(err.Error())
		if len(matchall)==2 {
			delay_sec,err := strconv.ParseInt(matchall[1],10,64)
			if err != nil {
				Error.Printf("Unable to parse delay value (%v) : %v\n",matchall[1],err)
				Info.Printf("Unable to parse delay value (%v) : %v\n",matchall[1],err)
				return
			}
			Info.Printf("Retrying channel name update in %v sec\n",delay_sec)
			time.Sleep(time.Duration(delay_sec) * time.Second)
			time.Sleep(1 * time.Second) // just extra for safety
			_,err = disc_client.Channel(channel_id).UpdateBuilder().SetName(new_name).Execute()
			if err != nil {
				Info.Printf("Couldn't change channel (second time) name to %v (channel id = %v) : %v\n",new_name,channel_id,err)
				Error.Printf("Couldn't change channel (second time) name to %v (channel id = %v) : %v\n",new_name,channel_id,err)
			}
		} else {
			Info.Printf("Retry skipped, len of matchall = %v\n",len(matchall))
		}
	}
}
func notify_twitter_one_message(token_id int64,evt_type int64,msg string,image_data []byte) {
	// only one message is sent, with image
	Info.Printf("notify_twitter(token_id=%v)\n",token_id)

	twitter_nonce++
	status_code,body,err := SendTweetWithImage(
		twitter_keys.ApiKey,
		twitter_keys.ApiSecret,
		twitter_keys.TokenKey,
		twitter_keys.TokenSecret,
		msg,
		twitter_nonce,
		image_data,
		"",
	)
	if err != nil {
		Info.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
		Error.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
	} else {
		Info.Printf(
			"Notification for evt type=%v of token id=%v to Twitter successful\n",
			evt_type,token_id,
		)
	}
}
func notify_twitter_two_messages(token_id int64,evt_type int64,msg string,image_data []byte,video_data []byte) {
	// first the message with image is sent, after that second message with video is posted as reply to first
	Info.Printf("notify_twitter_two_messages(token_id=%v)\n",token_id)

	twitter_nonce++
	status_code,body,err := SendTweetWithImage(
		twitter_keys.ApiKey,
		twitter_keys.ApiSecret,
		twitter_keys.TokenKey,
		twitter_keys.TokenSecret,
		msg,
		twitter_nonce,
		image_data,
		"",
	)
	if err != nil {
		Info.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
		Error.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
	} else {
		if status_code != 200 {
			Info.Printf("Notification of FIRST msg produced error: %v\n",err)
			Error.Printf("Notification of FIRST msg produced error: %v\n",err)
		} else {
			Info.Printf(
				"Notification of FIRST msg for evt type=%v of token id=%v to Twitter successful\n",
				evt_type,token_id,
			)
			var st_upd StatusUpdateResponse
			err := json.NewDecoder(bytes.NewReader([]byte(body))).Decode(&st_upd)
			if err != nil {
				Info.Printf("Could not decode response of FIRST msg, err: %v\n",err)
				Error.Printf("Could not decode response of FIRST msg, err: %v\n",err)
			} else {
				// everything ok, now send second msg
				msg_id_str := st_upd.IdStr
				msg := fmt.Sprintf("Video for token %v",token_id)
				twitter_nonce++
				status_code2,_,err:=SendTweetWithVideo(
					twitter_keys.ApiKey,
					twitter_keys.ApiSecret,
					twitter_keys.TokenKey,
					twitter_keys.TokenSecret,
					msg,
					twitter_nonce,
					video_data,
					msg_id_str,
				)
				if err != nil {
					Info.Printf("Error sending second message with video: %v\n",err)
					Error.Printf("Error sending second message with video: %v\n",err)
					time.Sleep(30*time.Second)	// sleep 30 sec to avoid spamming Twitter in case of error
				}
				Info.Printf("Status = %v for HTTP response of second tweet\n",status_code2)
			}
		}
	}
}
func notify_discord(token_id int64,msg string,image_data []byte,image_url string) error {

	if !*flag_discord {
		return nil
	}
	Info.Printf("notify_discord(token_id=%v)\n",token_id)
	rdr := bytes.NewReader(image_data)
	var err error
	// this is the Notification channel
	_, err = disc_client.Channel(disgord.Snowflake(discord_keys.ChannelId)).CreateMessage(
			&disgord.CreateMessageParams{
				Content: msg,
				Files: []disgord.CreateMessageFileParams{
					{rdr, "token.png", false},
				},
				Embed: &disgord.Embed{
					Description: image_url,
					URL: image_url,
				},
			},
	)
	if err != nil {
		return err
	}

	return err
}
func check_floor_price_change_and_emit() {

	no_offers,db_floor_price,_,token_id,err := storage.Get_floor_price(rwalk_ctrct_aid,market_ctrct_aid)
	if no_offers {
		return
	}
	if err != nil {
		Error.Printf("Can't get floor price: %v\n",err)
		Info.Printf("Can't get floor price: %v\n",err)
		return
	}
	if db_floor_price == cur_floor_price {
		return
	}
	cur_floor_price = db_floor_price

	var success bool
	var image_data []byte

	image_data,success = get_image(token_id)
	if !success {
		Error.Printf("Couldn't get image file in check_floor_price(), aborting.")
		return
	}
	Info.Printf("Floor price change detected (new price=%v)\n",cur_floor_price)
	if *flag_twitter {
		notif_msg := format_notification_message(4,token_id,cur_floor_price,0.0,true)
		twitter_nonce++
		status_code,body,err := SendTweetWithImage(
			twitter_keys.ApiKey,
			twitter_keys.ApiSecret,
			twitter_keys.TokenKey,
			twitter_keys.TokenSecret,
			notif_msg,
			twitter_nonce,
			image_data,
			"",
		)
		if err != nil {
			Info.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
		} else {
			Info.Printf("Notified about floor price change to Twitter (new price = %v)\n",cur_floor_price)
		}
	}
	if *flag_discord {
		notif_msg := format_notification_message(4,token_id,cur_floor_price,0.0,false)
		msg_url := format_url(token_id)
		err := notify_discord(token_id,notif_msg,image_data,msg_url)
		if err != nil {
			Info.Printf("Error sending msg to Discord: %v\n",err)
		} else {
			Info.Printf("Notified about floor price change to Discord (new price = %v)\n",cur_floor_price)
		}
	}
}
func update_last_minted_time() {
	// go-routine, updates last timed time ervery X amount of time
	for {
		if *flag_discord {
			if last_mint_ts > 0 {
				cur_time := time.Now()
				minted_time := time.Unix(last_mint_ts,0)
				duration := DurationToString(TimeDifference(minted_time,cur_time))
				new_channel_name := fmt.Sprintf(
					"Last mint: %v ago",
					duration,
				)
				set_channel_name(new_channel_name,LastDateChannelID)
				Info.Printf("Set last mint time to : %v",new_channel_name)
			}
		}
		time.Sleep(DEFAULT_LAST_MINTED_INTERVAL*time.Second)
	}
}
func update_last_reward(wamount float64) {

	if *flag_discord {
		new_channel_name := fmt.Sprintf("Last reward: %.2f%v",wamount,EthSign)
		set_channel_name(new_channel_name,LastRewardChannelID)
	}
}
func monitor_events(exit_chan chan bool,addr common.Address) {

	// notification types:
	//		1		Mint
	//		2		New Offer Sell			(otype=1)
	//		3		Item Bought
	//		4		Flour Price Changed
	//		5		New Offer Buy			(otype=0)
	rwalk_aid := storage.Lookup_address_id(addr.String())
	msg_status := storage.Get_messaging_status()
	cur_evtlog_id := msg_status.EvtLogId
	cur_ts := msg_status.TimeStamp
	Info.Printf(
		"monitor_events() starts with evtlog_id=%v (timestamp %v) (%v)\n",
		cur_evtlog_id,cur_ts,time.Unix(cur_ts,0).Format("2006-01-02T15:04:05"),
	)
	if DEV_MODE {
		// these are hardcoded values (for testing) obtained from the DB
		cur_ts = 1676508454 -1
		cur_evtlog_id=400625175 - 1
		Info.Printf("DEVELOPMENT mode on, starting from ts=%v(%v)\n",cur_ts,time.Unix(cur_ts,0).Format("2006-01-02T15:04:05"))
	}
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
		records := storage.Get_all_events_for_notification2(rwalk_aid,cur_evtlog_id)
		if len(records) > 0 {
			Info.Printf(
				"Got %v records for timestamp %v (%v) (evtlog_id=%v)\n",
				len(records),cur_ts,time.Unix(cur_ts,0).Format("2006-01-02T15:04:05"),cur_evtlog_id,
			)
		}
		os.Exit(1)
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
			Info.Printf(
				"Processing evt type=%v of token id=%v to Twitter (price= %v)\n",
				rec.EvtType,rec.TokenId,rec.Price,
			)
			var withdrawal_amount float64
			var success bool
			if rec.EvtType == 1 { // Mint
				withdrawal_amount,success = get_withdrawal_amount()
				if !success {
					Error.Printf("Couldn't get withdrawal amount, aborting.")
					break;
				}
				new_channel_name := fmt.Sprintf(
					"Cur. price %v : %.4f",
					LastPriceUnicodeChar,
					rec.Price,
				)
				set_channel_name(new_channel_name,PriceChannelID)
				new_channel_name = fmt.Sprintf(
					"Num. mints %v : %d",
					NumMintsUnicodeChar,
					rec.TokenId+1,
				)
				set_channel_name(new_channel_name,MintChannelID)
				last_mint_ts = rec.TimeStampMinted
				update_last_reward(withdrawal_amount)
			}
			is_403_code := web_returns_403_code(rec.TokenId,fmt_url_addr_for_image(rec.TokenId),tmp_img_filename())
			if is_403_code {
				Info.Printf("Skipping event for token %v due to 403 error\n",rec.TokenId)
				time.Sleep(1 * time.Second)
				break // 403 is not good for a perpetual fetch via HTTP, so we skip the event
			}
			image_data,success := get_image(rec.TokenId)
			var video_data []byte
			if rec.EvtType == 1 { // Mint
				// we get the video file only for Mint type events
				video_data,success = get_video(rec.TokenId)
				if !success {
					Error.Printf("Couldn't get VIDEO file for token %v, aborting.",rec.TokenId)
					time.Sleep(10 * time.Second)
					break
				}
			}
			cur_evtlog_id=rec.EvtLogId
			cur_ts = rec.TimeStampMinted
			Info.Printf(
				"Setting evtlog_id to %v, timestamp to %v (%v) (token_id=%v, evt_type=%v)\n",
				cur_evtlog_id,cur_ts,time.Unix(cur_ts,0).Format("2006-01-02T15:04:05"),rec.TokenId,rec.EvtType,
			)

			if *flag_twitter {
				notif_msg := format_notification_message(rec.EvtType,rec.TokenId,rec.Price,withdrawal_amount,true)
				if rec.EvtType == 1 { // Mint
					notify_twitter_two_messages(rec.TokenId,rec.EvtType,notif_msg,image_data,video_data)
				} else {
					notify_twitter_one_message(rec.TokenId,rec.EvtType,notif_msg,image_data)
				}
			}
			if *flag_discord {
				notif_msg := format_notification_message(rec.EvtType,rec.TokenId,rec.Price,withdrawal_amount,false)
				msg_url := format_url(rec.TokenId)
				err := notify_discord(rec.TokenId,notif_msg,image_data,msg_url)
				if err != nil {
					Error.Printf("Error on Discord notification: %v\n",err)
				} else {
					Info.Printf("Notification of event (token_id=%v) to Discord successful\n",rec.TokenId)
				}
			}
			msg_status.EvtLogId=cur_evtlog_id
			msg_status.TimeStamp=cur_ts
			storage.Update_messaging_status(&msg_status)
		}
		if DEV_MODE {
			Info.Printf("cur ts=%v(%v), records=%v\n",cur_ts,time.Unix(cur_ts,0).Format("2006-01-02T15:04:05"),len(records))
		}
		if len(records) == 0 {
			time.Sleep(30 * time.Second) // sleep only if there is no data
		}
	}
}
func main() {

	flag_twitter = flag.Bool("twitter", false, "Send messages to Twitter")
	flag_discord = flag.Bool("discord", false, "Send messages to Discord")
	flag.Parse()
	if !(*flag_twitter || *flag_discord) {
		fmt.Printf("Please use --twitter or --discord flag\n")
		os.Exit(1)
	}
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/notibot_db.log",log_dir)

	fname:=fmt.Sprintf("%v/notibot_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/notibot_error.log",log_dir)
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

	storage = Connect_to_storage(Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	rw_contracts = storage.Get_randomwalk_contract_addresses()
	rwalk_addr = common.HexToAddress(rw_contracts.RandomWalk)
	market_addr = common.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n",rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n",market_addr.String())

	if *flag_twitter {
		err = read_twitter_keys()
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		Info.Printf("Loaded Twitter keys\n")
	}
	if *flag_discord {
		err = read_discord_keys()
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		disc_client = disgord.New(
			disgord.Config{
				BotToken: discord_keys.TokenKey,
			},
		)
		Info.Printf("Loaded discord keys\n")
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
	rwalk_ctrct_aid=storage.Lookup_address_id(rwalk_addr.String())
	market_ctrct_aid=storage.Lookup_address_id(market_addr.String())
	_,cur_floor_price,_,_,err = storage.Get_floor_price(rwalk_ctrct_aid,market_ctrct_aid)
	//cur_floor_price = 0.0;
	last_mint_ts = storage.Get_last_mint_timestamp()
	go update_last_minted_time()

	wamount,success := get_withdrawal_amount()
	if success {
		update_last_reward(wamount)
	}
	monitor_events(exit_chan,rwalk_addr)

}
