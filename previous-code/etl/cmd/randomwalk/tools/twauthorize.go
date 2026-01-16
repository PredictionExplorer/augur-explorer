package main
import (
	"os"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"

//	"github.com/dghubble/go-twitter"
	"github.com/gomodule/oauth1/oauth"
)
const (
	//URL				string = "http://api.twitter.com/2/tweets"
	URL			string = "https://api.twitter.com/1.1/statuses/update.json"
	//URL			string = "https://api.twitter.com/oauth/authorize"
)
func decode_response(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
func main() {

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [access_token] [token_secret] [nonce]\n\t\t"+
			"Requests a twitter user to authorize the access token to send twits on his behalf\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	access_token := os.Args[1]
	token_secret := os.Args[2]
	nonce := os.ARgs[3]

	var client oauth.Client
	client.Credentials.Token = access_token // app key
	client.Credentials.Secret = token_secret // app secret

	form := url.Values{"status": {"hello"}}
	resp, err := client.Post(nil, &client.Credentials, URL, form)
	if err != nil {
		fmt.Printf("Post error: %v\n",err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	var data struct {
		d struct {
			id	string
			text string
		}
	}
	err = decode_response(resp, &data)
	if err != nil {
		fmt.Printf("Decode error: %v\n",err)
	}
	fmt.Printf("Data: %+v\n",data)
}
