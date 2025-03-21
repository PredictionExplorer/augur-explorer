package main

import (
	"sync"
	"math"
	"fmt"
	"github.com/nsf/termbox-go"
)

func init_application_layer_status_struct(s *AppLayerStatus,title,host,dbname,user,pass,table_name string,x,y int) {
	s.Title = title
	s.Host = host
	s.DbName = dbname
	s.User = user
	s.Pass = pass
	s.TableName = table_name
	s.X = x
	s.Y = y
}
func check_sql_db_status_application(status *AppLayerStatus,wg *sync.WaitGroup) {

	status.ErrStr = ""
	err,dbobj := pg_connect_db(status.Host,status.DbName,status.User,status.Pass)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v",err)
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	var chain_id_str string
	err = dbobj.QueryRow("SELECT chain_id FROM contract_addresses").Scan(&chain_id_str)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	var last_evt_id int64
	err = dbobj.QueryRow("SELECT last_evt_id FROM "+status.TableName).Scan(&last_evt_id)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	var bnum int64
	err = dbobj.QueryRow("SELECT block_num FROM evt_log WHERE id=$1",last_evt_id).Scan(&bnum)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	status.LastBlockNum = bnum
	if chain_id_str == "42161" {
		if Official_arbitrum_ptr != nil {
			if Official_arbitrum_ptr.LastBlockNum != 0 {
				status.OfficialLagDiff = Official_arbitrum_ptr.LastBlockNum - status.LastBlockNum
			}
		}
	}
	if chain_id_str == "421614" {
		if Official_sepolia_arb_ptr != nil {
			if Official_sepolia_arb_ptr.LastBlockNum != 0 {
				status.OfficialLagDiff = Official_sepolia_arb_ptr.LastBlockNum - status.LastBlockNum
			}
		}
	}
	wg.Done()
	defer dbobj.Close()
}
func print_application_layer_status_line(status *AppLayerStatus) {

	printAtPosition(status.X,status.Y,status.Title,termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(status.X+35,status.Y,fmt.Sprintf("%9d",status.LastBlockNum),termbox.ColorBlue,termbox.ColorDefault)
	if len(status.ErrStr) > 0 {
		update_global_errors(status.ErrStr)
	}
	var official_diff string = "------"
	if status.OfficialLagDiff != math.MaxInt64 {
		official_diff = fmt.Sprintf("%6v",status.OfficialLagDiff)
	}
	printAtPosition(status.X+45,status.Y,official_diff,termbox.ColorBlue,termbox.ColorDefault)
}
func print_current_application_layer_status() {
	printAtPosition(1, 24, "----------- Last Block Numbers in Postgres ----------",termbox.ColorWhite,termbox.ColorDefault)
	print_application_layer_status_line(&cosmic_app1)
	print_application_layer_status_line(&cosmic_app2)
	print_application_layer_status_line(&rwalk_app1)
	print_application_layer_status_line(&rwalk_app2)
	termbox.Flush()
}
