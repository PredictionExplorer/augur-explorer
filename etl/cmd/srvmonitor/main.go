// Used for monitoring our server installation

// df command line:
//		df --output=target,pcent '/dev/nvme0n1p1' '/dev/nvme1n1p1' '/dev/bcache0' /dev/mapper/ubuntu--vg-ubuntu--lv
package main

import (
//	"net/http"
	"os"
	"fmt"
	"log"
	"time"
	"sync"
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
type DfStatus struct {	// note: this is for passwordless execution, copy id_rsa.pub to authorized_hosts on the destination server
	Title				string
	User				string
	Ip					string
	DeviceList			string
	X					int
	Y					int
	ErrStr				string
}
type AppLayerStatus struct {	// fetches last block number that was processed by application layer
	Title				string
	LastBlockNum		int64
	DbName				string
	Host				string
	User				string
	Name				string
	Pass				string
	TableName			string
	ErrStr				string
	X					int
	Y					int
}
const (
	WAIT_RPC_BLOCK_NUM	= 60			// seconds to wait before second getBlock() call
	WAIT_DB_BLOCK_NUM = 60				// seconds to wait to detect incremental database update
	WAIT_BETWEEN_UPDATES = 30			// seconds to wait after each poll for data
	WAIT_BETWEEN_UPDATES_DFCMD = 600	// seconds to wait after each poll for data
)
var (
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage

	rpc0,rpc1,rpc2,rpc3,rpc4,rpc5,rpc6,rpc7		RPCStatus
	db1,db2,db3								Layer1Status
	df1,df2,df3								DfStatus
	rwalk_app1,rwalk_app2					AppLayerStatus
	cosmic_app1,cosmic_app2					AppLayerStatus
)
func check_rpc_services() {

	for {
		var wg_rpcs sync.WaitGroup
		wg_rpcs.Add(8);
		init_rpc_status_struct(&rpc0,os.Getenv("RPC0_NAME"),os.Getenv("RPC0_URL"),1,1)
		init_rpc_status_struct(&rpc1,os.Getenv("RPC1_NAME"),os.Getenv("RPC1_URL"),1,2)
		init_rpc_status_struct(&rpc2,os.Getenv("RPC2_NAME"),os.Getenv("RPC2_URL"),1,3)
		init_rpc_status_struct(&rpc3,os.Getenv("RPC3_NAME"),os.Getenv("RPC3_URL"),1,4)
		init_rpc_status_struct(&rpc4,os.Getenv("RPC4_NAME"),os.Getenv("RPC4_URL"),1,5)
		init_rpc_status_struct(&rpc5,os.Getenv("RPC5_NAME"),os.Getenv("RPC5_URL"),1,6)
		init_rpc_status_struct(&rpc6,os.Getenv("RPC6_NAME"),os.Getenv("RPC6_URL"),1,7)
		init_rpc_status_struct(&rpc7,os.Getenv("RPC7_NAME"),os.Getenv("RPC7_URL"),1,8)
		go check_rpc_status(&rpc0,&wg_rpcs); 
		go check_rpc_status(&rpc1,&wg_rpcs); 
		go check_rpc_status(&rpc2,&wg_rpcs); 
		go check_rpc_status(&rpc3,&wg_rpcs); 
		go check_rpc_status(&rpc4,&wg_rpcs); 
		go check_rpc_status(&rpc5,&wg_rpcs); 
		go check_rpc_status(&rpc6,&wg_rpcs); 
		go check_rpc_status(&rpc7,&wg_rpcs); 
		wg_rpcs.Wait() 
		print_current_rpc_status()
		time.Sleep(WAIT_BETWEEN_UPDATES * time.Second)
	}
}
func check_layer1() {

	init_layer1_status_struct(&db1,os.Getenv("DB_RWALK_L1_NAME_SRV1"),os.Getenv("DB_RWALK_L1_HOST_SRV1"),os.Getenv("DB_RWALK_L1_DBNAME_SRV1"),os.Getenv("DB_RWALK_L1_USER_SRV1"),os.Getenv("DB_RWALK_L1_PASS_SRV1"),1,10)
	init_layer1_status_struct(&db2,os.Getenv("DB_RWALK_L1_NAME_SRV2"),os.Getenv("DB_RWALK_L1_HOST_SRV2"),os.Getenv("DB_RWALK_L1_DBNAME_SRV2"),os.Getenv("DB_RWALK_L1_USER_SRV2"),os.Getenv("DB_RWALK_L1_PASS_SRV2"),1,11)
	init_layer1_status_struct(&db3,os.Getenv("DB_RWALK_L1_NAME_SRV3"),os.Getenv("DB_RWALK_L1_HOST_SRV3"),os.Getenv("DB_RWALK_L1_DBNAME_SRV3"),os.Getenv("DB_RWALK_L1_USER_SRV3"),os.Getenv("DB_RWALK_L1_PASS_SRV3"),1,12)

	for {
		var wg_db sync.WaitGroup
		wg_db.Add(3);
		go check_sql_db_status_layer1(&db1,&wg_db); 
		go check_sql_db_status_layer1(&db2,&wg_db); 
		go check_sql_db_status_layer1(&db3,&wg_db); 
		wg_db.Wait() 
		print_current_layer1_status()
		time.Sleep(WAIT_BETWEEN_UPDATES * time.Second)
	}
}
func show_disk_usage_statistics() {

	init_df_status_struct(&df1,os.Getenv("SSH_CMD_DF_SRV1_NAME"),os.Getenv("SSH_CMD_DF_SRV1_USER"),os.Getenv("SSH_CMD_DF_SRV1_IP"),os.Getenv("SSH_CMD_DF_SRV1_DEVICES"),1,14)
	init_df_status_struct(&df2,os.Getenv("SSH_CMD_DF_SRV2_NAME"),os.Getenv("SSH_CMD_DF_SRV2_USER"),os.Getenv("SSH_CMD_DF_SRV2_IP"),os.Getenv("SSH_CMD_DF_SRV2_DEVICES"),25,14)
	init_df_status_struct(&df3,os.Getenv("SSH_CMD_DF_SRV3_NAME"),os.Getenv("SSH_CMD_DF_SRV3_USER"),os.Getenv("SSH_CMD_DF_SRV3_IP"),os.Getenv("SSH_CMD_DF_SRV3_DEVICES"),50,14)
	for {
		var wg sync.WaitGroup
		wg.Add(3);
		go print_df_for_server(&df1,&wg)
		go print_df_for_server(&df2,&wg)
		go print_df_for_server(&df3,&wg)
		wg.Wait() 
		time.Sleep(WAIT_BETWEEN_UPDATES_DFCMD * time.Second)
	}
}
func show_application_layer_last_blocks() {
	init_application_layer_status_struct(&cosmic_app1,os.Getenv("APP_STATUS_SRV1_TITLE"),os.Getenv("APP_STATUS_SRV1_HOST"),os.Getenv("APP_STATUS_SRV1_DBNAME"),os.Getenv("APP_STATUS_SRV1_USER"),os.Getenv("APP_STATUS_SRV1_PASS"),"cg_proc_status",80,2)
	init_application_layer_status_struct(&cosmic_app2,os.Getenv("APP_STATUS_SRV2_TITLE"),os.Getenv("APP_STATUS_SRV2_HOST"),os.Getenv("APP_STATUS_SRV2_DBNAME"),os.Getenv("APP_STATUS_SRV2_USER"),os.Getenv("APP_STATUS_SRV2_PASS"),"cg_proc_status",80,3)
	init_application_layer_status_struct(&rwalk_app1,os.Getenv("APP_STATUS_SRV3_TITLE"),os.Getenv("APP_STATUS_SRV3_HOST"),os.Getenv("APP_STATUS_SRV3_DBNAME"),os.Getenv("APP_STATUS_SRV3_USER"),os.Getenv("APP_STATUS_SRV3_PASS"),"rw_proc_status",80,4)
	init_application_layer_status_struct(&rwalk_app2,os.Getenv("APP_STATUS_SRV4_TITLE"),os.Getenv("APP_STATUS_SRV4_HOST"),os.Getenv("APP_STATUS_SRV4_DBNAME"),os.Getenv("APP_STATUS_SRV4_USER"),os.Getenv("APP_STATUS_SRV4_PASS"),"rw_proc_status",80,5)

	for {
		var wg sync.WaitGroup
		wg.Add(4);
		go check_sql_db_status_application(&cosmic_app1,&wg)
		go check_sql_db_status_application(&cosmic_app2,&wg)
		go check_sql_db_status_application(&rwalk_app1,&wg)
		go check_sql_db_status_application(&rwalk_app2,&wg)
		wg.Wait()
		print_current_application_layer_status()
		time.Sleep(WAIT_BETWEEN_UPDATES * time.Second)
	}
}
func main() {
	/*
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
    }() */
	logfile, err := os.OpenFile("/tmp/srvmonitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)
	defer os.Remove("/tmp/srvmonitor.log")

	err = termbox.Init()
	if err != nil {
		log.Fatalf("Failed to initialize termbox: %v", err)
	}
	defer termbox.Close()
	fmt.Printf("\n\n\n\n\n\n")

	go check_rpc_services()
	go check_layer1()
	go show_disk_usage_statistics()
//	go show_application_layer_last_blocks()

	termbox.PollEvent()
}
