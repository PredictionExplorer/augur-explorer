package main
import (
	//"time"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	EXTRA_OFFSET_TS int64 = 60*5	// 1 hour of extra overlap to the next interval , for safety
)
func manage_stat_periods(s *SQLStorage,ethusd_schema string,default_duration int64) {

  again:
	last_period_ts,duration,_ :=s.Bigstats_get_last_period()
	if last_period_ts == 0 {
		duration = default_duration
		last_period_ts = s.Bigstats_get_first_block_timestamp()
		last_period_ts = last_period_ts / duration
		last_period_ts = last_period_ts * duration
	} else {
		last_period_ts = last_period_ts + duration
	}
	Info.Printf("Statistics: last_period_ts = %v, duration=%v\n",last_period_ts,duration)
	last_block_ts := storage.Bigstats_get_last_block_timestamp()
	if last_block_ts != 0 {
		Info.Printf("Stats: last_block_ts=%v, last_period_ts=%v\n",last_block_ts,last_period_ts)
		if (last_period_ts + duration + EXTRA_OFFSET_TS) < last_block_ts {
			// there is enough data for a new interval
			new_interval_ts := last_period_ts + duration
			Info.Printf(
				"Statistics: closing period. last_period_ts=%v, duration %v\n",
				last_period_ts,duration,
			)
			storage.Bigstats_close_period(ethusd_schema,last_period_ts,duration)
			Info.Printf("Statistics: new interval set to %v\n",new_interval_ts)
			goto again
		}
	} else {
		// no blocks in DB
	}
	Info.Printf("exiting manage_stat_periods()\n")

}
