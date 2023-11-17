package main

import (
	"net/http"
	"os"
	"fmt"
	"time"
	"log"
	"encoding/json"
	"bytes"
	"flag"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/cosmicgame"
)
const (
// old values, discontinued
//	REQUEST_URL	string = "https://randomwalknft-api.com/cosmicgame_tokens"
//	IMAGE_URL  string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
//	VIDEO_URL string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
)
var (
	REQUEST_URL	string = os.Getenv("IM_REQUEST_URL")
	IMAGE_URL  string = os.Getenv("IM_IMAGE_URL")
	VIDEO_URL string = os.Getenv("IM_VIDEO_URL")
	storagew                SQLStorageWrapper
	 Info                    *log.Logger
)
func generate_token_artifacts(token_id int64,seed string) {

	values := map[string]interface{}{"token_id": token_id, "seed": seed};
	json_data, err := json.Marshal(values)
	if err != nil {
		fmt.Printf("Error parsing json: %v\n",err)
	}
	resp, err := http.Post(REQUEST_URL, "application/json",bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Printf("Error submitting POST request: %v\n",err)
		return
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}
func token_artifacts_exist(token_id int64) (bool,error) {

	img_file := fmt.Sprintf("%v%06d.png",IMAGE_URL,token_id)
	video_file := fmt.Sprintf("%v%06d.mp4",VIDEO_URL,token_id)
	res,err := http.Head(img_file)
	if err != nil {
		fmt.Printf("Error getting img file: %v\n",err)
		return false,err
	}
	if res.StatusCode !=200 {
		return false,nil
	}
	res,err =  http.Head(video_file)
	if err != nil {
		fmt.Printf("Error getting video file: %v\n",err)
		return false,err
	}
	if res.StatusCode !=200 {
		fmt.Printf("Video file not found: %+v\n",res)
		return false,nil
	}
	return true,nil
}
func main() {


	regenerate_missing := flag.Bool("regenerate", false, "Regenerate missing images")
	flag.Parse()

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage(Info)
    storagew.S.Db_set_schema_name("public");
    storagew.S.Log_msg("Log initialized\n")

	if *regenerate_missing {
		fmt.Printf("Regenerating missing images/videos\n")
	} else {
		fmt.Printf("Checking image/video presence\n")
	}
	tokens := storagew.Get_cosmic_signature_nft_list(0, 100000)

	for i:=int64(0);i<int64(len(tokens));i++ {
		tok := tokens[i]
		fmt.Printf("token id = %v    ",tok.TokenId)
		exists,err := token_artifacts_exist(tok.TokenId);
		if err == nil {
			if !exists {
				if *regenerate_missing {
					fmt.Printf(" regenerating ...")
					generate_token_artifacts(tok.TokenId,tok.Seed)
					for {
						exists,err := token_artifacts_exist(tok.TokenId);
						if err != nil {
							fmt.Printf("aborting due to error: %v ",err)
							break
						} else {
							if exists { 
								break
							} else {
								fmt.Printf(".")
							}
						}
						time.Sleep(10* time.Second)
					}
					fmt.Printf(" done.\n")
				} else {
					fmt.Printf("doesn't exist\n")
				}
			} else {
				fmt.Printf("image/video present\n")
			}
		} else {
			fmt.Printf("error: %v\n",err)
		}
		time.Sleep(1* time.Second)
	}
}
