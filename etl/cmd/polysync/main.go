// Syncs PolyMarkets markets that are published at this URL:
//		https://strapi-matic.poly.market/markets

package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"io/ioutil"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	POLYMARKETS_API_URL string = "https://strapi-matic.poly.market/markets?_limit=-1"
)
var (
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage
)
func main() {


	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/polysync_db.log",log_dir)

	fname:=fmt.Sprintf("%v/polysync_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/polysync_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)
/*
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)
*/
	storage = Connect_to_storage(Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	resp,err := http.Get(POLYMARKETS_API_URL)
	if err != nil {
		Error.Printf("Couldn't successfuly GET at % : %v\n",POLYMARKETS_API_URL,err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response  %+v\n",body)
	ioutil.WriteFile("/tmp/body.txt",body,0644)
	scan_markets(body)
}
