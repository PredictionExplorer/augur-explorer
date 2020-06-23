package primitives

import (
	//"runtime"
	"fmt"
	"os"
	//"io"
)
func check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Exiting Augur extractor with error: %v",e))
	}
}
func Fatalf(format string, args ...interface{}) {
	/*
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	*/
	fmt.Printf(fmt.Sprintf("Fatal: "+format+"\n", args...))
	os.Exit(1)
}
func Short_address(long_addr string) string {

	var output string = long_addr[3:9] 
	output = output + "…"
	output = output + long_addr[25:31]
	return output
}
