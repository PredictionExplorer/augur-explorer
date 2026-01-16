package primitives

import (
	"strconv"
	"fmt"
	"os"
)
func check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Exiting Augur extractor with error: %v",e))
	}
}
func Fatalf(format string, args ...interface{}) {
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
		return fmt.Sprintf("inval_hash: %v",len(long_hash))
	}
	var output string = long_hash[2:8]
	output = output + "…"
	output = output + long_hash[59:65]
	return output
}
func ThousandsFormat(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
func Remove_duplicates_int64(nums []int64) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

