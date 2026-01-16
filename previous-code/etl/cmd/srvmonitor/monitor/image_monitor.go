package monitor

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
	
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/PredictionExplorer/augur-explorer/contracts"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/utils"
)

const (
	ImageCheckInterval = 900 // 15 minutes
)

// ImageCheckStatus holds status for a single image check
type ImageCheckStatus struct {
	TokenID   int64
	IsPresent bool
	ErrStr    string
}

// ImageMonitorData holds all image monitoring data
type ImageMonitorData struct {
	LatestTokens    [3]ImageCheckStatus
	RandomToken     ImageCheckStatus
	DbTokenID       int64
	ContractTokenID int64
	TokensMatch     bool
	ErrStr          string
}

// ImageMonitor monitors RWalk thumbnail images
type ImageMonitor struct {
	config   types.ImageServerConfig
	dbConfig types.DatabaseConfig
	data     ImageMonitorData
	position types.Position
}

// NewImageMonitor creates a new image monitor
func NewImageMonitor(imgCfg types.ImageServerConfig, dbCfg types.DatabaseConfig) *ImageMonitor {
	rand.Seed(time.Now().UnixNano())
	
	return &ImageMonitor{
		config:   imgCfg,
		dbConfig: dbCfg,
		position: types.Position{X: 0, Y: 37},
	}
}

// Name returns the monitor name
func (m *ImageMonitor) Name() string {
	return "Image Monitor"
}

// GetDisplayPosition returns the display position
func (m *ImageMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *ImageMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(ImageCheckInterval * time.Second)
		}
	}
}

// check performs a check cycle
func (m *ImageMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go m.checkImages(&wg, errorChan)
	wg.Wait()
	m.display(disp)
}

// checkImages checks all images and contract
func (m *ImageMonitor) checkImages(wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	// Get latest 3 token IDs from database
	tokenIDs, err := m.getLatestTokens(3)
	if err != nil {
		for i := 0; i < 3; i++ {
			m.data.LatestTokens[i] = ImageCheckStatus{TokenID: -1, ErrStr: fmt.Sprintf("DB Error: %v", err)}
		}
		m.data.RandomToken = ImageCheckStatus{TokenID: -1, ErrStr: fmt.Sprintf("DB Error: %v", err)}
		m.data.DbTokenID = -1
		m.data.ContractTokenID = -1
		m.data.TokensMatch = false
		m.data.ErrStr = fmt.Sprintf("DB Error: %v", err)
		errorChan <- m.data.ErrStr
		return
	}
	
	// Store DB token ID
	if len(tokenIDs) > 0 {
		m.data.DbTokenID = tokenIDs[0]
	} else {
		m.data.DbTokenID = -1
	}
	
	// Get contract token ID
	contractTokenID, err := m.getContractTokenID()
	if err != nil {
		m.data.ContractTokenID = -1
		m.data.TokensMatch = false
		m.data.ErrStr = fmt.Sprintf("Contract error: %v", err)
		errorChan <- m.data.ErrStr
	} else {
		m.data.ContractTokenID = contractTokenID
		m.data.TokensMatch = (m.data.DbTokenID == m.data.ContractTokenID)
		m.data.ErrStr = ""
		
		if !m.data.TokensMatch {
			errMsg := fmt.Sprintf("Token ID mismatch: DB=%d, Contract=%d", m.data.DbTokenID, m.data.ContractTokenID)
			m.data.ErrStr = errMsg
			errorChan <- errMsg
		}
	}
	
	// Check each of the latest 3 tokens
	for i := 0; i < len(tokenIDs) && i < 3; i++ {
		m.data.LatestTokens[i].TokenID = tokenIDs[i]
		isPresent, err := m.checkImage(tokenIDs[i])
		m.data.LatestTokens[i].IsPresent = isPresent
		if err != nil {
			m.data.LatestTokens[i].ErrStr = err.Error()
			errorChan <- err.Error()
		} else {
			m.data.LatestTokens[i].ErrStr = ""
		}
	}
	
	// Fill remaining slots
	for i := len(tokenIDs); i < 3; i++ {
		m.data.LatestTokens[i] = ImageCheckStatus{TokenID: -1}
	}
	
	// Check random token
	if len(tokenIDs) > 0 {
		maxTokenID := tokenIDs[0]
		if maxTokenID > 0 {
			randomTokenID := rand.Int63n(maxTokenID + 1)
			m.data.RandomToken.TokenID = randomTokenID
			isPresent, err := m.checkImage(randomTokenID)
			m.data.RandomToken.IsPresent = isPresent
			if err != nil {
				m.data.RandomToken.ErrStr = err.Error()
				errorChan <- err.Error()
			} else {
				m.data.RandomToken.ErrStr = ""
			}
		}
	}
}

// getLatestTokens gets the latest N token IDs from database
func (m *ImageMonitor) getLatestTokens(limit int) ([]int64, error) {
	dbobj, err := utils.ConnectPostgres(m.dbConfig)
	if err != nil {
		return nil, err
	}
	defer dbobj.Close()
	
	query := "SELECT token_id FROM rw_mint_evt ORDER BY id DESC LIMIT $1"
	rows, err := dbobj.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var tokenIDs []int64
	for rows.Next() {
		var tokenID int64
		err = rows.Scan(&tokenID)
		if err != nil {
			return nil, err
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	
	return tokenIDs, nil
}

// checkImage checks if an image exists
func (m *ImageMonitor) checkImage(tokenID int64) (bool, error) {
	url := fmt.Sprintf("%s/%06d_black_thumb.jpg", m.config.URL, tokenID)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	
	response, err := client.Head(url)
	if err != nil {
		return false, err
	}
	
	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("RWalk image for token %d is not present on image server", tokenID))
		return false, err
	}
	
	return true, nil
}

// getContractTokenID gets the last token ID from the contract
func (m *ImageMonitor) getContractTokenID() (int64, error) {
	rpcClient, err := rpc.DialContext(context.Background(), m.config.RPCURL)
	if err != nil {
		return -1, fmt.Errorf("RPC connection failed: %v", err)
	}
	defer rpcClient.Close()
	
	eclient := ethclient.NewClient(rpcClient)
	contractAddr := common.HexToAddress(m.config.ContractAddr)
	
	rwalkContract, err := contracts.NewRWalk(contractAddr, eclient)
	if err != nil {
		return -1, fmt.Errorf("Contract instantiation failed: %v", err)
	}
	
	opts := &bind.CallOpts{Context: context.Background()}
	nextTokenID, err := rwalkContract.NextTokenId(opts)
	if err != nil {
		return -1, fmt.Errorf("NextTokenId call failed: %v", err)
	}
	
	// Subtract 1 to get the last minted token ID
	lastTokenID := nextTokenID.Int64() - 1
	return lastTokenID, nil
}

// display renders the image monitoring status
func (m *ImageMonitor) display(disp display.Display) {
	y := 37
	
	// Header
	disp.DrawText(types.Position{X: 0, Y: y},
		"--------------------- RWalk Thumbnail Images ----------------",
		types.ColorWhite, types.ColorDefault)
	
	// Line 1: Latest 3 + Random
	x := 1
	for i := 0; i < 3; i++ {
		if m.data.LatestTokens[i].TokenID == -1 {
			continue
		}
		
		// Token ID
		tokenStr := fmt.Sprintf("%06d:", m.data.LatestTokens[i].TokenID)
		disp.DrawText(types.Position{X: x, Y: y + 1}, tokenStr, types.ColorWhite, types.ColorDefault)
		x += len(tokenStr)
		
		// Status (padded to 4 chars to fully overwrite previous text)
		statusStr := "Ok  "
		statusColor := types.ColorGreen
		if !m.data.LatestTokens[i].IsPresent {
			statusStr = "Fail"
			statusColor = types.ColorRed
		}
		disp.DrawText(types.Position{X: x, Y: y + 1}, statusStr, statusColor, types.ColorDefault)
		x += 4 + 3  // Always use 4 chars for status + 3 spaces
	}
	
	// Random token
	if m.data.RandomToken.TokenID >= 0 {
		randomStr := fmt.Sprintf("Rnd %06d:", m.data.RandomToken.TokenID)
		disp.DrawText(types.Position{X: x, Y: y + 1}, randomStr, types.ColorCyan, types.ColorDefault)
		x += len(randomStr)
		
		// Status (padded to 4 chars to fully overwrite previous text)
		statusStr := "Ok  "
		statusColor := types.ColorGreen
		if !m.data.RandomToken.IsPresent {
			statusStr = "Fail"
			statusColor = types.ColorRed
		}
		disp.DrawText(types.Position{X: x, Y: y + 1}, statusStr, statusColor, types.ColorDefault)
	}
	
	// Line 2: DB vs Contract
	x2 := 1
	dbStr := fmt.Sprintf("Last token DB: %06d", m.data.DbTokenID)
	disp.DrawText(types.Position{X: x2, Y: y + 2}, dbStr, types.ColorWhite, types.ColorDefault)
	x2 += len(dbStr) + 2
	
	contractStr := fmt.Sprintf("Last token ctrct: %06d", m.data.ContractTokenID)
	disp.DrawText(types.Position{X: x2, Y: y + 2}, contractStr, types.ColorWhite, types.ColorDefault)
	x2 += len(contractStr) + 2
	
	disp.DrawText(types.Position{X: x2, Y: y + 2}, "Match: ", types.ColorWhite, types.ColorDefault)
	x2 += 7
	
	// Status (padded to 4 chars to fully overwrite previous text)
	matchStr := "Ok  "
	matchColor := types.ColorGreen
	if !m.data.TokensMatch {
		matchStr = "Fail"
		matchColor = types.ColorRed
	}
	disp.DrawText(types.Position{X: x2, Y: y + 2}, matchStr, matchColor, types.ColorDefault)
	
	disp.Flush()
}




