package main

import (
	"fmt"
	"os"
	"errors"
	"net/http"
	"time"
	"sync"
	"math/rand"
	"context"
	"github.com/nsf/termbox-go"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	RWALK_IMAGES_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	RWALK_VIDEOS_URL			string = "https://randomwalknft.s3.us-east-2.amazonaws.com"
	CGAME_IMAGES_URL			string = "https://cosmic-game2.s3.us-east-2.amazonaws.com"
	CGAME_VIDEOS_URL			string = "https://cosmic-game2.s3.us-east-2.amazonaws.com"
	RWALK_THUMBS_URL			string = "https://nfts.cosmicsignature.com/images/randomwalk"
	IMAGE_CHECK_INTERVAL		int    = 900	// 15 minutes in seconds
)
type ImageCheckStatus struct {
	TokenId		int64
	IsPresent	bool
	ErrStr		string
}
type RWalkImageMonitor struct {
	LatestTokens	[3]ImageCheckStatus
	RandomToken		ImageCheckStatus
	DbTokenId		int64
	ContractTokenId	int64
	TokensMatch		bool
	ErrStr			string
}
func fmt_url_addr_for_image_randomwalk(token_id int64) string {

    url := fmt.Sprintf("%v/%06d_black.png",RWALK_IMAGES_URL,token_id)
    return url
}
func fmt_url_addr_for_video_randomwalk(token_id int64) string {
    url := fmt.Sprintf("%v/%06d_black_single.mp4",RWALK_VIDEOS_URL,token_id)
    return url
}
func fmt_url_addr_for_image_cosmicgame(seed string) string {

    url := fmt.Sprintf("%v/0x%s.png",CGAME_IMAGES_URL,seed)
    return url
}
func fmt_url_addr_for_video_cosmicgame(seed string) string {
    url := fmt.Sprintf("%v/0x%s.mp4",CGAME_VIDEOS_URL,seed)
    return url
}
func check_resource(url string) (bool,error) {

    response, err := http.Head(url)
    if err != nil {
        return false, err
    }
	if response.StatusCode != 200 {
		if response.StatusCode == 403 {
			err = errors.New(fmt.Sprintf("Resource %v not found (url=%v) (status = %v)",url,response.StatusCode))
		} else {
			err = errors.New(fmt.Sprintf("Error: HTTP status code = %v for url=%v",response.StatusCode,url))
		}
		return false,err
	}
	return true,nil
}
func get_last_token_id_randomwalk(host,dbname,user,pass string) (int64,error) {

	err,dbobj := pg_connect_db(host,dbname,user,pass)
	if err != nil {
		return -1,err
	}
	defer dbobj.Close()
	var query string
	query = "SELECT token_id FROM rw_mint_evt ORDER BY id DESC LIMIT 1"
	var last_token_id int64
	err = dbobj.QueryRow(query).Scan(&last_token_id)
	if err != nil {
		return -1,err
	}
	return last_token_id,nil
}
func get_last_token_seed_cosmicgame(host,dbname,user,pass string) (string,error) {

	err,dbobj := pg_connect_db(host,dbname,user,pass)
	if err != nil {
		return "",err
	}
	defer dbobj.Close()
	var query string
	query = "SELECT seed FROM cg_mint_event ORDER BY id DESC LIMIT 1"
	var last_token_seed string
	err = dbobj.QueryRow(query).Scan(&last_token_seed)
	if err != nil {
		return "",err
	}
	return last_token_seed,nil
}
func check_randomwalk_resource_availability() {

	token_id,err:=get_last_token_id_randomwalk(
		os.Getenv("IMG_RWALK_MOITOR_HOST"),
		os.Getenv("IMG_RWALK_MONITOR_DBNAME"),
		os.Getenv("IMG_RWALK_MONITOR_USER"),
		os.Getenv("IMG_RWALK_MONITOR_PASS"),
	)
	if err != nil {
		update_global_errors(fmt.Sprintf("Image check err: %v\n",err.Error()))
		return
	}
	url := fmt_url_addr_for_image_randomwalk(token_id)
	success,err := check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
	url = fmt_url_addr_for_video_randomwalk(token_id)
	success,err = check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
}
func check_cosmicgame_resource_availability() {

	seed,err:=get_last_token_seed_cosmicgame(
		os.Getenv("IMG_CGAME_MOITOR_HOST"),
		os.Getenv("IMG_CGAME_MONITOR_DBNAME"),
		os.Getenv("IMG_CGAME_MONITOR_USER"),
		os.Getenv("IMG_CGAME_MONITOR_PASS"),
	)
	if err != nil {
		update_global_errors(fmt.Sprintf("Image check err: %v\n",err.Error()))
		return
	}
	url := fmt_url_addr_for_image_cosmicgame(seed)
	success,err := check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
	url = fmt_url_addr_for_video_cosmicgame(seed)
	success,err = check_resource(url)
	if !success {
		update_global_errors(err.Error())
	}
}

// New functions for monitoring thumbnail images
func fmt_url_addr_for_thumb_randomwalk(token_id int64) string {
	url := fmt.Sprintf("%v/%06d_black_thumb.jpg", RWALK_THUMBS_URL, token_id)
	return url
}

func get_latest_token_ids_randomwalk(host, dbname, user, pass string, limit int) ([]int64, error) {
	err, dbobj := pg_connect_db(host, dbname, user, pass)
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
	
	var token_ids []int64
	for rows.Next() {
		var token_id int64
		err = rows.Scan(&token_id)
		if err != nil {
			return nil, err
		}
		token_ids = append(token_ids, token_id)
	}
	return token_ids, nil
}

func check_thumb_image(token_id int64) (bool, error) {
	url := fmt_url_addr_for_thumb_randomwalk(token_id)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := client.Head(url)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("RWalk image for token %d is not present on image server", token_id))
		return false, err
	}
	return true, nil
}

func get_last_token_from_contract(rpc_url, contract_addr_hex string) (int64, error) {
	// Connect to RPC
	rpcclient, err := rpc.DialContext(context.Background(), rpc_url)
	if err != nil {
		return -1, fmt.Errorf("RPC connection failed: %v", err)
	}
	defer rpcclient.Close()
	
	eclient := ethclient.NewClient(rpcclient)
	
	// Create contract instance
	contract_addr := common.HexToAddress(contract_addr_hex)
	rwalk_contract, err := NewRWalk(contract_addr, eclient)
	if err != nil {
		return -1, fmt.Errorf("Contract instantiation failed: %v", err)
	}
	
	// Call NextTokenId
	opts := &bind.CallOpts{Context: context.Background()}
	next_token_id, err := rwalk_contract.NextTokenId(opts)
	if err != nil {
		return -1, fmt.Errorf("NextTokenId call failed: %v", err)
	}
	
	// Subtract 1 to get the last minted token ID
	last_token_id := next_token_id.Int64() - 1
	return last_token_id, nil
}

func check_rwalk_thumb_images(monitor *RWalkImageMonitor, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Get database connection info from environment
	host := os.Getenv("DB_RWLK_HOST_SRV")
	dbname := os.Getenv("DB_RWLK_DBNAME_SRV")
	user := os.Getenv("DB_RWLK_USER_SRV")
	pass := os.Getenv("DB_RWLK_PASS_SRV")
	
	// Get latest 3 token IDs
	token_ids, err := get_latest_token_ids_randomwalk(host, dbname, user, pass, 3)
	if err != nil {
		for i := 0; i < 3; i++ {
			monitor.LatestTokens[i].TokenId = -1
			monitor.LatestTokens[i].IsPresent = false
			monitor.LatestTokens[i].ErrStr = fmt.Sprintf("DB Error: %v", err)
		}
		monitor.RandomToken.TokenId = -1
		monitor.RandomToken.IsPresent = false
		monitor.RandomToken.ErrStr = fmt.Sprintf("DB Error: %v", err)
		monitor.DbTokenId = -1
		monitor.ContractTokenId = -1
		monitor.TokensMatch = false
		monitor.ErrStr = fmt.Sprintf("DB Error: %v", err)
		update_global_errors(fmt.Sprintf("RWalk image check DB error: %v", err))
		return
	}
	
	// Store DB token ID (the latest one)
	if len(token_ids) > 0 {
		monitor.DbTokenId = token_ids[0]
	} else {
		monitor.DbTokenId = -1
	}
	
	// Get contract token ID
	rpc_url := os.Getenv("RPC1_URL")
	contract_addr := os.Getenv("RWALK_CONTRACT_ADDR")
	contract_token_id, err := get_last_token_from_contract(rpc_url, contract_addr)
	if err != nil {
		monitor.ContractTokenId = -1
		monitor.TokensMatch = false
		monitor.ErrStr = fmt.Sprintf("Contract error: %v", err)
		update_global_errors(fmt.Sprintf("RWalk contract check error: %v", err))
	} else {
		monitor.ContractTokenId = contract_token_id
		monitor.TokensMatch = (monitor.DbTokenId == monitor.ContractTokenId)
		monitor.ErrStr = ""
		
		if !monitor.TokensMatch {
			err_msg := fmt.Sprintf("Token ID mismatch: DB=%d, Contract=%d", monitor.DbTokenId, monitor.ContractTokenId)
			monitor.ErrStr = err_msg
			update_global_errors(err_msg)
		}
	}
	
	// Check each of the latest 3 tokens
	for i := 0; i < len(token_ids) && i < 3; i++ {
		monitor.LatestTokens[i].TokenId = token_ids[i]
		isPresent, err := check_thumb_image(token_ids[i])
		monitor.LatestTokens[i].IsPresent = isPresent
		if err != nil {
			monitor.LatestTokens[i].ErrStr = err.Error()
			update_global_errors(err.Error())
		} else {
			monitor.LatestTokens[i].ErrStr = ""
		}
	}
	
	// Fill remaining slots if fewer than 3 tokens
	for i := len(token_ids); i < 3; i++ {
		monitor.LatestTokens[i].TokenId = -1
		monitor.LatestTokens[i].IsPresent = false
		monitor.LatestTokens[i].ErrStr = ""
	}
	
	// Check random token
	if len(token_ids) > 0 {
		max_token_id := token_ids[0]  // First token is the latest (highest ID)
		if max_token_id > 0 {
			random_token_id := rand.Int63n(max_token_id + 1)  // Random from 0 to max_token_id inclusive
			monitor.RandomToken.TokenId = random_token_id
			isPresent, err := check_thumb_image(random_token_id)
			monitor.RandomToken.IsPresent = isPresent
			if err != nil {
				monitor.RandomToken.ErrStr = err.Error()
				update_global_errors(err.Error())
			} else {
				monitor.RandomToken.ErrStr = ""
			}
		} else {
			monitor.RandomToken.TokenId = -1
			monitor.RandomToken.IsPresent = false
			monitor.RandomToken.ErrStr = ""
		}
	}
}

func print_rwalk_thumb_status(monitor *RWalkImageMonitor) {
	termboxMutex.Lock()
	defer termboxMutex.Unlock()
	
	y := 37
	printAtPosition(0, y, "--------------------- RWalk Thumbnail Images ----------------", termbox.ColorWhite, termbox.ColorDefault)
	
	// Line 1: Print all 4 tokens on a single line
	x := 1
	
	// Print latest 3 tokens
	for i := 0; i < 3; i++ {
		if monitor.LatestTokens[i].TokenId == -1 {
			continue
		}
		
		// Token ID
		line := fmt.Sprintf("%06d:", monitor.LatestTokens[i].TokenId)
		printAtPosition(x, y+1, line, termbox.ColorWhite, termbox.ColorDefault)
		x += len(line)
		
		// Status
		if monitor.LatestTokens[i].IsPresent {
			printAtPosition(x, y+1, "Ok", termbox.ColorGreen, termbox.ColorDefault)
			x += 2
		} else {
			printAtPosition(x, y+1, "Fail", termbox.ColorRed, termbox.ColorDefault)
			x += 4
		}
		
		// Add spacing
		x += 3
	}
	
	// Print random token
	if monitor.RandomToken.TokenId >= 0 {
		line := fmt.Sprintf("Rnd %06d:", monitor.RandomToken.TokenId)
		printAtPosition(x, y+1, line, termbox.ColorCyan, termbox.ColorDefault)
		x += len(line)
		
		if monitor.RandomToken.IsPresent {
			printAtPosition(x, y+1, "Ok", termbox.ColorGreen, termbox.ColorDefault)
		} else {
			printAtPosition(x, y+1, "Fail", termbox.ColorRed, termbox.ColorDefault)
		}
	}
	
	// Line 2: DB/Contract comparison
	x2 := 1
	db_str := fmt.Sprintf("Last token DB: %06d", monitor.DbTokenId)
	printAtPosition(x2, y+2, db_str, termbox.ColorWhite, termbox.ColorDefault)
	x2 += len(db_str) + 2
	
	contract_str := fmt.Sprintf("Last token ctrct: %06d", monitor.ContractTokenId)
	printAtPosition(x2, y+2, contract_str, termbox.ColorWhite, termbox.ColorDefault)
	x2 += len(contract_str) + 2
	
	printAtPosition(x2, y+2, "Match: ", termbox.ColorWhite, termbox.ColorDefault)
	x2 += 7
	
	if monitor.TokensMatch {
		printAtPosition(x2, y+2, "Ok", termbox.ColorGreen, termbox.ColorDefault)
	} else {
		printAtPosition(x2, y+2, "Fail", termbox.ColorRed, termbox.ColorDefault)
	}
	
	termbox.Flush()
}

func monitor_rwalk_thumb_images() {
	var monitor RWalkImageMonitor
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go check_rwalk_thumb_images(&monitor, &wg)
		wg.Wait()
		print_rwalk_thumb_status(&monitor)
		time.Sleep(time.Duration(IMAGE_CHECK_INTERVAL) * time.Second)
	}
}
