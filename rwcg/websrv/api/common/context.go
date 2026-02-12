// Package common provides shared utilities and context for the RWCG web server
package common

import (
	"log"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
)

// ServerContext holds shared dependencies for all handlers
type ServerContext struct {
	Db        *SQLStorage
	EthClient *ethclient.Client
	Info      *log.Logger
	Error     *log.Logger
}

var (
	// Ctx is the global server context, initialized by InitContext
	Ctx *ServerContext
)

// InitContext initializes the global server context
func InitContext(db *SQLStorage, ethClient *ethclient.Client, info, errorLog *log.Logger) {
	Ctx = &ServerContext{
		Db:        db,
		EthClient: ethClient,
		Info:      info,
		Error:     errorLog,
	}
}

// IsAddressValid validates an Ethereum address and returns the checksummed version
func IsAddressValid(c *gin.Context, jsonOutput bool, addr string) (string, bool) {
	if (len(addr) != 40) && (len(addr) != 42) {
		var errMsg = fmt.Sprintf("Provided address has invalid length (len=%v)", len(addr))
		if jsonOutput {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"error":  errMsg,
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "RWCG: Error",
				"ErrDescr": errMsg,
			})
		}
		return "", false
	}
	if (addr[0] == '0') && (addr[1] == 'x') {
		addr = addr[2:]
	}
	if len(addr) != 40 {
		if jsonOutput {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"error":  fmt.Sprintf("Invalid address length"),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "RWCG: Error",
				"ErrDescr": fmt.Sprintf("Invalid address length"),
			})
		}
		return "", false
	}
	var formattedAddr string
	addrBytes, err := hex.DecodeString(addr)
	if err == nil {
		addr := common.BytesToAddress(addrBytes)
		formattedAddr = addr.String()
	} else {
		if jsonOutput {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"error":  fmt.Sprintf("Provided address parameter is invalid : %v", err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "RWCG: Error",
				"ErrDescr": fmt.Sprintf("Provided address parameter is invalid : %v", err),
			})
		}
		return "", false
	}
	return formattedAddr, true
}

