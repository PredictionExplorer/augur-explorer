package main
import (
	//"time"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	EXTRA_OFFSET_TS int64 = 60*10	// 1 hour of extra overlap to the next interval , for safety
)
func manage_stat_periods(s *SQLStorage,default_duration int64) {

	last_period_ts,duration,_ :=s.Bigstats_get_last_period()
	if last_period_ts == 0 {
		duration = default_duration
		last_period_ts = s.Bigstats_get_first_block_timestamp()
		last_period_ts = last_period_ts / duration
		last_period_ts = last_period_ts * duration
	}
	Info.Printf("Statistics: last_period_ts = %v, duration=%v\n",last_period_ts,duration)
	last_block_ts := storage.Bigstats_get_last_block_timestamp()
	if last_block_ts != 0 {
		Info.Printf("Stats: last_block_ts=%v, last_period_ts=%v\n",last_block_ts,last_period_ts)
		Info.Printf("If condition: %v > %v \n",last_period_ts + duration + EXTRA_OFFSET_TS,last_block_ts)
		if (last_period_ts + duration + EXTRA_OFFSET_TS) < last_block_ts {
			// there is enough data for a new interval
			new_interval_ts := last_period_ts + duration
			Info.Printf(
				"Statistics: closing period. last_period_ts=%v, duration %v, new_interval_ts=%v\n",
				last_period_ts,duration,new_interval_ts,
			)
			storage.Bigstats_close_period(new_interval_ts,duration)
		}
	} else {
		// no blocks in DB
	}


}
