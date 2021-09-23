// Will balancer-swap (buy or sell operation) a token until it reaches a desired price level
package main

import (
	"os"
	"fmt"
	"time"
	"math/big"
	"strconv"
	"strings"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CHAIN_ID		int64 = 80001
	SLEEP_TIME		int64 =180 // seconds
	SLEEP_TIME_RAND		int64 = 100 // max seconds for random sleep
	MAX_RECEIPT_WAIT_TIME int64 = 120 // seconds to wait for transaction receipt before considering timeout
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 9 {
		fmt.Printf(
			"Usage: \n\t\t%v "+
			"[priv_key] [amm_factory_addr] [market_factory] "+
			"[market_id] [outcome_in] [outcome_out] "+
			"[target_price] [amount]\n\n\t\t"+
			"Swaps outcome1 share tokens for outcome2\n",
			os.Args[0],
		)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long (including 0x prefix)\n")
		os.Exit(1)
	}

	amm_factory_addr := common.HexToAddress(os.Args[2])
	market_factory_addr := common.HexToAddress(os.Args[3])
	market_id,err := strconv.ParseInt(os.Args[4],10,64)
	if err != nil {
		fmt.Printf("Error parsing outcome1 field: %v\n",err)
		os.Exit(1)
	}
	outc1,err := strconv.ParseInt(os.Args[5],10,64)
	if err != nil {
		fmt.Printf("Error parsing outcome1 field: %v\n",err)
		os.Exit(1)
	}
	outc2,err := strconv.ParseInt(os.Args[6],10,64)
	if err != nil {
		fmt.Printf("Error parsing outcome1 field: %v\n",err)
		os.Exit(1)
	}
	target_price_str := os.Args[7]
	amount_str := os.Args[8]
	target_price := big.NewInt(0)
	_,success := target_price.SetString(target_price_str,10)
	if !success {
		fmt.Printf("Bad integer for target_price parameter\n")
		os.Exit(1)
	}
	big_amount := big.NewInt(0)
	_,success = big_amount.SetString(amount_str,10)
	if !success {
		fmt.Printf("Bad integer for pirce increment parameter\n")
		os.Exit(1)
	}

	amm_factory,err := NewAMMFactory(amm_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate AMMFactory contract: %v\n",err)
		os.Exit(1)
	}

	market_factory,err := NewSportsLinkMarketFactory(market_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Market Factory contract: %v\n",err)
		os.Exit(1)
	}
	sharefactor,err := market_factory.ShareFactor(copts)
	if err!=nil {
		fmt.Printf("Error getting Share factor: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("ShareFactor = %v\n",sharefactor.String())

	big_market_id := big.NewInt(market_id)
	market_obj,err:=market_factory.GetMarket(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	if outc1 >= int64(len(market_obj.ShareTokens)) {
		fmt.Printf("Outcome 1 is larger than the size of outcomes (%v)\n",outc1)
		os.Exit(1)
	}
	if outc2 >= int64(len(market_obj.ShareTokens)) {
		fmt.Printf("Outcome 2 is larger than the size of outcomes (%v)\n",outc2)
		os.Exit(1)
	}

	token1 := market_obj.ShareTokens[outc1]
	token2 := market_obj.ShareTokens[outc2]

	erc20_token1,err := NewOwnedERC20(token1,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract for %v: %v\n",token1.String(),err)
		os.Exit(1)
	}
	symbol1,err:=erc20_token1.Symbol(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	erc20_token2,err := NewOwnedERC20(token2,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract for %v: %v\n",token2.String(),err)
		os.Exit(1)
	}
	symbol2,err:=erc20_token2.Symbol(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}
	decimals,err:=erc20_token1.Decimals(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	bpool_addr,err := amm_factory.Pools(copts,market_factory_addr,big_market_id)
	if err!=nil {
		fmt.Printf("Failed to retrieve Balancer Pool (BPool) contract from market factory: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Params:\n")
	fmt.Printf("\tAMM factory:    %v\n",amm_factory_addr.String())
	fmt.Printf("\tMarket factory: %v\n",market_factory_addr.String())
	fmt.Printf("\tMarket ID:      %v\n",market_id)
	fmt.Printf("\tOutcome IN:     %v (%v)\n",outc1,symbol1)
	fmt.Printf("\tOutcome OUT:    %v (%v)\n",outc2,symbol2)
	fmt.Printf("\n")

	fmt.Printf("Pool addr: %v\n",bpool_addr.String())
	fmt.Printf("Market {\n")
	fmt.Printf("\tSettlement Address: %v\n",market_obj.SettlementAddress.String())
	fmt.Printf("\tShareTokens:\n")
	for i:=0; i<len(market_obj.ShareTokens) ; i++ {
		fmt.Printf("\t\t%v\n",market_obj.ShareTokens[i].String())
	}
	fmt.Printf("\tEndTime: %v\n",market_obj.EndTime.String())
	fmt.Printf("\tWinner: %v\n",market_obj.Winner.String())
	fmt.Printf("\tSettlement Fee: %v\n",market_obj.SettlementFee.String())
	fmt.Printf("\tProtocol Fee: %v\n",market_obj.ProtocolFee.String())
	fmt.Printf("\tStakerFee: %v\n",market_obj.StakerFee.String())
	fmt.Printf("\tCreation Timestamp: %v\n",market_obj.CreationTimestamp.String())
	fmt.Printf("}\n")

	bpool,err := NewBPool(bpool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate BPool contract: %v\n",err)
		os.Exit(1)
	}

	spot_price,err:=bpool.GetSpotPrice(copts,token1,token2)
	if err!=nil {
		fmt.Printf("Error during GetSpotPrice() call: %v\n",err)
		os.Exit(1)
	}

	divisor:=big.NewInt(0)
	if decimals == 0 {
		divisor = big.NewInt(1)	//to avoid divide by 0 error
	} else {
		multiplier_str := strings.Repeat("0",int(decimals))
		multiplier_str = "1" + multiplier_str
		divisor.SetString(multiplier_str,10)
	}
	fmt.Printf("decimals=%v, divisor=%v\n",decimals,divisor.String())

	diff := big.NewInt(0)
	diff.Sub(target_price,spot_price)
	sign := diff.Sign()
	if sign == 0 {
		fmt.Printf("Swap price is already the same as target price\n")
		os.Exit(2)
	} else {
		if sign < 0 {
			fmt.Printf(
				"Target price is lower than current price. "+
				"This script can only make price higher.",
			)
			os.Exit(2)
		}
	}

	compact_price:= big.NewInt(0)
	reminder := big.NewInt(0)
	compact_price.QuoRem(spot_price,divisor,reminder)
	fmt.Printf("Spot price:\n")
	fmt.Printf(
		"Swap price of \n\t%v (outcome %v (%v)) for \n\t%v (outcome %v (%v)) = %v\n",
		token1.String(),outc1,symbol1,
		token2.String(),outc2,symbol2,
		spot_price.String(),
	)
	fmt.Printf(
		"Price with floating point: %v.%018s \n",compact_price.String(),reminder.String(),
	)

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Sprintf("Error making private key: %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)

	for {
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
		big_chain_id := big.NewInt(CHAIN_ID)
		fmt.Printf("Using chain_id=%v\n",big_chain_id.String())
		txopts := bind.NewKeyedTransactor(from_PrivateKey)
		txopts.Nonce = big.NewInt(int64(from_nonce))
		txopts.Value = big.NewInt(0)     // in wei
		txopts.GasLimit = uint64(10000000) // in units
		txopts.GasPrice = gasPrice
		fmt.Printf("Gas price = %v\n",gasPrice.String())

		signfunc := func(signer_disabled types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewEIP155Signer(big_chain_id)
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_PrivateKey)
			if err != nil {
				fmt.Printf("Error signing: %v\n",err)
				os.Exit(1)
				return nil,nil
			}
			return tx.WithSignature(signer, signature)
		}
		txopts.Signer = signfunc

		max_price := big.NewInt(0)
		max_price.SetString("1000000000000000000000000000000",10)
		amount_out := big.NewInt(1)
		big_amount := big.NewInt(0)
		big_amount.SetString(amount_str,10)
		fmt.Printf("Swapping %v %v for %v (token out)\n",big_amount.String(),symbol1,symbol2)
		fmt.Printf("Token In: %v\n",token1.String())
		fmt.Printf("Token Out: %v\n",token2.String())
		fmt.Printf("Amount In: %v\n",big_amount.String())
		fmt.Printf("Amount Out (min): %v\n",amount_out.String())
		fmt.Printf("Max price: %v\n",max_price.String())
		tx,err := bpool.SwapExactAmountIn(txopts,
			token1,
			big_amount,
			token2,
			amount_out,
			max_price,
		)
		if err!=nil {
			fmt.Printf("Error sending tx: %v\n",err)
			os.Exit(1)
		}
		fmt.Printf("Tx hash = %v\n",tx.Hash().String())

		time_counter := int64(0)
		for {
			time.Sleep(1 * time.Second)
			receipt,err := eclient.TransactionReceipt(context.Background(),tx.Hash())
			if err != nil {
				if err == ethereum.NotFound {
					fmt.Printf("Wating ...\n")
				} else {
					fmt.Printf("Error getting tx receipt: %v\n",err)
					os.Exit(1)
				}
			}
			if receipt != nil {
				if receipt.Status == types.ReceiptStatusSuccessful {
					fmt.Printf("Transaction executed\n")
					break
				} else {
					fmt.Printf("Transaction failed.")
					os.Exit(1)
				}
			}
			time_counter = time_counter + 1
			if time_counter >= MAX_RECEIPT_WAIT_TIME {
				fmt.Printf("Can't get receipt for too long, exiting\n")
				os.Exit(1)
			}
		}

		spot_price,err:=bpool.GetSpotPrice(copts,token1,token2)
		if err!=nil {
			fmt.Printf("Error during GetSpotPrice() call: %v\n",err)
			os.Exit(1)
		}

		divisor:=big.NewInt(0)
		if decimals == 0 {
			divisor = big.NewInt(1)	//to avoid divide by 0 error
		} else {
			multiplier_str := strings.Repeat("0",int(decimals))
			multiplier_str = "1" + multiplier_str
			divisor.SetString(multiplier_str,10)
		}

		reminder := big.NewInt(0)
		compact_price:= big.NewInt(0)
		compact_price.QuoRem(spot_price,divisor,reminder)
		fmt.Printf("Spot price:\n")
		fmt.Printf(
			"Swap price of \n\t%v (outcome %v (%v)) for \n\t%v (outcome %v (%v)) = %v\n",
			token1.String(),outc1,symbol1,
			token2.String(),outc2,symbol2,
			spot_price.String(),
		)
		fmt.Printf(
			"Price with floating point: %v.%018s \n",compact_price.String(),reminder.String(),
		)

		diff := big.NewInt(0)
		diff.Sub(target_price,spot_price)
		sign := diff.Sign()
		if sign == 0 {
			fmt.Printf("Swap price is already the same as target price\n")
			os.Exit(2)
		} else {
			if sign < 0 {
				fmt.Printf(
					"Target price reached\n",
				)
				os.Exit(2)
			}
		}

	}// end for sending transactions for swaps
}
