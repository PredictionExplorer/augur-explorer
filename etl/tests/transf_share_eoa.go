/// Makes a transaction to transfer ShareToken to another account, for testing
// This is only for accounts that are EOAs, not valid for Wallet contract accounts

package main

import (
	"os"
	"context"
	"strconv"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	// this address is for local testnet
	SHARE_TOKEN_ADDR = "0xE60c9fe85aEE7B4848a97271dA8c86323CdFb897"
)
var (
	RPC_URL string
	share_token_addr common.Address = common.HexToAddress(SHARE_TOKEN_ADDR)
)

func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC, please set AUGUR_ETH_NODE_RPC_URL env variable : %v\n",err)
	}

	if len(os.Args) < 5  {
		fmt.Printf(
			"usage: \n\t%v [sender_private_key] [recipient_private_key] [market_addr] [outcome]\n" +
			"\tWill transfer all tokens from sender to recipient \n",
			os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long (including 0x prefix)\n")
		os.Exit(1)
	}
	to_pkey_str := os.Args[2]
	if len(to_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long (including 0x prefix)\n")
		os.Exit(1)
	}
	mkt_addr_str := os.Args[3]
	mkt_addr := common.HexToAddress(mkt_addr_str)
	outcome_idx,err := strconv.ParseInt(os.Args[4],10,32)
	if err!=nil {
		Fatalf("Bad integer for 'outcome_idx': %v\n",err)
	}

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		log.Fatal(err)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		Fatalf("Couldn't derive public key for Sender")
	}

	to_PrivateKey, err := crypto.HexToECDSA(to_pkey_str)
	if err != nil {
		log.Fatal(err)
	}
	to_publicKey := to_PrivateKey.Public()
	to_publicKeyECDSA, ok := to_publicKey.(*ecdsa.PublicKey)
	if !ok {
		Fatalf("Couldn't derive public key for Recipient")
	}

	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		log.Fatal(err)
	}
	to_address := crypto.PubkeyToAddress(*to_publicKeyECDSA)
	to_nonce, err := eclient.PendingNonceAt(context.Background(), to_address)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Transferring Share Tokens from %v to %v on Market %v outcome %v\n",
		from_address.String(),to_address.String(),mkt_addr.String(),outcome_idx,
	)

	auth := bind.NewKeyedTransactor(from_PrivateKey)
	auth.Nonce = big.NewInt(int64(from_nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("Sending tx from %v with nonce %v\n",from_address.String(),from_nonce)
	ctrct_share_token,err := NewShareToken(share_token_addr,eclient)
	if err!=nil {
		Fatalf("Failed to instantiate ShareToken contract: %v\n",err)
	}

	big_outcome_idx:=big.NewInt(outcome_idx)
	var copts = new(bind.CallOpts)
	balance,err:=ctrct_share_token.BalanceOfMarketOutcome(copts,mkt_addr,big_outcome_idx,from_address)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	} else {
		fmt.Printf("%v\n",balance.String())
	}

	fmt.Printf("Holder's account has %v tokens of outcome %v\n",balance.String(),outcome_idx)
	zero:=big.NewInt(0)
	if zero.Cmp(balance) == 0 {
		fmt.Printf("Exiting, no share tokens to transfer\n")
		os.Exit(0)
	}
	token_id,err := ctrct_share_token.GetTokenId(copts,mkt_addr,big_outcome_idx)
	if err!=nil {
		fmt.Printf("getTokenId() on ShareToken contract returns error: %v\n",err)
		os.Exit(1)
	}

	fmt.Printf(
		"Checking balance on token_id %v for the purpose of verification of token_id encoding\n",
		hex.EncodeToString(token_id.Bytes()),
	)
	balance2,err:=ctrct_share_token.BalanceOf(copts,from_address,token_id)
	if err!=nil {
		fmt.Printf("Second balance_of() failed with err: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Confirmed balance is %v\n",balance2.String())

	fmt.Printf("Preparing transaction, gas price=%v\n",auth.GasPrice.String())
	tx, err := ctrct_share_token.SetApprovalForAll(auth, to_address, true)
	if err != nil {
		log.Fatal(err)
	}
	_ = tx
	fmt.Printf("Waiting for transaction to be processed..")
	time.Sleep(15 * time.Second)

	approved,err:=ctrct_share_token.IsApprovedForAll(copts,from_address,to_address)
	if err!=nil {
		fmt.Printf("isApprovedForAll() failed with err: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Approval status: %v\n",approved)
	if approved == false {
		fmt.Printf("Couldn't get approval , exiting\n")
		os.Exit(1)
	}

	fmt.Printf("Sending safeTransferFrom tx from %v with nonce %v\n",to_address.String(),to_nonce)
	// Do actual safeTransfeFrom()
	auth2 := bind.NewKeyedTransactor(to_PrivateKey)
	auth2.Nonce = big.NewInt(int64(to_nonce))
	auth2.Value = big.NewInt(0)     // in wei
	auth2.GasLimit = uint64(3000000) // in units
	auth2.GasPrice = gasPrice

	var empty_data []byte
	tx, err = ctrct_share_token.SafeTransferFrom(auth2, from_address,to_address, token_id,balance2,empty_data)
	if err != nil {
		log.Fatal(err)
	}
	_ = tx
	fmt.Printf("Waiting for safeTransferFrom() transaction to be processed\n")
	time.Sleep(15 * time.Second)
}
