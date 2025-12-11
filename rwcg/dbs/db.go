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

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
)
var (
	zero *big.Int = big.NewInt(0)
	hundred *big.Int = big.NewInt(100)
)
type SQLStorage struct {
	db					*sql.DB
	db_logger			*log.Logger
	Info				*log.Logger
	schema_name			string
}
func (ss *SQLStorage) SchemaName() string { return ss.schema_name }
func (ss *SQLStorage) Db() *sql.DB {
	return ss.db
}
func New_sql_storage(info_log *log.Logger,db_log *log.Logger,host_port,db_name,user,pass string) *SQLStorage {
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
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC 
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	ss.db_logger = db_log
	ss.Info = info_log
	ss.Info.Printf("DB: connected to %v:%v",host,port)
	return ss
}
func show_connect_error() {
	fmt.Printf(`Extractor: can't connect to PostgreSQL database.
				Check that you have set EXTRACTOR_USERNAME,EXTRACTOR_PASSWORD,EXTRACTOR_DATABASE
				EXTRACTOR_HOST environment variables`);
}
func Connect_to_storage(info_log *log.Logger) *SQLStorage {
	var err error
	host,port,err:=net.SplitHostPort(os.Getenv("EXTRACTOR_HOST"))
	if (err!=nil) {
		host=os.Getenv("EXTRACTOR_HOST")
		port="5432"
	}
	conn_str := "user='"+
				os.Getenv("EXTRACTOR_USERNAME") +
				"' dbname='" +
				os.Getenv("EXTRACTOR_DATABASE") +
				"' password='" +
				os.Getenv("EXTRACTOR_PASSWORD") +
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
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	ss.Info = info_log
	return ss
}
func Connect_to_storage_with_schema(info_log *log.Logger,schema_name string) *SQLStorage {
	var err error
	host,port,err:=net.SplitHostPort(os.Getenv("EXTRACTOR_HOST"))
	if (err!=nil) {
		host=os.Getenv("EXTRACTOR_HOST")
		port="5432"
	}
	conn_str := "user='"+
				os.Getenv("EXTRACTOR_USERNAME") +
				"' dbname='" +
				os.Getenv("EXTRACTOR_DATABASE") +
				"' password='" +
				os.Getenv("EXTRACTOR_PASSWORD") +
				"' host='" +
				host +
				"' port='" +
				port +
				"'" +
				"search_path='"+
				schema_name+
				"'";
	db,err := sql.Open("postgres",conn_str);
	if (err!=nil) {
		show_connect_error()
	} else {

	}
	_,err = db.Exec("SET timezone TO 0")		// Setting timezone to UTC 
	if (err!=nil) {
		p.Fatalf(fmt.Sprintf("DB Error: %v",err));
	}

	ss := new(SQLStorage)
	ss.db = db
	ss.Info = info_log
	ss.Db_set_schema_name(schema_name)
	return ss
}
func (ss *SQLStorage) Init_log(fname string) {

	f, err := os.OpenFile(fname,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Exiting extractor with error: %v",err)
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
func (ss *SQLStorage) Db_set_schema_name(name string) {
	ss.schema_name = name
}
func (ss *SQLStorage) Set_search_path_to_schema_name() {

	ss.Info.Printf("Setting search path to %v\n",ss.schema_name)
	_,err := ss.db.Exec("SET SEARCH_PATH TO "+ss.schema_name)
	if (err!=nil) {
		ss.Info.Printf("DB Error: %v",err);
		os.Exit(1)
	}

}
func (ss *SQLStorage) Get_search_path() string {

	var query string
	query="SHOW search_path"
	var spath string
	err:=ss.db.QueryRow(query).Scan(&spath);
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
		os.Exit(1)
	}
	return spath
}
