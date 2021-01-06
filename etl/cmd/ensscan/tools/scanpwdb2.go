// Scans another password database
package main

import (
	"os"
	"path/filepath"
	"fmt"
	"bufio"
	"encoding/hex"
	"strings"
	"unicode/utf8"
	"os/signal"
	"syscall"
	"io/ioutil"
	"log"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	"github.com/ethereum/go-ethereum/rlp"
)
const (
	MAX_FILE_NAME int = 256
)
type StatusFile struct {
	CurrentFile			string
	AllFiles			[]string
}
var (
	statusf				StatusFile
	all_labels			map[string]struct{} = make(map[string]struct{})
	Info				*log.Logger
	storage				*SQLStorage
	market_order_id int64 = 0
	status_file_name	string
)
func dump_all_files() {

	fmt.Printf("Dumping file list:\n")
	for i:=0;i<len(statusf.AllFiles);i++ {
		fmt.Printf("\t%v\n",statusf.AllFiles[i])
	}
	fmt.Printf("\n\n")
}
func process_token(token string) {
	// Note: passwords are all unique, so we don't have unique index check here
	if len(token)==0 {
		return
	}
	token = strings.ReplaceAll(token,` `,``)
	token = strings.ReplaceAll(token,`'`,``)
	token = strings.ReplaceAll(token,"\t",``)
	token = strings.ReplaceAll(token,"\r",``)
	token = strings.ReplaceAll(token,`\`,``)
	if len(token)==0 {
		return
	}
	if !utf8.ValidString(token) {
		return
	}
	hash,err := LabelHash(token)
	if err != nil {
		return
	}
	hash_str := hex.EncodeToString(hash[:])
	_,exists := all_labels[hash_str]
	if exists {
		if !storage.Label_exists_in_ens_labels(hash_str) {
			fmt.Printf("%v\t%v\n",token,hash_str)
			storage.Insert_word_for_ens_label(token,hash_str)
			all_labels[hash_str]=struct{}{}
		}
	}
	return
}
func process_single_file(fname string) {

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error when reading %v: %v\n",fname,err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf,32*1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := strings.Split(line,":")
		for j:=0 ; j < len(parsed) ; j ++ {
			subsub_tok := parsed[j]
			subsub_tokens := strings.Split(subsub_tok,"@")
			for n:=0; n<len(subsub_tokens); n++ {
				sub_tok := subsub_tokens[n]
				tokens := strings.Split(sub_tok,".")
				for k:=0 ; k<len(tokens) ; k++ {
					tok := tokens[k]
					process_token(tok)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}
}
func process(exit_chan chan bool) {

//	dump_all_files()
	fmt.Printf("Current file name: %v\n",statusf.CurrentFile)
	var current_idx int = 0
	if len(statusf.CurrentFile) > 0 {
		var i int = 0
		for i=0 ; i<len(statusf.AllFiles); i++ {
			if statusf.CurrentFile == statusf.AllFiles[i] {
				break;
			}
		}
		current_idx = i+1
	}

	for {
		if current_idx > (len(statusf.AllFiles)-1) {
			fmt.Printf("Finishing...\n")
			break
		}
		filename := statusf.AllFiles[current_idx]
		fmt.Printf("Processing file %v\n",filename)
		process_single_file(filename)
		statusf.CurrentFile = filename
		write_status(status_file_name)
		current_idx++
		select {
			case exit_flag := <-exit_chan:
			if exit_flag {
				fmt.Println("Exiting by user request.")
				fmt.Printf("Done. Exiting...\n")
				os.Exit(0)
			}
			default:
		}
	}
}
func write_status(status_file string) {
	encoded,err := rlp.EncodeToBytes(&statusf)
	if err != nil {
		fmt.Printf("Error at RLP encoding: %v\n",err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(status_file, encoded[:], 0644)
	if err != nil {
		fmt.Printf("Error : %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Latest processed file was %v\n",statusf.CurrentFile)
}
func main() {

	if len(os.Args) != 3 {
		fmt.Printf("usage: %v [directory_to_scan] [status_work_file_name]\n"+
					"\t(the status file is going to store internal variables of the process\n\n")
		os.Exit(1)
	}
	source_dir := os.Args[1]
	status_file_name = os.Args[2]

	_, err := os.Stat(status_file_name);
	if 	os.IsNotExist(err) {
		fmt.Printf("Starting a brand new process. (for continuation of pevious process specify a non-empty status file)\n")
		fmt.Printf("Reading directory structure from %v\n",source_dir)
		var num_files int = 0
		err = filepath.Walk(source_dir, func(path string, info os.FileInfo, err error) error {
			if  info.Mode().IsRegular() {
				statusf.AllFiles = append(statusf.AllFiles,path)
				num_files++
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		fmt.Printf("%v files were read\n",num_files)
	} else {
		fmt.Printf("Loading previous state variables from %v\n",status_file_name)
		encoded, err := ioutil.ReadFile(status_file_name)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		rlp.DecodeBytes(encoded,&statusf)
		fmt.Printf("%v files were loaded\n",len(statusf.AllFiles))
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	storage.Get_all_ens_labels_from_owners(&all_labels)
	fmt.Printf("Loaded %v ENS label records.\n",len(all_labels))

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("Got SIGINT signal, will exit after processing is over." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()
	process(exit_chan)
}
