// Used for monitoring our server installation

package main

import (
//	"net/http"
	"os"
	"fmt"
	"log"
	"time"
	"net"
	"errors"
//	"io/ioutil"
	"context"
	"sync"
	"database/sql"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nsf/termbox-go"

//	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
type RPCStatus struct {
	LastBlockNum		int64
	Alive				bool	// if there is block difference over last 60 seconds, node is alive
	RPCUrl				string
	RPCName				string
	ErrStr				string
	X					int
	Y					int
}
type Layer1Status struct {
	LastBlockNum		int64
	Alive				bool
	DbName				string
	Host				string
	User				string
	Name				string
	Pass				string
	ErrStr				string
	X					int
	Y					int
}
const (
	WAIT_RPC_BLOCK_NUM	= 60		// seconds to wait before second getBlock() call
	WAIT_DB_BLOCK_NUM = 60			// seconds to wait to detect incremental database update
)
var (
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage

	rpc0,rpc1,rpc2,rpc3,rpc4,rpc5,rpc6		RPCStatus
	db1,db2,db3								Layer1Status
)
func printAtPosition(x, y int, text string, fg, bg termbox.Attribute) {
	for i, r := range text {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
func init_rpc_status_struct(s *RPCStatus,name string,url string,x int,y int) {
	s.RPCName = name
	s.RPCUrl = url
	s.X = x
	s.Y = y
}
func check_rpc_status(status *RPCStatus, wg *sync.WaitGroup) {

	if len(status.RPCUrl) == 0 {
		status.RPCUrl = "*** not set ***"
		wg.Done()
		return
	}
	status.ErrStr = ""
	status.Alive = false
	rpc_obj, err:=rpc.DialContext(context.Background(), status.RPCUrl)
	if err != nil {
		status.ErrStr = err.Error()
		wg.Done()
		return
	}
	eclient := ethclient.NewClient(rpc_obj)
	latestBlock1, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
		wg.Done()
		return
    }
	time.Sleep(WAIT_RPC_BLOCK_NUM*time.Second)
	latestBlock2, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
		wg.Done()
		return
    }
	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	if diff == 0 {
		status.ErrStr=fmt.Sprintf("Block difference is zero (last block = %v)",latestBlock2.Number.Int64())
	} else {
		status.Alive = true
	}
	status.LastBlockNum = latestBlock2.Number.Int64()
	wg.Done()

}
func print_rpc_status_line(status *RPCStatus) {

	printAtPosition(status.X,status.Y,status.RPCName,termbox.ColorWhite,termbox.ColorBlack)
	printAtPosition(status.X+22,status.Y,status.RPCUrl,termbox.ColorWhite,termbox.ColorBlack)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN"
		printAtPosition(status.X+55,status.Y,alive_str,termbox.ColorRed,termbox.ColorBlack)
	} else {
		printAtPosition(status.X+55,status.Y,alive_str,termbox.ColorGreen,termbox.ColorBlack)
	}
	printAtPosition(status.X+70,status.Y,fmt.Sprintf("%v",status.LastBlockNum),termbox.ColorBlue,termbox.ColorBlack)
	printAtPosition(status.X+85,status.Y,fmt.Sprintf("%v",status.ErrStr),termbox.ColorYellow,termbox.ColorBlack)
}
func print_current_rpc_status() {
	print_rpc_status_line(&rpc0)
	print_rpc_status_line(&rpc1)
	print_rpc_status_line(&rpc2)
	print_rpc_status_line(&rpc3)
	print_rpc_status_line(&rpc4)
	print_rpc_status_line(&rpc5)
	termbox.Flush()
}
func check_rpc_services() {

	var wg_rpcs sync.WaitGroup
	wg_rpcs.Add(5);
	init_rpc_status_struct(&rpc0,os.Getenv("RPC0_NAME"),os.Getenv("RPC0_URL"),1,0)
	init_rpc_status_struct(&rpc1,os.Getenv("RPC1_NAME"),os.Getenv("RPC1_URL"),1,1)
	init_rpc_status_struct(&rpc2,os.Getenv("RPC2_NAME"),os.Getenv("RPC2_URL"),1,2)
	init_rpc_status_struct(&rpc3,os.Getenv("RPC3_NAME"),os.Getenv("RPC3_URL"),1,3)
	init_rpc_status_struct(&rpc4,os.Getenv("RPC4_NAME"),os.Getenv("RPC4_URL"),1,4)
	init_rpc_status_struct(&rpc5,os.Getenv("RPC5_NAME"),os.Getenv("RPC4_URL"),1,5)
	go check_rpc_status(&rpc0,&wg_rpcs); 
	go check_rpc_status(&rpc1,&wg_rpcs); 
	go check_rpc_status(&rpc2,&wg_rpcs); 
	go check_rpc_status(&rpc3,&wg_rpcs); 
	go check_rpc_status(&rpc4,&wg_rpcs); 
	go check_rpc_status(&rpc5,&wg_rpcs); 
	wg_rpcs.Wait() 
	print_current_rpc_status()
}
func init_layer1_status_struct(s *Layer1Status,name,host,dbname,user,pass string,x,y int) {
	s.Name = name
	s.Host = host
	s.DbName = dbname
	s.User = user
	s.Pass = pass
	s.X = x
	s.Y = y
}
func pg_connect_db(host_port,db_name,user,pass string) (error , *sql.DB) {
    var err error
    host,port,err:=net.SplitHostPort(host_port)
    if (err!=nil) {
        host=host_port
        port="5432"
    }   
    conn_str := "user='"+user+"' dbname='" + db_name + "' password='" + pass +
                "' host='" + host + "' port='" + port + "'";
    dbobj,err := sql.Open("postgres",conn_str);
    if (err!=nil) {
        return errors.New(fmt.Sprintf("Error connecting: %v\n",err)),nil
    }   
    _,err = dbobj.Exec("SET timezone TO 0")        // Setting timezone to UTC 
    if (err!=nil) {
        return errors.New(fmt.Sprintf("DB Error: %v",err)),nil
    }
	return nil,dbobj
}
func check_sql_db_status_layer1(status *Layer1Status,wg *sync.WaitGroup) {

	status.ErrStr = ""
	status.Alive = false
	err,dbobj := pg_connect_db(status.Host,status.DbName,status.User,status.Pass)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v",err)
		wg.Done()
		return
	}
	var bnum1 int64
	err = dbobj.QueryRow("SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1").Scan(&bnum1)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		wg.Done()
		return
	}
	time.Sleep(WAIT_RPC_BLOCK_NUM*time.Second)
	var bnum2 int64
	err = dbobj.QueryRow("SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1").Scan(&bnum2)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		wg.Done()
		return
	}
	diff := bnum2 - bnum1
	if diff == 0 {
		status.ErrStr=fmt.Sprintf("Block difference is zero (last block = %v)",bnum2)
	} else {
		status.Alive = true
	}
//	fmt.Printf("sql diff = %v\n",diff)
	status.LastBlockNum = bnum2
	wg.Done()
}
func print_layer1_status_line(status *Layer1Status) {

	printAtPosition(status.X,status.Y,status.Name,termbox.ColorWhite,termbox.ColorBlack)
	printAtPosition(status.X+22,status.Y,status.Host,termbox.ColorWhite,termbox.ColorBlack)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN"
		printAtPosition(status.X+55,status.Y,alive_str,termbox.ColorRed,termbox.ColorBlack)
	} else {
		printAtPosition(status.X+55,status.Y,alive_str,termbox.ColorGreen,termbox.ColorBlack)
	}
	printAtPosition(status.X+70,status.Y,fmt.Sprintf("%v",status.LastBlockNum),termbox.ColorBlue,termbox.ColorBlack)
	printAtPosition(status.X+85,status.Y,fmt.Sprintf("%v",status.ErrStr),termbox.ColorYellow,termbox.ColorBlack)
}
func print_current_layer1_status() {
	print_layer1_status_line(&db1)
	print_layer1_status_line(&db2)
	print_layer1_status_line(&db3)
	termbox.Flush()
}
func check_layer1() {

	var wg_db sync.WaitGroup
	wg_db.Add(3);
	init_layer1_status_struct(&db1,os.Getenv("DB_RWALK_L1_NAME_SRV1"),os.Getenv("DB_RWALK_L1_HOST_SRV1"),os.Getenv("DB_RWALK_L1_DBNAME_SRV1"),os.Getenv("DB_RWALK_L1_USER_SRV1"),os.Getenv("DB_RWALK_L1_PASS_SRV1"),1,10)
	init_layer1_status_struct(&db2,os.Getenv("DB_RWALK_L1_NAME_SRV2"),os.Getenv("DB_RWALK_L1_HOST_SRV2"),os.Getenv("DB_RWALK_L1_DBNAME_SRV2"),os.Getenv("DB_RWALK_L1_USER_SRV2"),os.Getenv("DB_RWALK_L1_PASS_SRV2"),1,11)
	init_layer1_status_struct(&db3,os.Getenv("DB_RWALK_L1_NAME_SRV3"),os.Getenv("DB_RWALK_L1_HOST_SRV3"),os.Getenv("DB_RWALK_L1_DBNAME_SRV3"),os.Getenv("DB_RWALK_L1_USER_SRV3"),os.Getenv("DB_RWALK_L1_PASS_SRV3"),1,12)
	go check_sql_db_status_layer1(&db1,&wg_db); 
	go check_sql_db_status_layer1(&db2,&wg_db); 
	go check_sql_db_status_layer1(&db3,&wg_db); 
	wg_db.Wait() 
	print_current_layer1_status()
}
func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatalf("Failed to initialize termbox: %v", err)
	}
	defer termbox.Close()
	//Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	fmt.Printf("\n\n\n\n\n\n")

//	storage = Connect_to_storage(Info)
	//check_rpc_services()
//	check_sql_db_status_layer1(os.Getenv("DB_RWALK_HOST_S1"),os.Getenv("DB_RWALK_DBNAME_S1"),os.Getenv("DB_RWALK_USER_S1"),os.Getenv("DB_RWALK_PASS_S1"),os.Getenv("DB_RWALK_L1_)
	check_layer1()
	printAtPosition(3,20,fmt.Sprintf("Press any key to exit"),termbox.ColorGreen,termbox.ColorBlack)
	termbox.PollEvent()
}
