package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

//./internal/ethapi/api.go
type ExecutionResult struct {
	Gas         uint64         `json:"gas"`
	Failed      bool           `json:"failed"`
	ReturnValue string         `json:"returnValue"`
	StructLogs  []StructLogRes `json:"structLogs"`
}

type StructLogRes struct {
	Pc      uint64             `json:"pc"`
	Op      string             `json:"op"`
	Gas     uint64             `json:"gas"`
	GasCost uint64             `json:"gasCost"`
	Depth   int                `json:"depth"`
	Error   error              `json:"error,omitempty"`
	Stack   *[]string          `json:"stack,omitempty"`
	Memory  *[]string          `json:"memory,omitempty"`
	Storage *map[string]string `json:"storage,omitempty"`
}
type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}


func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: \t%v [json file]\n",os.Args[0])
		os.Exit(1)
	}
	file,err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening %v: %v\n",os.Args[0],err)
		os.Exit(1)
	}
	defer file.Close()
	data,err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ReadAll() error: %v\n",err)
		os.Exit(1)
	}
	var rpcmsg jsonrpcMessage
	err = json.Unmarshal(data,&rpcmsg)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n",err)
		os.Exit(1)
	}
	var  result ExecutionResult
	err = json.Unmarshal(rpcmsg.Result,&result)
	//fmt.Printf("%+v\n",log_result)
	logs := result.StructLogs
	fmt.Printf("     PC           OP           Gas          Gas cost      Depth     Stack/Memory/Storage\n")
	for i:=0 ; i<len(logs); i++ {
		e:=&logs[i]
		fmt.Printf("%v % 5d\t% 15s\t% 10d\t% 10d\t% 3d\ts=%+v\tm=%+v\tsto=%+v\n",e.Error,e.Pc,e.Op,e.Gas,e.GasCost,e.Depth,e.Stack,e.Memory,e.Storage)
	}

}
