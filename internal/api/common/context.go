// Package common provides shared middleware, response envelopes and request
// helpers for the RWCG web server.
package common

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// IsAddressValid validates an Ethereum address and returns the checksummed version.
// The jsonOutput flag used to select between JSON and HTML error rendering;
// HTML pages are gone, so both branches now emit the same JSON error.
func IsAddressValid(c *httpx.Context, jsonOutput bool, addr string) (string, bool) {
	if (len(addr) != 40) && (len(addr) != 42) {
		var errMsg = fmt.Sprintf("Provided address has invalid length (len=%v)", len(addr))
		c.JSON(http.StatusOK, httpx.H{
			"status": 0,
			"error":  errMsg,
		})
		return "", false
	}
	if (addr[0] == '0') && (addr[1] == 'x') {
		addr = addr[2:]
	}
	if len(addr) != 40 {
		c.JSON(http.StatusOK, httpx.H{
			"status": 0,
			"error":  "Invalid address length",
		})
		return "", false
	}
	var formattedAddr string
	addrBytes, err := hex.DecodeString(addr)
	if err == nil {
		addr := common.BytesToAddress(addrBytes)
		formattedAddr = addr.String()
	} else {
		c.JSON(http.StatusOK, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("Provided address parameter is invalid : %v", err),
		})
		return "", false
	}
	return formattedAddr, true
}
