// Copied from: github.com/febriliankr/whatsapp-cloud-api
// (had to copy because the library has a bug (issue #2 at Github) wich I fixed in my copy)
package wanotif
import (
	"fmt"
	"bytes"
	"io"
	"errors"
	"net/http"
	"encoding/json"
)
type Whatsapp struct {
	Token         string
	APIVersion    string
	PhoneNumberID string
	Language      TemplateLanguage
}
type TemplateLanguage struct {
	Code string `json:"code,omitempty"`
}
var (
	LanguageEnglish = TemplateLanguage{
		Code: "en",
	}
)
func NewWhatsapp(token string, phoneNumberID string) *Whatsapp {
	return &Whatsapp{
		Language:      LanguageEnglish,
		Token:         token,
		APIVersion:    "v14.0",
		PhoneNumberID: phoneNumberID,
	}
}
func parseHTTPError(body io.Reader) (err error) {
	var errRes map[string]map[string]interface{}
	err = json.NewDecoder(body).Decode(&errRes)
	if err != nil {
		return fmt.Errorf("unparsed error message")
	}
	msg := fmt.Sprintf("%s", errRes["error"]["message"])
	return errors.New(msg)
}
func (wa *Whatsapp) sendMessage(request interface{}) (res map[string]interface{}, err error) {

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(jsonRequest)

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", wa.APIVersion, wa.PhoneNumberID)
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Authorization", "Bearer "+wa.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	fmt.Printf("Do(): resp=%v, err=%v\n",resp,err)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := parseHTTPError(resp.Body)
		return res, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	fmt.Printf("body bytes=%+v\n",string(bodyBytes))
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		return res,err
	}
	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)
	if err != nil {
		return res, err
	}
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
func (wa *Whatsapp) SendText(toPhoneNumber string, text string) (res map[string]interface{}, err error) {

	request := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                toPhoneNumber,
		"type":              "text",
		"text": map[string]string{
			"body": string(text),
		},
	}
	return wa.sendMessage(request)

}
/*
	Sample use
	note: you must initiate conversation with the sending account first, otherwise
			the message will be classified as spam and won't be delivered, though no err is reported
func main() {

	my_phone_id := os.Getenv("WHATSAPP_PHONE_ID")
	token := os.Getenv("WHATSAPP_TOKEN")
	wa := NewWhatsapp(token,my_phone_id)
	wa.APIVersion = "v15.0"
	res, err := wa.SendText("phone_num", "your_message")
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("Message sent successfuly\n")
	fmt.Printf("res = %v\n",res)
	os.Exit(0)
}
*/
