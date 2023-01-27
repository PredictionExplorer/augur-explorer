// Daemon to monitor services and send alarms in case notification is required
// URL list file format: URL\tMessage_header\n  per line, separate with \n
// Notified people list format: person_name[separator1]phone[separator2]
//		separator1 - colon, to separate name and number
//		separator2 - comma, to separate records
package main
import (
	"os"
	"fmt"
	"net/http"
	"log"
	"strings"
	"syscall"
	"os/signal"

	"github.com/PredictionExplorer/augur-explorer/wanotif"
)
var (
	UrlList			map[string]string		// list of urls to check
	NotifiedPeople	map[string]string

	Info			*log.Logger
	wa				*wanotif.Whatsapp
)
func notify_failure(notif_msg string) {

	for person,phone := range NotifiedPeople {
		res,err := wa.SendText(phone,notif_msg)
		if err != nil {
			Info.Printf(
				"Error sending whatsapp request to %v: %v (res=%+v,  phone=%v, msg=%v)",
				person,err,res,phone,notif_msg,
			)
		} else {
			Info.Printf("Notified failure to %v (%v): %v",person,phone,notif_msg)
		}
	}
}
func check_single_url(url string,msg_header string) {

	var err_str string
	resp,err := http.Get(url)
	if err != nil {
		err_str = fmt.Sprintf("Networking error: %v",err.Error())
	} else {
		if resp.StatusCode == 200 {
			return	// no problem with URL
		}
		err_str = fmt.Sprintf("%v. HTTP status: %v",msg_header,resp.StatusCode)
	}
	notify_failure(err_str)
}
func check_urls() {
	for k,v := range UrlList {
		check_single_url(k,v)
	}
}
func read_urls(filename string) {

	data,err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading urls: %v\n",err)
		os.Exit(1)
	}
	parsed := strings.Split(string(data),"\n")
	if len(parsed)==0 {
		fmt.Printf("URL list is empty, aborting.")
		os.Exit(1)
	}
	counter:=int(0)
	for row_num,entry := range parsed {
		counter=row_num+1
		fields:=strings.Split(entry,"\t")
		if len(fields) != 2 {
			fmt.Printf("Missing tab separator at line %v (%v)",counter,entry)
			os.Exit(1)
		}
		UrlList[fields[1]]=fields[0]
	}
	Info.Printf("Loaded %v URLs\n",counter)
}
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v [url_list_file]\n",os.Args[0])
		os.Exit(1)
	}
	read_urls(os.Args[1])

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	phones := os.Getenv("PHONE_LIST")
	counter := int(0)
	for row_num,entry := range strings.Split(phones,",") {
		counter = row_num+1
		person_phone := strings.Split(entry,":")
		if len(person_phone) != 2 {
			fmt.Printf("Entry line %v has invalid format (read %v fields)",counter,len(person_phone))
			os.Exit(1)
		}
		person:=person_phone[0]
		phone:=person_phone[1]
		NotifiedPeople[person]=phone
	}
	if len(NotifiedPeople) == 0 {
		fmt.Printf("After parsing phone list, the map is empty, aborting.")
		os.Exit(1)
	}
	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over ." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	token := os.Getenv("WHATSAPP_TOKEN")
	my_phone_id := os.Getenv("WHATSAPP_PHONE_ID")
	if len(token) == 0 {
		fmt.Printf("WHATSAPP_TOKEN environment variable is empty, aborting.\n")
		os.Exit(1)
	}
	if len(my_phone_id) == 0 {
		fmt.Printf("WHATSAPP_PHONE_ID environment variable is empty, aborting\n")
		os.Exit(1)
	}
	wa = wanotif.NewWhatsapp(token,my_phone_id)
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Print("Exiting upon user request\n")
				}
				default:
			}
	}
}
