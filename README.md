## Augur data Extractor

Extracts data from Augur Prediction Marketplace (http://augur.net) and stores it in an SQL database.

### Dependencies

 * Latest golang `crypto` package containing `NewLegacyKeccak256` function
 * Latest Ethereum `go-ethereum` package with Tuple ABI Unpack function (v1.9.13)
 * PostgreSQL libpq package

### Compile

	go build main.go types.go db.go utils.go

### Specifying database location

	cat aux/dev-db.env

### Database initialization

	psql [dbname] < sql/tables.sql
