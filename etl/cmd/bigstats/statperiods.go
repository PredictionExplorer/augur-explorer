package main
import (
	//"time"
)
const (
	EXTRA_OFFSET_TS int64 = 60*60	// 1 hour of extra overlap to the next interval , for safety
)
func manage_stat_periods(default_duration int64) {


	last_period,duration,err:=storage.Bigstats_get_last_period()
	if err != nil {
		duration = default_duration
	}

//	for {
		last_block_ts := storage.Bigstats_get_last_block_timestamp()
		if last_block_ts != 0 {
			last_period_ts,_,err := storage.Bigstats_get_last_period()
			if err != nil {
				if (last_period_ts + duration + EXTRA_OFFSET_TS) > last_block_ts {
					// there is enough data for a new interval
					new_interval_ts := last_period + duration
					storage.Bigstats_close_period(new_interval_ts,duration)
				}
			} else {
				Error.Printf("Error getting last period in manage_state_periods(): %v\n",err)
			}
		} else {
			// no blocks in DB
		}

//		time.Sleep(1 * time.Hour)
//	}

}
