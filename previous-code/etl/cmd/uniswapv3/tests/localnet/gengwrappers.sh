#!/bin/bash
BINARY=$HOME/bin/abigen-1.10.23
if [ -f "$BINARY" ]
then
	echo generating Golang wrappers
else
	echo $BINARY binary doesn\'t exist , install abigen \(from go-ethereum/cmd/abigen in Geth sources\)
	exit 1
fi
SYM="UniswapV3Factory"
LSYM="Local$SYM"
$BINARY --combined-json ./$SYM/combined.json --pkg main --type $LSYM --out "$LSYM.go"
echo $LSYM.go has been created, move it to destination with \'mv\'
SYM="UniswapV3Pool"
LSYM="Local$SYM"
$BINARY --combined-json ./$SYM/combined.json --pkg main --type $LSYM --out "$LSYM.go"
echo $LSYM.go has been created, move it to destination with \'mv\'

SYM="NonfungiblePositionManager"
LSYM="Local$SYM"
$BINARY --combined-json ./$SYM/combined.json --pkg main --type $LSYM --out "$LSYM.go"
echo $LSYM.go has been created, move it to destination with \'mv\'

SYM="NonfungibleTokenPositionDescriptor"
LSYM="Local$SYM"
$BINARY --combined-json ./$SYM/combined.json --pkg main --type $LSYM --out "$LSYM.go"
echo $LSYM.go has been created, move it to destination with \'mv\'

SYM="SwapRouter"
LSYM="Local$SYM"
$BINARY --combined-json ./$SYM/combined.json --pkg main --type $LSYM --out "$LSYM.go"
echo $LSYM.go has been created, move it to destination with \'mv\'
