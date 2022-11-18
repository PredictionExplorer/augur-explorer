package main
import (
	"os"
	"fmt"

	 . "github.com/PredictionExplorer/augur-explorer/uevm"
)
func main() {
	mchain,err := OpenMiniChain("/tmp/minichain.dat")
	if err != nil {
		fmt.Printf("Error opening minichain: %v\n",err)
		os.Exit(1)
	}
	r,err:=mchain.ReadLastLine()
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	DumpRecord(&r)
}
