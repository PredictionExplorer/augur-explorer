package main

import (
	"sync"
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
	wg.Done()
	defer dbobj.Close()
}
func print_application_layer_status_line(status *AppLayerStatus) {

	printAtPosition(status.X,status.Y,status.Title,termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(status.X+35,status.Y,fmt.Sprintf("%9d",status.LastBlockNum),termbox.ColorBlue,termbox.ColorDefault)
	if len(status.ErrStr) > 0 {
		update_global_errors(status.ErrStr)
	}
}
func print_current_application_layer_status() {
	printAtPosition(90, 0, "---- Last Block Numbers for App layer ---",termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(100, 1, "( Postgres database)",termbox.ColorWhite,termbox.ColorDefault)
	print_application_layer_status_line(&cosmic_app1)
	print_application_layer_status_line(&cosmic_app2)
	print_application_layer_status_line(&rwalk_app1)
	print_application_layer_status_line(&rwalk_app2)
	termbox.Flush()
}
