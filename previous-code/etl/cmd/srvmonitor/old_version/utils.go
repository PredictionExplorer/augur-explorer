package main

import (
	"fmt"
	"net"
	"errors"
	"github.com/nsf/termbox-go"
	"database/sql"
)

func printAtPosition(x, y int, text string, fg, bg termbox.Attribute) {
	for i, r := range text {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
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
