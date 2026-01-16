package main

import (
	"sync"
	"os/exec"
	"strings"
	"github.com/nsf/termbox-go"

)

func init_df_status_struct(s *DfStatus,title,user,ip,device_list string,x,y int) {
	s.Title = title
	s.User= user
	s.Ip = ip
	s.DeviceList = device_list
	s.X = x
	s.Y = y
}
func print_df_for_server(status *DfStatus,wg *sync.WaitGroup) {

	cmd := exec.Command("/usr/bin/ssh","-l",status.User,status.Ip,"df --output=target,pcent",status.DeviceList)
	output,err := cmd.Output()
	if err != nil {
		status.ErrStr = err.Error()
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	termboxMutex.Lock()
	defer termboxMutex.Unlock()
	printAtPosition(status.X+3,status.Y,status.Title,termbox.ColorYellow,termbox.ColorDefault)
	lines := strings.Split(string(output),"\n")
	for i:=1; i<(len(lines)-1);i++ {
		line:=lines[i];
		printAtPosition(status.X,status.Y+i,line,termbox.ColorWhite,termbox.ColorDefault)
	}
	termbox.Flush()
	wg.Done()

}
