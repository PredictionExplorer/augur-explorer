#!/bin/bash

DB=$1
if test -z "$DB"
then
	echo "usage:" $0 "[db_name]"
	exit 1
fi
HOST="-Uatest -h augur"
psql $HOST $DB -P pager=off --quiet -c "\copy (SELECT FLOOR(EXTRACT(EPOCH FROM time_stamp)),evt_code,order_hash,chain_id,exchange_addr,maker_addr,maker_asset_data,maker_fee_asset_data,ROUND(maker_asset_amount*1e+18)_amount,ROUND(maker_fee*1e+18),taker_address,taker_asset_data,taker_fee_asset_data,ROUND(taker_asset_amount*1e+18),ROUND(taker_fee*1e+18),sender_address,fee_recipient_address,EXTRACT(EPOCH FROM expiration_time)::BIGINT,salt,signature,ROUND(fillable_amount*1e+18),ma.addr,m.num_ticks FROM mesh_evt JOIN market AS m ON mesh_evt.market_aid=m.market_aid JOIN address AS ma ON m.market_aid=ma.address_id) TO STDOUT DELIMITER E'\t' "
