package tweets
import (
	//"os"
	"io"
	"fmt"
	"errors"
	"strconv"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
)
const (
	URL			string = "https://api.twitter.com/1.1/statuses/update.json"
	MEDIA_URL   string = "https://upload.twitter.com/1.1/media/upload.json"
)
type TwitterKeys struct {
	ApiKey          string
	ApiSecret       string
	TokenKey        string
	TokenSecret     string
}
type MediaImage struct {
	Image_type      string      `json:"image_type"`
	W               int64       `json:"w"`
	H               int64       `json:"h"`
}
type ImageResponse struct {
	Media_id            int64       `json:"media_id"`
	Media_id_string     string      `json:"media_id_string"`
	Media_key           string      `json:"media_key"`
	Size                int64       `json:"size"`
	Expires_after_secs  int64       `json:"expires_after_secs"`
	Image               MediaImage  `json:"image"`
}
type StatusUpdateResponse struct {
	Id					int64		`json:"id"`
	IdStr				string		`json:"id_str"`
}
func decode_response(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
func format_nonce(nonce_int uint64) string {
	return strconv.FormatUint(nonce_int, 16)
}
func SendTweet(api_key,api_secret,access_token,token_secret,message string,session_nonce uint64) (int,string,error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
/*
	session_nonce,err := strconv.ParseUint(session_nonce_hex,16,64)
	if err != nil {
		err_str := fmt.Sprintf("Parsing nonce error: %v\n",err)
		return 0,"",errors.New(err_str)
	}
*/
	var client Client
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token
	client.Nonce=format_nonce(session_nonce)

	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	form := url.Values{"status": {message}}
	resp, err := client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		err_str := fmt.Sprintf("Post error: %v\n",err)
		return 0, "",errors.New(err_str)
	}
	defer resp.Body.Close()

	var err_str string
	var err_out error
	if resp.StatusCode != 200 {
		err_str = fmt.Sprintf("Error at Twitter occured: Status = %v\n",resp.Status)
		err_out = errors.New(err_str)
	}
	b, err := io.ReadAll(resp.Body)
	body_str := string(b)
	return resp.StatusCode,body_str,err_out
}
func SendTweetWithImage(api_key,api_secret,access_token,token_secret,message string,session_nonce uint64,image_data []byte,reply_tweet string) (int,string,error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
/*
	session_nonce,err := strconv.ParseUint(session_nonce_hex,16,64)
	if err != nil {
		err_str := fmt.Sprintf("Parsing nonce error: %v\n",err)
		return 0,"",errors.New(err_str)
	}
*/
	var client Client
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token
	client.Nonce=format_nonce(session_nonce)

	encoded_image_data := base64.StdEncoding.EncodeToString([]byte(image_data))
	form := url.Values{
		"media_data": {encoded_image_data},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	resp, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("Error occured: Status = %v\n",resp.Status)
	} else {
		fmt.Printf("Successfuly sent\n")
	}
	var image_response ImageResponse
	err = decode_response(resp, &image_response)
	if err != nil {
		fmt.Printf("Decode error: %v\n",err)
	}
	fmt.Printf("Media id= %v\n",image_response.Media_id_string)
	fmt.Printf("Image Response: %+v\n",image_response)

	form = url.Values {
		"status": {message},
		"media_ids": {image_response.Media_id_string},
		"in_reply_to_status_id" : {reply_tweet},
	}
	resp, err = client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		err_str := fmt.Sprintf("Post error: %v\n",err)
		return 0, "",errors.New(err_str)
	}
	defer resp.Body.Close()

	var err_str string
	var err_out error
	if resp.StatusCode != 200 {
		err_str = fmt.Sprintf("Error at Twitter occured: Status = %v\n",resp.Status)
		err_out = errors.New(err_str)
	}
	b, err := io.ReadAll(resp.Body)
	body_str := string(b)

	return resp.StatusCode,body_str,err_out
}
