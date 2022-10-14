// Extracts epoch fields for analyzing randomness in the assignment of validators
package main

import (
	//"net/http"
	"os"
	"fmt"
	"log"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	//API_URL string = "http://127.0.0.1:3500"
	API_URL string ="http://170.187.142.12:14343"
	BLOCK_INCREMENT int = -1
	REQUEST_DELAY int = 5
)
var (
	Error   *log.Logger
	Info	*log.Logger
)
func main() {


	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)

	fname:=fmt.Sprintf("%v/beacon_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/beacon_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	main_event_loop()
}
