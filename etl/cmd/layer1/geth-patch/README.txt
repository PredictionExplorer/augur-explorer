In order to keep with Polygon's Network speed we add a new RPC method to Geth's RPC API to retrieve all receipts of a block in single request. This accelerates receipt fetching by the number of transactions in the block, on average 200x. Therefore per block we now only have 2 calls: getBlock() and getBlockReceipts()

To apply the patch:
	cd go-ethereum-1.10.6
	patch -p1 < block-receipts-patch.diff

New RPC request will be available at 'eth' API as :

	eth_getBlockReceipts

Example:
	curl -X POST -H "Content-Type: application/json" \
	--data '{"jsonrpc":"2.0","method":"eth_getBlockReceipts", \
	"params":["0x3c9c46a46b17361cd1ac3ed3401c9a268095c1810bf991c470c139f8441e1d0b"],"id":67}' \
	http://host:port

For Polygon you have to apply additional patch because Polygon adds one transaction of its own at the end of the block, and the receipt for this transaction is stored separately inside the DB. This patch appends Bor's receipt at the end of the list:

	matic-bor-patch.diff

