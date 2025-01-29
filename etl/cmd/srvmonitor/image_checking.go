package main

import (
	"fmt"
	"os"
	"errors"
	"net/http"
)
const (
	RWALK_IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	RWALK_VIDEOS_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	CGAME_IMAGES_URL			string = "https://cosmic-game2.s3.us-east-2.amazonaws.com"
	CGAME_VIDEOS_URL			string = "https://cosmic-game2.s3.us-east-2.amazonaws.com"
)
func fmt_url_addr_for_image_randomwalk(token_id int64) string {

    url := fmt.Sprintf("%v/%06d_black.png",RWALK_IMAGES_URL,token_id)
    return url
}
func fmt_url_addr_for_video_randomwalk(token_id int64) string {
    url := fmt.Sprintf("%v/%06d_black_single.mp4",RWALK_VIDEOS_URL,token_id)
    return url
}
func fmt_url_addr_for_image_cosmicgame(seed string) string {

    url := fmt.Sprintf("%v/0x%s.png",CGAME_IMAGES_URL,seed)
    return url
}
func fmt_url_addr_for_video_cosmicgame(seed string) string {
    url := fmt.Sprintf("%v/0x%s.mp4",CGAME_VIDEOS_URL,seed)
    return url
}
func check_resource(url string) (bool,error) {

    response, err := http.Head(url)
    if err != nil {
        return false, err
    }
	if response.StatusCode != 200 {
		if response.StatusCode == 403 {
			err = errors.New(fmt.Sprintf("Resource %v not found (url=%v) (status = %v)",url,response.StatusCode))
		} else {
			err = errors.New(fmt.Sprintf("Error: HTTP status code = %v for url=%v",response.StatusCode,url))
		}
		return false,err
	}
	return true,nil
}
func get_last_token_id_randomwalk(host,dbname,user,pass string) (int64,error) {

	err,dbobj := pg_connect_db(host,dbname,user,pass)
	if err != nil {
		return -1,err
	}
	defer dbobj.Close()
	var query string
	query = "SELECT token_id FROM rw_mint_evt ORDER BY id DESC LIMIT 1"
	var last_token_id int64
	err = dbobj.QueryRow(query).Scan(&last_token_id)
	if err != nil {
		return -1,err
	}
	return last_token_id,nil
}
func get_last_token_seed_cosmicgame(host,dbname,user,pass string) (string,error) {

	err,dbobj := pg_connect_db(host,dbname,user,pass)
	if err != nil {
		return "",err
	}
	defer dbobj.Close()
	var query string
	query = "SELECT seed FROM cg_mint_event ORDER BY id DESC LIMIT 1"
	var last_token_seed string
	err = dbobj.QueryRow(query).Scan(&last_token_seed)
	if err != nil {
		return "",err
	}
	return last_token_seed,nil
}
func check_randomwalk_resource_availability() {

	token_id,err:=get_last_token_id_randomwalk(
		os.Getenv("IMG_RWALK_MOITOR_HOST"),
		os.Getenv("IMG_RWALK_MONITOR_DBNAME"),
		os.Getenv("IMG_RWALK_MONITOR_USER"),
		os.Getenv("IMG_RWALK_MONITOR_PASS"),
	)
	if err != nil {
		update_global_errors(fmt.Sprintf("Image check err: %v\n",err.Error()))
		return
	}
	url := fmt_url_addr_for_image_randomwalk(token_id)
	success,err := check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
	url = fmt_url_addr_for_video_randomwalk(token_id)
	success,err = check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
}
func check_cosmicgame_resource_availability() {

	seed,err:=get_last_token_seed_cosmicgame(
		os.Getenv("IMG_CGAME_MOITOR_HOST"),
		os.Getenv("IMG_CGAME_MONITOR_DBNAME"),
		os.Getenv("IMG_CGAME_MONITOR_USER"),
		os.Getenv("IMG_CGAME_MONITOR_PASS"),
	)
	if err != nil {
		update_global_errors(fmt.Sprintf("Image check err: %v\n",err.Error()))
		return
	}
	url := fmt_url_addr_for_image_cosmicgame(seed)
	success,err := check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
	url = fmt_url_addr_for_video_cosmicgame(seed)
	success,err = check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
}
