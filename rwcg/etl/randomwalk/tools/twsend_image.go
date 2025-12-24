package main
import (
	"os"
	"fmt"
	"io"
	"time"
	"encoding/base64"
	"strconv"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"

//	"github.com/dghubble/go-twitter"
//	"github.com/gomodule/oauth1/oauth"
)
const (
	URL			string = "https://api.twitter.com/1.1/statuses/update.json"
	MEDIA_URL	string = "https://upload.twitter.com/1.1/media/upload.json"
)
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
func format_nonce(nonce_int uint64) string {
	return strconv.FormatUint(nonce_int, 16)
}
func main() {

	if len(os.Args) < 7 {
		fmt.Printf(
			"Usage: \n\t\t%v [api_key] [api_secret] [access_token] [token_secret] [nonce] [image]\n\t\t"+
			"Sends a tweet with image attached on behalf of a user\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	api_key := os.Args[1]
	api_secret := os.Args[2]
	access_token := os.Args[3]
	token_secret := os.Args[4]
	session_nonce,err := strconv.ParseUint(os.Args[5],16,64)
	if err != nil {
		fmt.Printf("Parsing nonce error: %v\n",err)
		os.Exit(1)
	}
	session_nonce++

	image_filename := os.Args[6]
	image_data, err := os.ReadFile(image_filename)
	if err != nil {
		fmt.Printf("Can't read image at %v : %v\n",image_filename)
		os.Exit(1)
	}
	var client Client
	var signing_credentials Credentials
	signing_credentials.Token = api_key
	signing_credentials.Secret = api_secret
	client.Credentials.Token = api_key // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey=api_key
	client.ClientToken = access_token

	nonce_counter := uint64(time.Now().UnixNano())
	session_nonce = nonce_counter
	client.Nonce=format_nonce(session_nonce)

	encoded_image_data := base64.StdEncoding.EncodeToString([]byte(image_data))
	form := url.Values{
		"status" : {"upload test"},
		"media_data": {encoded_image_data},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret
	resp, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		fmt.Printf("Post error: %v\n",err)
		os.Exit(1)
	}
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

	//now submit attachment
	form = url.Values{
		"status": {"random walk tokens"},
		"media_ids": {image_response.Media_id_string},
	}
	resp, err = client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		fmt.Printf("Post error: %v\n",err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("Error occured: Status = %v\n",resp.Status)
	} else {
		fmt.Printf("Successfuly sent\n")
	}
	b, err := io.ReadAll(resp.Body)
	var data struct {
		d struct {
			created_at string
			id	string
			text string
		}
	}
	err = decode_response(resp, &data)
	if err != nil {
		fmt.Printf("Decode error: %v\n",err)
	}
	fmt.Printf("Data: %+v\n",data)
	fmt.Printf("body = %+v\n",resp.Body)
	fmt.Printf("dump body:\n")
	fmt.Println(string(b))
}
