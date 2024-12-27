package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_NFT_donations(offset,limit int) []p.CGNFTDonation{

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"d.round_num,"+
				"nft.address_id,"+
				"d.idx,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"ORDER BY d.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGNFTDonation
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.NFTTokenId,
			&rec.RoundNum,
			&rec.TokenAddressId,
			&rec.Index,
			&rec.TokenAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_NFT_donation_info(id int64) (bool,p.CGNFTDonation) {

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"nft.address_id,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE d.id=$1"

	row := sw.S.Db().QueryRow(query,id)
	var err error
	var rec p.CGNFTDonation
	rec.RecordId = id
	err=row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.NFTTokenId,
		&rec.TokenAddressId,
		&rec.TokenAddr,
		&rec.NFTTokenURI,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_NFT_donation_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return true,rec
}
func (sw *SQLStorageWrapper) Get_donated_nft_claims(offset,limit int) []p.CGDonatedNFTClaimRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"c.id,"+
				"c.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT,"+
				"c.time_stamp,"+
				"c.round_num,"+
				"ta.addr,"+
				"c.token_id, "+
				"c.idx, "+
				"c.winner_aid,"+
				"wa.addr, "+
				"da.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_donated_nft_claimed c "+
				"LEFT JOIN transaction t ON t.id=c.tx_id "+
				"LEFT JOIN address ta ON c.token_aid=ta.address_id "+
				"LEFT JOIN address wa ON c.winner_aid=wa.address_id "+
				"LEFT JOIN cg_nft_donation d ON c.idx=d.idx "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id " +
			"ORDER BY c.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGDonatedNFTClaimRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGDonatedNFTClaimRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.TokenAddr,
			&rec.NFTTokenId,
			&rec.WinnerIndex,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.DonorAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_nft_donations_by_prize(round_num int64) []p.CGNFTDonation {

	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.round_num,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"d.idx,"+
				"nft.address_id,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE d.round_num= $1 " +
			"ORDER BY d.id DESC"

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGNFTDonation
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.NFTTokenId,
			&rec.Index,
			&rec.TokenAddressId,
			&rec.TokenAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unclaimed_donated_nfts_by_prize(round_num int64) []p.CGNFTDonation {

	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.round_num,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"d.idx,"+
				"nft.address_id,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_donated_nft_claimed dc ON d.idx = dc.idx "+
			"WHERE d.round_num= $1 AND dc.idx IS NULL " +
			"ORDER BY d.id DESC"

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGNFTDonation
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.NFTTokenId,
			&rec.Index,
			&rec.TokenAddressId,
			&rec.TokenAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donated_token_distribution() []p.CGDonatedTokenDistrRec {

	var query string
	query = "SELECT "+
				"ca.addr,"+
				"ns.num_donated "+
			"FROM "+sw.S.SchemaName()+".cg_nft_stats ns "+
				"LEFT JOIN address ca ON ns.contract_aid=ca.address_id "+
			"ORDER BY ns.num_donated DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGDonatedTokenDistrRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGDonatedTokenDistrRec
		err=rows.Scan(
			&rec.ContractAddr,
			&rec.NumDonatedTokens,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
