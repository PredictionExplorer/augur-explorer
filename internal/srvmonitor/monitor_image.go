package srvmonitor

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// imageProbeTimeout bounds one HEAD request against the image server.
const imageProbeTimeout = 10 * time.Second

// ImageCheckStatus holds status for a single image check.
type ImageCheckStatus struct {
	TokenID   int64
	IsPresent bool
	ErrStr    string
}

// ImageMonitorData holds all image monitoring data.
type ImageMonitorData struct {
	LatestTokens    [3]ImageCheckStatus
	RandomToken     ImageCheckStatus
	DBTokenID       int64
	ContractTokenID int64
	TokensMatch     bool
	ErrStr          string
}

// ImageMonitor verifies RandomWalk thumbnails: the latest three minted
// tokens plus one random token must be present on the image server, and the
// database's newest mint must match the contract's NextTokenId.
type ImageMonitor struct {
	config   ImageServerConfig
	dbConfig DatabaseConfig
	data     ImageMonitorData
	position Position
	interval time.Duration
	connect  Connector
	client   *http.Client
	// randInt63n picks the random spot-check token; tests pin it.
	randInt63n func(n int64) int64
}

// NewImageMonitor creates a new image monitor.
func NewImageMonitor(imgCfg ImageServerConfig, dbCfg DatabaseConfig, baseY int, iv Intervals) *ImageMonitor {
	return &ImageMonitor{
		config:   imgCfg,
		dbConfig: dbCfg,
		position: Position{X: 0, Y: baseY},
		interval: iv.Image,
		connect:  ConnectPostgres,
		client:   &http.Client{Timeout: imageProbeTimeout},
		// Non-cryptographic sampling of a token id to spot-check the server.
		randInt63n: rand.Int63n,
	}
}

// Name returns the monitor name.
func (m *ImageMonitor) Name() string {
	return "Image Monitor"
}

// Start begins monitoring.
func (m *ImageMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *ImageMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	m.checkImages(ctx, errorChan)
	m.display(disp)
}

// checkImages checks all images and contract.
func (m *ImageMonitor) checkImages(ctx context.Context, errorChan chan<- string) {
	// Get latest 3 token IDs from database
	tokenIDs, err := m.getLatestTokens(ctx, 3)
	if err != nil {
		for i := range 3 {
			m.data.LatestTokens[i] = ImageCheckStatus{TokenID: -1, ErrStr: fmt.Sprintf("DB Error: %v", err)}
		}
		m.data.RandomToken = ImageCheckStatus{TokenID: -1, ErrStr: fmt.Sprintf("DB Error: %v", err)}
		m.data.DBTokenID = -1
		m.data.ContractTokenID = -1
		m.data.TokensMatch = false
		m.data.ErrStr = fmt.Sprintf("DB Error: %v", err)
		sendErr(ctx, errorChan, m.data.ErrStr)
		return
	}

	// Store DB token ID
	if len(tokenIDs) > 0 {
		m.data.DBTokenID = tokenIDs[0]
	} else {
		m.data.DBTokenID = -1
	}

	// Get contract token ID
	contractTokenID, err := m.getContractTokenID(ctx)
	if err != nil {
		m.data.ContractTokenID = -1
		m.data.TokensMatch = false
		m.data.ErrStr = fmt.Sprintf("Contract error: %v", err)
		sendErr(ctx, errorChan, m.data.ErrStr)
	} else {
		m.data.ContractTokenID = contractTokenID
		m.data.TokensMatch = (m.data.DBTokenID == m.data.ContractTokenID)
		m.data.ErrStr = ""

		if !m.data.TokensMatch {
			errMsg := fmt.Sprintf("Token ID mismatch: DB=%d, Contract=%d", m.data.DBTokenID, m.data.ContractTokenID)
			m.data.ErrStr = errMsg
			sendErr(ctx, errorChan, errMsg)
		}
	}

	// Check each of the latest 3 tokens
	for i := 0; i < len(tokenIDs) && i < 3; i++ {
		m.data.LatestTokens[i].TokenID = tokenIDs[i]
		isPresent, err := m.checkImage(ctx, tokenIDs[i])
		m.data.LatestTokens[i].IsPresent = isPresent
		if err != nil {
			m.data.LatestTokens[i].ErrStr = err.Error()
			sendErr(ctx, errorChan, err.Error())
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
			randomTokenID := m.randInt63n(maxTokenID + 1)
			m.data.RandomToken.TokenID = randomTokenID
			isPresent, err := m.checkImage(ctx, randomTokenID)
			m.data.RandomToken.IsPresent = isPresent
			if err != nil {
				m.data.RandomToken.ErrStr = err.Error()
				sendErr(ctx, errorChan, err.Error())
			} else {
				m.data.RandomToken.ErrStr = ""
			}
		}
	}
}

// getLatestTokens gets the latest N token IDs from database.
func (m *ImageMonitor) getLatestTokens(ctx context.Context, limit int) ([]int64, error) {
	dbobj, err := m.connect(ctx, m.dbConfig)
	if err != nil {
		return nil, err
	}
	defer func() { _ = dbobj.Close(ctx) }() // best-effort close of per-check connection

	rows, err := dbobj.Query(ctx, "SELECT token_id FROM rw_mint_evt ORDER BY id DESC LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokenIDs []int64
	for rows.Next() {
		var tokenID int64
		if err := rows.Scan(&tokenID); err != nil {
			return nil, err
		}
		tokenIDs = append(tokenIDs, tokenID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokenIDs, nil
}

// checkImage checks if an image exists.
func (m *ImageMonitor) checkImage(ctx context.Context, tokenID int64) (bool, error) {
	url := fmt.Sprintf("%s/%06d_black_thumb.jpg", m.config.URL, tokenID)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return false, err
	}

	response, err := m.client.Do(req)
	if err != nil {
		return false, err
	}
	_ = response.Body.Close() // HEAD responses carry no body

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("RWalk image for token %d is not present on image server", tokenID)
	}

	return true, nil
}

// getContractTokenID gets the last token ID from the contract.
func (m *ImageMonitor) getContractTokenID(ctx context.Context) (int64, error) {
	rpcClient, err := rpc.DialContext(ctx, m.config.RPCURL)
	if err != nil {
		return -1, fmt.Errorf("RPC connection failed: %w", err)
	}
	defer rpcClient.Close()

	eclient := ethclient.NewClient(rpcClient)
	contractAddr := common.HexToAddress(m.config.ContractAddr)

	rwalkContract, err := randomwalk.NewRWalk(contractAddr, eclient)
	if err != nil {
		return -1, fmt.Errorf("contract instantiation failed: %w", err)
	}

	opts := &bind.CallOpts{Context: ctx}
	nextTokenID, err := rwalkContract.NextTokenId(opts)
	if err != nil {
		return -1, fmt.Errorf("NextTokenId call failed: %w", err)
	}

	// Subtract 1 to get the last minted token ID
	lastTokenID := nextTokenID.Int64() - 1
	return lastTokenID, nil
}

// display renders the image monitoring status.
func (m *ImageMonitor) display(disp Display) {
	y := m.position.Y

	// Header
	disp.DrawText(Position{X: 0, Y: y},
		"--------------------- RWalk Thumbnail Images ----------------",
		ColorWhite, ColorDefault)

	// Line 1: Latest 3 + Random
	x := 1
	for i := range 3 {
		if m.data.LatestTokens[i].TokenID == -1 {
			continue
		}

		// Token ID
		tokenStr := fmt.Sprintf("%06d:", m.data.LatestTokens[i].TokenID)
		disp.DrawText(Position{X: x, Y: y + 1}, tokenStr, ColorWhite, ColorDefault)
		x += len(tokenStr)

		// Status (padded to 4 chars to fully overwrite previous text)
		statusStr := "Ok  "
		statusColor := ColorGreen
		if !m.data.LatestTokens[i].IsPresent {
			statusStr = "Fail"
			statusColor = ColorRed
		}
		disp.DrawText(Position{X: x, Y: y + 1}, statusStr, statusColor, ColorDefault)
		x += 4 + 3 // Always use 4 chars for status + 3 spaces
	}

	// Random token
	if m.data.RandomToken.TokenID >= 0 {
		randomStr := fmt.Sprintf("Rnd %06d:", m.data.RandomToken.TokenID)
		disp.DrawText(Position{X: x, Y: y + 1}, randomStr, ColorCyan, ColorDefault)
		x += len(randomStr)

		// Status (padded to 4 chars to fully overwrite previous text)
		statusStr := "Ok  "
		statusColor := ColorGreen
		if !m.data.RandomToken.IsPresent {
			statusStr = "Fail"
			statusColor = ColorRed
		}
		disp.DrawText(Position{X: x, Y: y + 1}, statusStr, statusColor, ColorDefault)
	}

	// Line 2: DB vs Contract
	x2 := 1
	dbStr := fmt.Sprintf("Last token DB: %06d", m.data.DBTokenID)
	disp.DrawText(Position{X: x2, Y: y + 2}, dbStr, ColorWhite, ColorDefault)
	x2 += len(dbStr) + 2

	contractStr := fmt.Sprintf("Last token ctrct: %06d", m.data.ContractTokenID)
	disp.DrawText(Position{X: x2, Y: y + 2}, contractStr, ColorWhite, ColorDefault)
	x2 += len(contractStr) + 2

	disp.DrawText(Position{X: x2, Y: y + 2}, "Match: ", ColorWhite, ColorDefault)
	x2 += 7

	// Status (padded to 4 chars to fully overwrite previous text)
	matchStr := "Ok  "
	matchColor := ColorGreen
	if !m.data.TokensMatch {
		matchStr = "Fail"
		matchColor = ColorRed
	}
	disp.DrawText(Position{X: x2, Y: y + 2}, matchStr, matchColor, ColorDefault)

	disp.Flush()
}
