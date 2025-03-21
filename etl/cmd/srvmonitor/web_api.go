package main

import (
	"time"
	"fmt"
	"sync"
	"github.com/nsf/termbox-go"
	"net/http"
)
func init_web_api_status_struct(s *WebApiStatus,title string,host string,port string,uri string,x int,y int) {
	s.Title = title
	s.Host = host
	s.Port = port
	s.URI = uri
	s.X = x
	s.Y = y
}
func check_web_api_status(status *WebApiStatus, wg *sync.WaitGroup) {
	url := fmt.Sprintf("http://%s:%s%v", status.Host, status.Port, status.URI)
    client := http.Client{
        Timeout: 10 * time.Second,
    }
	resp, err := client.Get(url)
	if resp != nil {
	    defer resp.Body.Close()
	}
    if err != nil {
        status.Alive = false
    } else {
		status.Alive = (resp.StatusCode >= 200) && (resp.StatusCode < 500)
	}
	wg.Done()
}
func print_web_api_status_line(status *WebApiStatus) {

	printAtPosition(status.X,status.Y,status.Title,termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(status.X+25,status.Y,status.Host+":"+status.Port,termbox.ColorWhite,termbox.ColorDefault)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN "
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorRed,termbox.ColorDefault)
	} else {
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorGreen,termbox.ColorDefault)
	}
	if len(status.ErrStr) > 0 {
		update_global_errors(status.ErrStr)
	}
}
func print_current_web_api_status() {
	printAtPosition(0, 30 , "--------------------- Web API ------------------------------",termbox.ColorWhite,termbox.ColorDefault)
	print_web_api_status_line(&web1)
	print_web_api_status_line(&web2)
	print_web_api_status_line(&web3)
	print_web_api_status_line(&web4)
	termbox.Flush()
}
