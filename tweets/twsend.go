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
)
const (
	URL			string = "https://api.twitter.com/1.1/statuses/update.json"
)
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
		return 0,"",errors.New(err_str)
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
