package main
import (
	"fmt"
	"os"
	"io"
	"errors"
	"net/http"
	"log"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"

)
const (
	RWALK_ADDR				string = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"
	TMP_IMAGE_FILE			string = "randomwalk_tmp.png"
	IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
)
var (
	storagew				*rwdb.SQLStorageWrapper
	Error					*log.Logger
	Info					*log.Logger
	rwalk_aid				int64
)

func tmp_img_filename() string {
	return fmt.Sprintf("%v/%v",os.TempDir(),TMP_IMAGE_FILE)
}
func fetch_image(url string) (int,error) {

	response, err := http.Get(url)
	if err != nil {
		Error.Printf("Can't fetch image %v : %v\n",url,err)
		Info.Printf("Can't fetch image %v : %v\n",url,err)
		return 0,err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		img_file_name := tmp_img_filename()
		os.Remove(img_file_name)
		file, err := os.Create(img_file_name)
		if err != nil {
			Info.Printf("Can't create temporal image file %v : %v\n",img_file_name,err)
			return 0,err
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			Info.Printf("Can't copy image data to tmp file: %v\n",err)
			return 0,err
		}
		return response.StatusCode,nil
	} else {
		err_str := fmt.Sprintf("HTTP response was not 'Ok' : %v (url=%v)\n",response.StatusCode,url)
		Info.Printf("%v\n",err_str)
		return response.StatusCode,errors.New(err_str)
	}
}
func fmt_url_addr(token_id int64) string {

	url := fmt.Sprintf("%v/%06d_black.png",IMAGES_URL,token_id)
	return url
}
func web_returns_403_code(token_id int64) bool {
	//this is needed to recover from 403 code which the image server returns for
	//token ids from 1 to 269 (a bug at the Rwalk webserver)

	url := fmt_url_addr(token_id)
	status,_:= fetch_image(url)
	if status == 403 {
		Info.Printf("Image server returns 403 code for token %v...\n",token_id)
		return true
	}
	return false
}
func main() {

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	Error = Info

	storage := dbs.Connect_to_storage(Info)
	storagew = &rwdb.SQLStorageWrapper{S: storage}
	rwalk_aid = storagew.S.Lookup_address_id(RWALK_ADDR)
	rw_stats := storagew.Get_random_walk_stats(rwalk_aid)
	num_tokens := rw_stats.TokensMinted
	Info.Printf("num_tokens = %v\n",num_tokens)

	counter := int64(0)
	for i:=int64(0);i<num_tokens;i++ {
		is_403_fail := web_returns_403_code(i)
		if is_403_fail {
			Info.Printf("token %v: FAIL\n",i)
			counter++
		} else {
			//Info.Printf("token %v: OK",i)
		}
	}
	Info.Printf("Process ended, failed tokens: %v\n",counter)
}
