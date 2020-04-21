## Augur Data Extractor

Extracts data from Augur Prediction Marketplace (http://augur.net) and stores it in an SQL database.

### Dependencies

 * Latest golang `crypto` package containing `NewLegacyKeccak256` function
 * Latest Ethereum `go-ethereum` package with Tuple ABI Unpack function (v1.9.13)
 * PostgreSQL libpq package

### Compile & Run

	go ./
	#set Ethereum Node RPC port (TODO)
	. aux/dev-db.env ;#set datasource here
	./augur-extractor

### Specifying database location

	cat aux/dev-db.env

### Database initialization

	psql [dbname] < sql/tables.sql

### Database Schema and documentation

	cat sql/tables.sql
