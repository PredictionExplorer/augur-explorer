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
	"time"
	"strings"
	"syscall"
	"os/signal"

	"github.com/PredictionExplorer/augur-explorer/wanotif"
)
const (
	NUM_TRIES				int = 5
	DELAY_SECONDS			= 2
)
var (
	UrlList					map[string]string		// list of urls to check
	NotifiedPeople			map[string]string
	UrlStatusNumFails		map[string]int

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
			Info.Printf("Res: %+v\n",res)
		} else {
			Info.Printf("Notified failure to %v (%v): %v",person,phone,notif_msg)
		}
	}
}
func check_single_url(url string,msg_header string) {

	fmt.Printf("check_single_url(): %v\n",url)
	var err_str string
	resp,err := http.Get(url)
	fmt.Printf("err=%v\n",err)
	if err != nil {
		err_str = fmt.Sprintf("Networking error: %v",err.Error())
	} else {
		if resp.StatusCode == 200 {
			UrlStatusNumFails[url]=0
			return	// no problem with URL
		}
		err_str = fmt.Sprintf("%v. HTTP status: %v",msg_header,resp.StatusCode)
	}
	num_fails := UrlStatusNumFails[url] + 1
	UrlStatusNumFails[url]=num_fails
	if num_fails >= NUM_TRIES {
		notify_failure(err_str)
		UrlStatusNumFails[url]=0
		Info.Printf("Notifying failure of url %v\n",url)
	} else {
		Info.Printf("Url %v , error: %v, num_fails=%v (%v < %v)\n",url,err,num_fails,num_fails,NUM_TRIES)
	}
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
		fmt.Printf("URL list is empty, aborting.\n")
		os.Exit(1)
	}
	counter:=int(0)
	for row_num,entry := range parsed {
		fmt.Printf("entry = %v\n",entry)
		fields:=strings.Split(entry,"\t")
		if len(fields) != 2 {
			if len(entry) > 0 {
				fmt.Printf("Missing tab separator at line %v (%v)\n",row_num,entry)
				os.Exit(1)
			} else {// empty line was found
				continue
			}
		} else {
			counter++
		}
		fmt.Printf("fields=%+v (len=%v)\n",fields,len(fields))
		fmt.Printf("fields[0]=%v\n",fields[0])
		fmt.Printf("fields[1]=%v\n",fields[1])
		UrlList[fields[0]]=fields[1]
	}
	Info.Printf("Loaded %v URLs\n",counter)
}
func main() {

	UrlList = make(map[string]string)
	NotifiedPeople = make(map[string]string)
	UrlStatusNumFails = make(map[string]int)
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v [url_list_file]\n",os.Args[0])
		os.Exit(1)
	}
	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	read_urls(os.Args[1])



	phones := os.Getenv("PHONE_LIST")
	if len(phones)==0 {
		fmt.Printf("PHONE_LIST environment variable not set\n")
		os.Exit(1)
	}
	counter := int(0)
	for row_num,entry := range strings.Split(phones,",") {
		person_phone := strings.Split(entry,":")
		if len(person_phone) != 2 {
			fmt.Printf("person_phone=%v\n",person_phone)
			fmt.Printf("Loading phones. Entry line %v has invalid format (read %v fields)\n",row_num,len(person_phone))
			os.Exit(1)
		} else {
			counter++
		}
		person:=person_phone[0]
		phone:=person_phone[1]
		NotifiedPeople[person]=phone
	}
	if len(NotifiedPeople) == 0 {
		fmt.Printf("Loading phones. After parsing phone list, the map is empty, aborting.\n")
		os.Exit(1)
	}
	Info.Printf("Loaded %v phones for notification\n",counter)
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
		check_urls()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Print("Exiting upon user request\n")
					os.Exit(1)
				}
				default:
		}
		time.Sleep(DELAY_SECONDS * time.Second)
	}
}
