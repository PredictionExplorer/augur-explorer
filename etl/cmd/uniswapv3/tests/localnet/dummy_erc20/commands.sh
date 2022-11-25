~/bin/solc-0.7.0 --combined-json bin,abi,userdoc,devdoc,metadata -o output.json erc20unlimited.sol 
~/bin/abigen-1.10.23 --combined-json combined.json --pkg contracts --type ERC20Unlimited --out ERC20Unlimited.go
