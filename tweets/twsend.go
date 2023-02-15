package tweets
import (
	//"os"
	"io"
	"fmt"
	"time"
	"errors"
	"bytes"
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
type ProcessingInfo struct {
	State				string		`json:"in_progress"`
	Check_after_secs	int64		`json:"check_after_secs"`
	Progress_percent	int64		`json:"progress_percent"`
}
type VideoUploadStatus struct {
	Media_id			int64		`json:"media_id"`
	Media_id_string		string		`json:"media_id_string"`
	Expires_after_secs  int64       `json:"expires_after_secs"`
	Processing_info		ProcessingInfo	`json:"processing_info"`
}
func decode_response(resp *http.Response, data interface{}) error {
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error at ioutilReadAll(): %v",err))
	}
	// note: we did ReadAll() because we need to have an option to dump the received body
	fmt.Printf("Body: %v\n",string(p))
	body_stream:=bytes.NewReader(p)
	if (resp.StatusCode != 200) && (resp.StatusCode !=202) && (resp.StatusCode != 204) {
		fmt.Printf("statusCode=%v, this is an error\n",resp.StatusCode)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, string(p))
	}
	return json.NewDecoder(body_stream).Decode(data)
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
func SendTweetWithAttachment(api_key,api_secret,access_token,token_secret,message string,session_nonce uint64,reply_tweet_id string,attachment_url string) (int,string,error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	var client Client
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token
	client.Nonce=format_nonce(session_nonce)

	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	form := url.Values{
		"status": {message},
		"in_reply_to_status_id" : {reply_tweet_id},
		"attachment_url" : {attachment_url},
	}
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
func SendTweetWithMedia(api_key,api_secret,access_token,token_secret,message string,session_nonce uint64,media_type string,image_data []byte,reply_tweet string) (int,string,error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	var client Client
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token
	client.Nonce=format_nonce(session_nonce)
	fmt.Printf("using media_type=%v\n",media_type)
	encoded_image_data := base64.StdEncoding.EncodeToString([]byte(image_data))
	form := url.Values{
		"media_data": {encoded_image_data},
		"media_category":{ "tweet_video"},
	//	"media_type": {media_type},
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
func SendTweetWithVideo(api_key,api_secret,access_token,token_secret,message string,session_nonce uint64,video_data []byte,reply_tweet string) (int,string,error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	var client Client
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token
	client.Nonce=format_nonce(session_nonce)
	encoded_video_data := base64.StdEncoding.EncodeToString([]byte(video_data))
	string_data := string(video_data)
	fmt.Printf("length of video=%v, legth of string data=%v\n",len(video_data),len(string_data))
	form := url.Values{
		"command": {"INIT"},
		"total_bytes" : {fmt.Sprintf("%v",len(video_data))},
		"media_category":{ "tweet_video"},
		"media_type": {"video/mp4"},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	//---------- INIT
	resp_init, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	defer resp_init.Body.Close()
	if resp_init.StatusCode != 202 {
		fmt.Printf("Error occured: Status = %v\n",resp_init.Status)
		var img_resp ImageResponse
		err = decode_response(resp_init, &img_resp)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
		}
		b, err := io.ReadAll(resp_init.Body)
		body_str := string(b)
		fmt.Printf("Body of erroneous req: %v\n",body_str)
		return 0, "",err
	} else {
		fmt.Printf("INIT successfuly sent\n")
	}
	var img_response_init ImageResponse
	err = decode_response(resp_init, &img_response_init)
	if err != nil {
		fmt.Printf("Decode error: %v\n",err)
		fmt.Printf("Media id= %v\n",img_response_init.Media_id_string)
		fmt.Printf("Image Response INIT: %+v\n",img_response_init)
		return 0, "",err
	}
	upload_id := img_response_init.Media_id_string
	fmt.Printf("Media id= %v\n",img_response_init.Media_id_string)
	fmt.Printf("Image Response: %+v\n",img_response_init)

	//------------ APPEND
	form = url.Values{
		"command": {"APPEND"},
		"media_id": {upload_id},
		"media_data": {encoded_video_data+fmt.Sprintf("\r\n")},
		//"media": {string_data},
		"segment_index":{fmt.Sprintf("%v",0)},
	}
	_=encoded_video_data
	client.Header = make(map[string][]string)
	hdrval:=client.Header.Get("Content-Type")
	fmt.Printf("Before APPEND content type = %v\n",hdrval)
	client.Header.Set("Content-Type","application/octet-stream")
	client.Header.Set("Content-Transfer-Encoding","base64")
	resp_append, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		err_str := fmt.Sprintf("Post error: %v\n",err)
		return 0, "",errors.New(err_str)
	}
	defer resp_append.Body.Close()
	if (resp_append.StatusCode != 200) && (resp_append.StatusCode != 204) {
		fmt.Printf("Error occured: Status = %v\n",resp_append.Status)
		var img_resp ImageResponse
		err = decode_response(resp_append, &img_resp)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
		}
		b, err := io.ReadAll(resp_append.Body)
		body_str := string(b)
		fmt.Printf("Body of erroneous req: %v\n",body_str)
		return 0, "",err
	} else {
		fmt.Printf("APPEND successfuly sent\n")
	}
	if resp_append.StatusCode != 204 {
		var img_response_append ImageResponse
		err = decode_response(resp_append, &img_response_append)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
			fmt.Printf("Media id= %v\n",img_response_append.Media_id_string)
			fmt.Printf("Image Response APPEND: %+v\n",img_response_append)
			return 0, "",err
		}
	}
	client.Header.Del("Content-Type")
	client.Header.Del("Content-Transfer-Encoding")
	//----------- FINALIZE
	form = url.Values {
		"command": {"FINALIZE"},
		"media_id": {upload_id},
	}
	resp_finalize, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		err_str := fmt.Sprintf("Post error: %v\n",err)
		return 0, "",errors.New(err_str)
	}
	defer resp_finalize.Body.Close()
	if (resp_finalize.StatusCode != 200) && (resp_finalize.StatusCode != 204) {
		fmt.Printf("Error occured: Status = %v\n",resp_finalize.Status)
		var img_resp ImageResponse
		err = decode_response(resp_finalize, &img_resp)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
		}
		b, err := io.ReadAll(resp_finalize.Body)
		body_str := string(b)
		fmt.Printf("Body of erroneous req: %v\n",body_str)
		return 0, "",err
	} else {
		fmt.Printf("FINALZIE successfuly sent (code=%v)\n",resp_finalize.StatusCode)
	}
	if resp_finalize.StatusCode != 204 {
		var img_response_finalize ImageResponse
		err = decode_response(resp_finalize, &img_response_finalize)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
			fmt.Printf("Media id= %v\n",img_response_finalize.Media_id_string)
			fmt.Printf("Image Response FINALIZE: %+v\n",img_response_finalize)
			return 0, "",err
		}
	}

	// Now check if Twitter have processed the video and finished on its end
	form = url.Values {
		"command": {"STATUS"},
		"media_id": {upload_id},
	}
	for {
		resp_media_status, err := client.Get(nil, &token_credentials, MEDIA_URL, form)
		if err != nil {
			err_str := fmt.Sprintf("Post error: %v\n",err)
			return 0, "",errors.New(err_str)
		}
		defer resp_media_status.Body.Close()
		var vid_upload_status VideoUploadStatus
		if resp_finalize.StatusCode != 200 {
			err = decode_response(resp_finalize, &vid_upload_status)
			if err != nil {
				fmt.Printf("Decode error: %v\n",err)
				fmt.Printf("Media id= %v\n",vid_upload_status.Media_id_string)
				fmt.Printf("Video upload status: %+v\n",vid_upload_status)
				return 0, "",err
			}
		}
		err = decode_response(resp_media_status, &vid_upload_status)
		if err != nil {
			fmt.Printf("Decode error: %v\n",err)
			fmt.Printf("Media id= %v\n",vid_upload_status.Media_id_string)
			fmt.Printf("Video upload status: %+v\n",vid_upload_status)
			return 0, "",err
		}
		fmt.Printf("Progress: %v %%\n",vid_upload_status.Processing_info.Progress_percent)
		fmt.Printf("Check after: %v sec\n",vid_upload_status.Processing_info.Check_after_secs)
		if vid_upload_status.Processing_info.Progress_percent == 100 {
			fmt.Printf("Upload progress reached 100%, continuing\n")
			break
		}
		time.Sleep(2*time.Second)
	}
	fmt.Printf("Upload of video successful, sending status update (message)\n")

	// Send STATUS UPDATE request
	form = url.Values {
		"status": {message},
		"media_ids": {upload_id},
		"in_reply_to_status_id" : {reply_tweet},
	}
	resp_update_status, err := client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		err_str := fmt.Sprintf("Post error: %v\n",err)
		return 0, "",errors.New(err_str)
	}
	defer resp_update_status.Body.Close()

	var err_str string
	var err_out error
	if resp_update_status.StatusCode != 200 {
		err_str = fmt.Sprintf("Error at Twitter occured: Status = %v\n",resp_update_status.Status)
		err_out = errors.New(err_str)
		b, err := io.ReadAll(resp_update_status.Body)
		if err !=nil {
			fmt.Printf("Error in ReadAll() for body of STATUS: %v\n",err)
		}
		body_str := string(b)
		fmt.Printf("Body of erroneous req: %v\n",body_str)
		return 0, "",err_out
	}
	b, err := io.ReadAll(resp_update_status.Body)
	body_str := string(b)

	return resp_update_status.StatusCode,body_str,err_out
}
