// Database Storage package
package dbs

// Notes:
// 		the functions contained in this package should be only storing objects defined in 
//		primitves::types.go , thus having ver low amount of business logic of the application
//		(though, some exceptions may apply)

import (
	"fmt"
	"net"
	"os"
	"log"
	"math/big"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	zero *big.Int = big.NewInt(0)
	hundred *big.Int = big.NewInt(100)
)
type SQLStorage struct {
	db					*sql.DB
	db_logger			*log.Logger
	Info				*log.Logger
	mkt_order_id_ptr	*int64		// global var indicating we have an OrderEvent going on in event chain
}
func (ss *SQLStorage) Db() *sql.DB {
	return ss.db
}
func New_sql_storage(mkt_order_ptr *int64,info_log *log.Logger,db_log *log.Logger,host_port,db_name,user,pass string) *SQLStorage {
	var err error
	host,port,err:=net.SplitHostPort(host_port)
	if (err!=nil) {
		host=host_port
		port="5432"
	}
	conn_str := "user='"+user+"' dbname='" + db_name + "' password='" + pass +
				"' host='" + host + "' port='" + port +	"'";
	db,err := sql.Open("postgres",conn_str);
	if (err!=nil) {
		info_log.Printf("Error connecting: %v\n",err)
	}
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC (which Augur uses)
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	ss.mkt_order_id_ptr = mkt_order_ptr
	ss.db_logger = db_log
	ss.Info = info_log
	ss.Info.Printf("DB: connected to %v:%v",host,port)
	return ss
}
func show_connect_error() {
	fmt.Printf(`AugurExtractor: can't connect to PostgreSQL database.
				Check that you have set AUGUR_EXTRACTOR_USERNAME,AUGUR_EXTRACTOR_PASSWORD,AUGUR_EXTRACTOR_DATABASE
				and AUGUR_EXTRACTOR_HOST environment variables`);
}
func Connect_to_storage(mkt_order_ptr *int64,info_log *log.Logger) *SQLStorage {
	var err error
	host,port,err:=net.SplitHostPort(os.Getenv("AUGUR_EXTRACTOR_HOST"))
	if (err!=nil) {
		host=os.Getenv("AUGUR_EXTRACTOR_HOST")
		port="5432"
	}
	conn_str := "user='"+
				os.Getenv("AUGUR_EXTRACTOR_USERNAME") +
				"' dbname='" +
				os.Getenv("AUGUR_EXTRACTOR_DATABASE") +
				"' password='" +
				os.Getenv("AUGUR_EXTRACTOR_PASSWORD") +
				"' host='" +
				host +
				"' port='" +
				port +
				"'";
	db,err := sql.Open("postgres",conn_str);
	if (err!=nil) {
		show_connect_error()
	} else {

	}
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC (which Augur uses)
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	ss.mkt_order_id_ptr = mkt_order_ptr
	ss.Info = info_log
	ss.Info.Printf("DB: connected to %v:%v",host,port)
	return ss
}
func (ss *SQLStorage) Init_log(fname string) {

	f, err := os.OpenFile(fname,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Exiting Augur extractor with error: %v",err)
		os.Exit(1)
	}
	ss.db_logger = log.New(f,"DB: ",log.LstdFlags)
}
func (ss *SQLStorage) Log_msg(msg string) {
	if ss.db_logger !=nil {
		ss.db_logger.Printf(msg)
	} else {
		ss.Info.Printf(msg)
	}
}
