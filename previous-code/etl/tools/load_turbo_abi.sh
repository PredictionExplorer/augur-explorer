#!/bin/bash

EXECUTABLE=$1

if test -z "$EXECUTABLE" 
then
	echo "usage: [load_abi_binary]"
	exit 1
fi
echo Loading AMM Module contract ABIs ...
$EXECUTABLE ./abi/contracts/chainlink/TheRundownChainlink.sol/TheRundownChainlink.json TheRundownChainlink
$EXECUTABLE ./abi/contracts/balancer/BToken.sol/BToken.json BToken
$EXECUTABLE ./abi/contracts/balancer/BToken.sol/BTokenBase.json BTokenBase
$EXECUTABLE ./abi/contracts/balancer/BNum.sol/BNum.json BNum
$EXECUTABLE ./abi/contracts/balancer/BPool.sol/BPool.json BPool
$EXECUTABLE ./abi/contracts/balancer/BMath.sol/BMath.json BMath
$EXECUTABLE ./abi/contracts/balancer/BColor.sol/BBronze.json BBronze
$EXECUTABLE ./abi/contracts/balancer/BColor.sol/BColor.json BColor
$EXECUTABLE ./abi/contracts/balancer/BConst.sol/BConst.json BConst
$EXECUTABLE ./abi/contracts/balancer/BFactory.sol/BFactory.json BFactory
$EXECUTABLE ./abi/contracts/turbo/AbstractMarketFactory.sol/AbstractMarketFactory.json AbstractMarketFactory
$EXECUTABLE ./abi/contracts/turbo/AMMFactory.sol/AMMFactory.json AMMFactory
$EXECUTABLE ./abi/contracts/turbo/SportsLinkMarketFactory.sol/SportsLinkMarketFactory.json SportsLinkMarketFactory
$EXECUTABLE ./abi/contracts/turbo/PriceMarketFactory.sol/TestPriceMarketFactory.json TestPriceMarketFactory
$EXECUTABLE ./abi/contracts/turbo/MMALinkMarketFactory.sol/MMALinkMarketFactory.json MMALinkMarketFactory
$EXECUTABLE ./abi/contracts/turbo/FeePot.sol/FeePot.json
$EXECUTABLE ./abi/contracts/turbo/SportsLinkProxy.sol/SportsLinkProxy.json SportsLinkProxy
$EXECUTABLE ./abi/contracts/turbo/TrustedMarketFactory.sol/TrustedMarketFactory.json TrustedMarketFactory
echo Done
