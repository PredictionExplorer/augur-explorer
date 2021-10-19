package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_active_offers() []p.RW_API_Offer {

	records := make([]p.RW_API_Offer,0,16)

	var query string
	query = "SELECT " +
				

	return records
}
