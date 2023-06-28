package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
	"bytes"
)
const (
	REQUEST_URL	string = "https://randomwalknft-api.com/cosmicgame_tokens"
	IMAGE_URL  string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
	VIDEO_URL string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
)
var (
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
	//fmt.Printf("imgfile = %v\nvideofile = %v\n",img_file,video_file)
	res,err := http.Head(img_file)
	if err != nil {
		fmt.Printf("Error getting img file: %v\n",err)
		return false,err
	}
	if res.StatusCode !=200 {
	//	fmt.Printf("Image file not found: %+v\n",res)
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

	if len(os.Args) < 3 {
		fmt.Printf("usage: %v [token_id] [seed]\n")
		os.Exit(1)
	}
	token_id_str := os.Args[1]
	seed := os.Args[2]
	token_id,err := strconv.ParseInt(token_id_str,10,64)
    if err != nil {
        fmt.Printf("error parsing tokenid: %v\n",err)
        os.Exit(1)
    }
	generate_token_artifacts(token_id,seed)
}
