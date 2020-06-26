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

	if len(long_addr)!=42 {
		return "inval_addr"
	}
	var output string = long_addr[2:8]
	output = output + "…"
	output = output + long_addr[36:42]
	return output
}
func Short_hash(long_hash string) string {

	if len(long_hash)!=66 {
		return "inval_hash"
	}
	var output string = long_hash[2:8]
	output = output + "…"
	output = output + long_hash[59:65]
	return output
}
