package main

import (
	"sync"
	"fmt"
	"time"
	"github.com/nsf/termbox-go"
)
func init_layer1_status_struct(s *Layer1Status,name,host,dbname,user,pass string,x,y int) {
	s.Name = name
	s.Host = host
	s.DbName = dbname
	s.User = user
	s.Pass = pass
	s.X = x
	s.Y = y
}
func print_layer1_status_line(status *Layer1Status) {

	printAtPosition(status.X,status.Y,status.Name,termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(status.X+35,status.Y,status.Host,termbox.ColorWhite,termbox.ColorDefault)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN "
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorRed,termbox.ColorDefault)
	} else {
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorGreen,termbox.ColorDefault)
	}
	printAtPosition(status.X+70,status.Y,fmt.Sprintf("%v",status.LastBlockNum),termbox.ColorBlue,termbox.ColorDefault)
	if len(status.ErrStr) > 0 {
		update_global_errors(status.ErrStr)
	}
}
func print_current_layer1_status() {
	printAtPosition(0, 11, "--------------------- SQL DB --------------------------------",termbox.ColorWhite,termbox.ColorDefault)
	print_layer1_status_line(&db1)
	print_layer1_status_line(&db2)
	print_layer1_status_line(&db3)
	print_layer1_status_line(&db4)
	termbox.Flush()
}
func check_sql_db_status_layer1(status *Layer1Status,wg *sync.WaitGroup) {

	status.ErrStr = ""
	status.Alive = false
	err,dbobj := pg_connect_db(status.Host,status.DbName,status.User,status.Pass)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v",err)
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	var bnum1 int64
	err = dbobj.QueryRow("SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1").Scan(&bnum1)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v",err)
		update_global_errors(status.ErrStr)
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
	status.LastBlockNum = bnum2
	wg.Done()
	defer dbobj.Close()
}
