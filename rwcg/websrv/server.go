package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
)
const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
	DEFAULT_MARKET_ROWS_LIMIT int	= 500
	DEFAILT_MARKET_TRADES_LIMIT int = 20
	DEFAULT_USER_REPORTS_LIMIT int = 30
	DEFAULT_MARKET_REPORTS_LIMIT int = 40

	JSON				bool = true
	HTTP				bool = false
)
type RWCGServer struct {
	db				*SQLStorage
	db_arbitrum		*SQLStorage
}
func (self *RWCGServer) arbitrum_initialized() bool {

	if self.db_arbitrum == nil {
		return false
	}
	return true
}
func connect_to_arbitrum(srv *RWCGServer) {

	arb_user := os.Getenv("ARB_USERNAME")
	arb_passwd := os.Getenv("ARB_PASSWORD")
	arb_db_name := os.Getenv("ARB_DATABASE")
	arb_host_port := os.Getenv("ARB_HOST")
	if len(arb_user) > 0 {
		log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
		db_log_file:=fmt.Sprintf("%v/%v",log_dir,"arbitrum-db.log")
		arbitrum_db_logfile, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Can't start: %v\n",err)
			os.Exit(1)
		}
		arb_DB := log.New(arbitrum_db_logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

		srv.db_arbitrum= New_sql_storage(
			Info,
			arb_DB,
			arb_host_port,
			arb_db_name,
			arb_user,
			arb_passwd,
		)
		Info.Printf("Arbitrum database connection initialized (ARB_USERNAME=%v, ARB_HOST=%v, ARB_DATABASE=%v)\n",
			arb_user, arb_host_port, arb_db_name)
	} else {
		Info.Printf("Arbitrum database connection NOT configured. ARB_USERNAME environment variable is not set. Arbitrum features will be unavailable.\n")
		fmt.Printf("INFO: Arbitrum database connection NOT configured (ARB_USERNAME not set). Arbitrum features will be unavailable.\n")
	}
}
func create_rwcg_server() *RWCGServer {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	web_db_log_file:=fmt.Sprintf("%v/%v",log_dir,"webserver-db.log")

	fname:=fmt.Sprintf("%v/webserver_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/webserver_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)
	srv := new(RWCGServer)
	srv.db = Connect_to_storage(Info)
	srv.db.Init_log(web_db_log_file)
	connect_to_arbitrum(srv)

	return srv
}
func is_address_valid(c *gin.Context,json_output bool,addr string) (string,bool) {

	if (len(addr) != 40) && (len(addr)!=42) {
		var err_msg = fmt.Sprintf("Provided address has invalid length (len=%v)",len(addr))
		if json_output {
			c.JSON(http.StatusOK,gin.H{
				"status": 0,
				"error": err_msg,
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "RWCG: Error",
				"ErrDescr": err_msg,
			})
		}
		return "",false
	}
	if (addr[0]=='0') && (addr[1] == 'x') {
		addr = addr[2:]
	}
	if len(addr) != 40 {
		if json_output {
			c.JSON(http.StatusOK,gin.H{
				"status": 0,
				"error": fmt.Sprintf("Invalid address length"),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "RWCG: Error",
				"ErrDescr": fmt.Sprintf("Invalid address length"),
			})
		}
		return "",false
	}
	var formatted_addr string
	addr_bytes,err := hex.DecodeString(addr)
	if err == nil {
		addr := common.BytesToAddress(addr_bytes)
		formatted_addr = addr.String()
	} else {
		if json_output {
			c.JSON(http.StatusOK,gin.H{
				"status": 0,
				"error": fmt.Sprintf("Provided address parameter is invalid : %v",err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "RWCG: Error",
				"ErrDescr": fmt.Sprintf("Provided address parameter is invalid : %v",err),
			})
		}
		return "",false
	}
	return formatted_addr,true
}
