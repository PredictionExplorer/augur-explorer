package main

import (
	"net/http"
	//"os"
	"fmt"
	"encoding/json"
	"bytes"

	 . "github.com/PredictionExplorer/augur-explorer/dbs/biddingwar"
)
const (
	REQUEST_URL	string = "https://randomwalknft-api.com/cosmicgame_tokens"
	IMAGE_URL  string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
	VIDEO_URL string = "https://cosmic-game.s3.us-east-2.amazonaws.com/"
)
var (
	storagew                SQLStorageWrapper
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
	fmt.Printf("resp: %+v\n",resp)
	fmt.Printf("res: \n%v\n",res)
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Printf("json:\n%v\n",res["json"])
}
func token_artifacts_exist(token_id string) (bool,error) {

	img_file := fmt.Sprintf("%v%v.png",IMAGE_URL,token_id)
	video_file := fmt.Sprintf("%v%v.mp4",VIDEO_URL,token_id)
	res,err := http.Head(img_file)
	if err != nil {
		fmt.Printf("Error getting img file: %v\n",err)
		return false,err
	}
	if res.StatusCode !=200 {
		fmt.Printf("Image file not found: %+v\n",res)
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
	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(Info)
	tokens := storage
/*
	if len(os.Args) < 3 {
		fmt.Printf("usage: %v [token_id] [seed]\n");
		os.Exit(1)
	}
*/
	starting_token_id := 0
	for {

	}
	//token_artifacts_exist("000072")
	//generate_token_artifacts(67,"4be9157e3676c30594ec2b1d8b08971264d3d7d73ed91aa44e0a492a22e08ba5")
}
