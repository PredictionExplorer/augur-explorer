// Sets initialDurationUntilMainPrizeDivisor so that after first bid,
// time until main prize equals the specified duration (in seconds)
package main

import (
	"os"
	"fmt"
	"math/big"
	"context"
	"crypto/ecdsa"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
var (
	RPC_URL    string
	chain_id   *big.Int
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}
	chain_id,err = eclient.ChainID(context.Background())
	if err != nil {
		fmt.Printf("Error getting chain id : %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 4 {
		fmt.Printf("Usage: \n\t\t%v [priv_key] [contract_addr] [desired_duration_seconds]\n\n"+
			"\t\tSets initialDurationUntilMainPrizeDivisor so that after first bid,\n"+
			"\t\tthe time until main prize equals desired_duration_seconds.\n"+
			"\t\tExample: %v [key] [addr] 300  (for 5 minutes)\n",os.Args[0],os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}

	cosmic_game_addr := common.HexToAddress(os.Args[2])

	desired_seconds, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Printf("Error parsing desired_duration_seconds: %v\n", err)
		os.Exit(1)
	}
	if desired_seconds <= 0 {
		fmt.Printf("desired_duration_seconds must be positive\n")
		os.Exit(1)
	}

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	// Read current mainPrizeTimeIncrementInMicroSeconds from contract
	var copts bind.CallOpts
	mainPrizeTimeIncrementInMicroSeconds, err := cosmic_game_ctrct.MainPrizeTimeIncrementInMicroSeconds(&copts)
	if err != nil {
		fmt.Printf("Error reading mainPrizeTimeIncrementInMicroSeconds: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("mainPrizeTimeIncrementInMicroSeconds = %v\n", mainPrizeTimeIncrementInMicroSeconds.String())

	// Calculate divisor: initialDuration = mainPrizeTimeIncrementInMicroSeconds / divisor
	// So: divisor = mainPrizeTimeIncrementInMicroSeconds / desired_seconds
	// With rounding: divisor = (mainPrizeTimeIncrementInMicroSeconds + desired/2) / desired
	desired := big.NewInt(desired_seconds)
	half_desired := new(big.Int).Div(desired, big.NewInt(2))
	numerator := new(big.Int).Add(mainPrizeTimeIncrementInMicroSeconds, half_desired)
	divisor := new(big.Int).Div(numerator, desired)
	
	fmt.Printf("\nFormula: initialDuration = mainPrizeTimeIncrementInMicroSeconds / divisor\n")
	fmt.Printf("         %v = %v / %v\n", desired_seconds, mainPrizeTimeIncrementInMicroSeconds.String(), divisor.String())

	fmt.Printf("Desired initial duration: %v seconds\n", desired_seconds)
	fmt.Printf("Calculated divisor: %v\n", divisor.String())

	// Verify: actual_duration = mainPrizeTimeIncrementInMicroSeconds / divisor
	actual_duration := new(big.Int).Div(mainPrizeTimeIncrementInMicroSeconds, divisor)
	fmt.Printf("Actual initial duration will be: %v seconds\n", actual_duration.String())

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Printf("Error making private key: %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender\n")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)

	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting account's nonce: %v\n",err)
		os.Exit(1)
	}
	gasPrice, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("Using chain_id=%v\n",chain_id.String())
	fmt.Printf("From address: %v\n", from_address.String())

	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)
	txopts.GasLimit = uint64(100000)
	txopts.GasPrice = gasPrice

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chain_id)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_PrivateKey)
		if err != nil {
			fmt.Printf("Error signing: %v\n",err)
			os.Exit(1)
			return nil,nil
		}
		return tx.WithSignature(signer, signature)
	}
	txopts.Signer = signfunc

	fmt.Printf("Setting initialDurationUntilMainPrizeDivisor to %v...\n", divisor.String())
	tx,err := cosmic_game_ctrct.SetInitialDurationUntilMainPrizeDivisor(txopts, divisor)
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	fmt.Printf("Done! After first bid, time until main prize will be ~%v seconds.\n", actual_duration.String())
}

