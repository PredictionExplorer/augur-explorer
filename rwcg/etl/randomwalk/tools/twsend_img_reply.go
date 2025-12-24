package main
import (
	"os"
	"fmt"
	"time"
	"strings"
	"encoding/json"
	"io/ioutil"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/tweets"
)
const (
)
var (
	TWITTER_KEYS_FILE = os.Getenv("TWITTER_KEYS_FILE")
	twitter_keys TwitterKeys
	twitter_nonce   uint64 = uint64(time.Now().UnixNano())
)
func read_twitter_keys() error {
	file_name := fmt.Sprintf("%v/configs/%v",os.Getenv("HOME"),TWITTER_KEYS_FILE)
	b, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Can't read configuration file with Twitter account keys in %v: %v\n",file_name,err)
		os.Exit(1)
	}
	return json.Unmarshal(b, &twitter_keys)
}

func notify_twitter(token_id int64,msg string,image_data []byte,reply_tweet string) string {

	fmt.Printf("notify_twitter(token_id=%v)\n",token_id)

	twitter_nonce++
	status_code,body,err := SendTweetWithImage(
		twitter_keys.ApiKey,
		twitter_keys.ApiSecret,
		twitter_keys.TokenKey,
		twitter_keys.TokenSecret,
		msg,
		twitter_nonce,
		image_data,
		reply_tweet,
	)
	if err != nil {
		fmt.Printf("Error sending tweet: %v (status %v; body = %v)\n",err,status_code,body)
		os.Exit(1)
	}
	fmt.Printf("body after send: %v\n",body)
	var status_resp StatusUpdateResponse
	err = json.NewDecoder(strings.NewReader(body)).Decode(&status_resp)
	if err != nil {
		fmt.Printf("Error at decode response: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("status_resp.Id=%v\n",status_resp.Id)
	fmt.Printf("status_resp.IdStr=%v\n",status_resp.IdStr)
	return fmt.Sprintf("%v",status_resp.Id)
}
func notify_twitter_media(token_id int64,msg string,media_type string,media_data []byte,reply_id string) string {

	fmt.Printf("notify_twitter(token_id=%v)\n",token_id)

	twitter_nonce++
	status_code,body,err := SendTweetWithVideo(
		twitter_keys.ApiKey,
		twitter_keys.ApiSecret,
		twitter_keys.TokenKey,
		twitter_keys.TokenSecret,
		msg,
		twitter_nonce,
		media_data,
		reply_id,
	)
	if err != nil {
		fmt.Printf("Error sending tweet with media: %v (status %v; body = %v)\n",err,status_code,body)
		os.Exit(1)
	}
	fmt.Printf("body after send: %v\n",body)
	var status_resp StatusUpdateResponse
	err = json.NewDecoder(strings.NewReader(body)).Decode(&status_resp)
	if err != nil {
		fmt.Printf("Error at decode response: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("status_resp.Id=%v\n",status_resp.Id)
	fmt.Printf("status_resp.IdStr=%v\n",status_resp.IdStr)
	return fmt.Sprintf("%v",status_resp.Id)
}
func main() {

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [reply_to_id] [media_file] [message]\n\t\t"+
			"Sends a tweet with media_file (image or video) attached as reply-to\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	err := read_twitter_keys()
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}

	reply_to_id := os.Args[1]
	media_filename := os.Args[2]
	message := os.Args[3]
	fmt.Printf("Media file: %v\n",media_filename)
	fmt.Printf("Reply to id: %v\n",reply_to_id)
	media_data, err := os.ReadFile(media_filename)
	if err != nil {
		fmt.Printf("Can't read media data at %v : %v\n",media_filename)
		os.Exit(1)
	}

	nonce_counter := uint64(time.Now().UnixNano())
	twitter_nonce = nonce_counter

	/*
	msg_id2 := notify_twitter(11,"test message",image_data2,msg_id)
	fmt.Printf("msg id2 = %v\n",msg_id2)
	*/
//	url := "https://randomwalknft.s3.us-east-2.amazonaws.com/003913_black_single.mp4"

	//media_type := "image/png"
	media_type := 	"video/mp4"
	notify_twitter_media(11,message,media_type,media_data,reply_to_id)
}
