#!/bin/bash

rm ./etl/cmd/augur/augur
rm ./etl/cmd/tokens/tokens
rm ./etl/cmd/layer1/layer1
rm ./etl/cmd/layer1/scripts/verify_transactions
rm ./etl/cmd/balancer/balancer
rm ./etl/cmd/uniswap/uniswap
rm ./etl/cmd/ensscan/ensscan
rm ./etl/cmd/ensscan/tools/scanpropnames
rm ./etl/cmd/ensscan/tools/labelhash
rm ./etl/cmd/ensscan/tools/namehash
rm ./etl/cmd/ensscan/tools/lookup
rm ./etl/cmd/ensscan/tools/pk
rm ./etl/cmd/ensscan/tools/check_ens
rm ./etl/cmd/arbitrum/arbitrum
rm ./etl/cmd/augur-turbo/augur-turbo
rm ./etl/cmd/erc20/erc20
rm ./etl/cmd/erc1155/erc1155
rm ./etl/cmd/polysync/polysync
rm ./etl/cmd/polymarkets/polymarkets
rm ./etl/dmesh/dmesh
rm ./etl/tools/augur_blocks
rm ./etl/tools/dai_balances
rm ./etl/tools/toprated
rm ./etl/tools/uniqueaddrs
rm ./etl/tools/gas_usage
rm ./etl/tools/stbalance
rm ./etl/tools/check_owner
rm ./etl/tools/check_wallet
rm ./etl/tools/load_abi
rm ./etl/tools/load_abi_artifacts
rm ./etl/tests/verif_dai_balances
rm ./etl/tests/verif_cash_flow
rm ./etl/ens/scanpropnames
rm ./etl/ens/scantlds
rm ./server/server
rm ./etl/cmd/arbitrum-augur/scripts/cash_faucet
rm ./etl/cmd/arbitrum-augur/scripts/new_sports
rm ./etl/cmd/arbitrum-augur/scripts/new_trusted
rm ./etl/cmd/arbitrum-augur/multicall/deploy
rm ./etl/cmd/arbitrum-augur/tools/balswap
rm ./etl/cmd/arbitrum-augur/tools/erc20bal
rm ./etl/cmd/arbitrum-augur/tools/getmarket
rm ./etl/cmd/arbitrum-augur/tools/approve
rm ./etl/cmd/arbitrum-augur/tools/burnshares
rm ./etl/cmd/arbitrum-augur/tools/dump-vm-trace
rm ./etl/cmd/arbitrum-augur/tools/getcode
rm ./etl/cmd/arbitrum-augur/tools/spotprice
rm ./etl/cmd/arbitrum-augur/tools/trade2price
rm ./etl/cmd/arbitrum-augur/tools/usershares
rm ./etl/cmd/arbitrum-augur/tools/allowance
rm ./etl/cmd/randomwalk/randomwalk
rm ./etl/cmd/notibot/notibot
rm ./etl/cmd/randomwalk/scripts/accept_offer
rm ./etl/cmd/randomwalk/scripts/approve
rm ./etl/cmd/randomwalk/scripts/cancel_offer
rm ./etl/cmd/randomwalk/scripts/mint
rm ./etl/cmd/randomwalk/scripts/new_offer
rm ./etl/cmd/randomwalk/scripts/setname
rm ./etl/cmd/randomwalk/scripts/transfer
rm ./etl/cmd/randomwalk/scripts/ownerof
rm ./etl/cmd/randomwalk/scripts/price
rm ./etl/cmd/randomwalk/scripts/scan_transfers
rm ./etl/cmd/randomwalk/scripts/verify_erc20_transfers
rm ./etl/cmd/randomwalk/scripts/verify_owner
rm ./etl/cmd/randomwalk/tools/discord_bot
rm ./etl/cmd/randomwalk/tools/discord_ch_name
rm ./etl/cmd/randomwalk/tools/discord_user_limit
rm ./etl/cmd/randomwalk/tools/rw_toprated
rm ./etl/cmd/randomwalk/tools/twauthorize
rm ./etl/cmd/randomwalk/tools/tweet_mints
rm ./etl/cmd/randomwalk/tools/twitteroob
rm ./etl/cmd/randomwalk/tools/twsend
rm ./etl/cmd/randomwalk/tools/twsend_image
rm ./etl/cmd/randomwalk/tools/twsend_img_reply
rm ./etl/cmd/randomwalk/tools/ffmpeg-convert
rm ./etl/cmd/randomwalk/scripts/withdrawal
rm ./etl/cmd/randomwalk/scripts/tokenuri
rm ./etl/cmd/randomwalk/rwalkalarm/rwalkalarmd
rm ./etl/cmd/bigstats/bigstats
rm ./etl/cmd/bigstats/scripts/verify_transactions
rm ./etl/cmd/ethprice/ethprice
rm ./etl/cmd/uniswapv3/tests/localnet/test_lastline
rm ./etl/cmd/uniswapv3/tests/localnet/locnewpool
rm ./etl/cmd/uniswapv3/tests/localnet/locmint
rm ./etl/cmd/uniswapv3/tests/localnet/calcaddr
rm ./etl/cmd/uniswapv3/tests/localnet/compute-pool-addr
rm ./etl/cmd/uniswapv3/tests/localnet/decode-price
rm ./etl/cmd/uniswapv3/tests/localnet/delstate
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-dummyerc20
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-getaddr
rm ./etl/cmd/uniswapv3/tests/localnet/dump_rlogs
rm ./etl/cmd/uniswapv3/tests/localnet/dump_statedb
rm ./etl/cmd/uniswapv3/tests/localnet/exact-input
rm ./etl/cmd/uniswapv3/tests/localnet/exact-output
rm ./etl/cmd/uniswapv3/tests/localnet/get-code
rm ./etl/cmd/uniswapv3/tests/localnet/get-factory
rm ./etl/cmd/uniswapv3/tests/localnet/get-pool-init-code
rm ./etl/cmd/uniswapv3/tests/localnet/locevm
rm ./etl/cmd/uniswapv3/tests/localnet/locswap
rm ./etl/cmd/uniswapv3/tests/localnet/initu3
rm ./etl/cmd/uniswapv3/tests/localnet/initminichain
rm ./etl/cmd/uniswapv3/tests/localnet/initlocalevm
rm ./etl/cmd/uniswapv3/tests/localnet/WrappedEth
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-usdc
rm ./etl/cmd/uniswapv3/tests/localnet/send-erc20
rm ./etl/cmd/uniswapv3/tests/localnet/weth_bal
rm ./etl/cmd/uniswapv3/tests/localnet/create-init-pool
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-pinit
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-weth10
rm ./etl/cmd/uniswapv3/tests/localnet/deposit_weth
rm ./etl/cmd/uniswapv3/tests/localnet/erc20bal
rm ./etl/cmd/uniswapv3/tests/localnet/approve
rm ./etl/cmd/uniswapv3/tests/localnet/allowance
rm ./etl/cmd/uniswapv3/tests/localnet/exact_input
rm ./etl/cmd/uniswapv3/tests/localnet/encode_price
rm ./etl/cmd/uniswapv3/tests/localnet/erc20decimals
rm ./etl/cmd/uniswapv3/tests/localnet/mint_liq
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-uv3-factory
rm ./etl/cmd/uniswapv3/tests/localnet/make-pool
rm ./etl/cmd/uniswapv3/tests/localnet/init_pool
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-nftp-descriptor
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-nftp-mgr
rm ./etl/cmd/uniswapv3/tests/localnet/deploy-swaprouter
rm ./etl/cmd/uniswapv3/tests/localnet/mk-position
rm ./etl/cmd/uniswapv3/tests/localnet/dbg_liq
rm ./etl/cmd/uniswapv3/tests/localnet/swap
rm ./etl/cmd/uniswapv3/tests/localnet/dbg_swap
rm ./etl/cmd/uniswapv3/tests/localnet/slot0
rm ./etl/cmd/uniswapv3/uniswapv3
rm ./etl/cmd/uswapv3/uswapv3
rm ./etl/cmd/layer1/layer1
rm ./etl/cmd/biddingwar/biddingwar
rm ./etl/cmd/beacon/beacon
rm ./statserv/statserv
rm ./etl/cmd/balancerv2/balancerv2
