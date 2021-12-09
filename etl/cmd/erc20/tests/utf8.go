// Some tokens have invalid unicode characters which prevents them to be inserted in the DB
//		this script tests one of such unicode sequences
package main
import (
	"fmt"
	"os"
	"bytes"
	"encoding/hex"
	"strings"

	 p "github.com/PredictionExplorer/augur-explorer/primitives"
)

const (
	//INVALID_SEQUENCE		string = "532d0000000000000000000000000000000000000000000000000000000000000001"
	INVALID_SEQUENCE		string = "424c4f4b00000000000000000000000000000000000000000000000000000000"
)

func main() {

	sequence,err := hex.DecodeString(INVALID_SEQUENCE)
	if err != nil {
		fmt.Printf("Can't decode: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("err=%v\n",err)
	pos := bytes.Index(sequence,[]byte{0})
	fmt.Printf("pos = %v\n",pos)
	str := p.Bytes_to_string([]byte(sequence))
	fmt.Printf("string is %v\n",str)
	fmt.Printf("hex : %v\n",hex.EncodeToString([]byte(str)))
	fmt.Printf("utf : %v\n",strings.ToValidUTF8(str," "))
}
