package main

import (
	"os"
	"fmt"
	"time"
	"github.com/slack-go/slack"
)
const (
	SEND_ALARMS_INTERVAL		int64	= 60 * 60 *2	// the alarm will repeat every time this interval expires
)
const (
	ALARM_CODE_COSMIC_APP_LAYER_BLOCK_LAG	int64			= 1
)
var (
	slack_channel				string = "#cosmic-signature"
	bot_token					string = os.Getenv("OAUTH_BOT_TOKEN")
	last_alarm_timestamp		int64 = 0
	alarm_status				map[int64]int64 = make(map[int64]int64)

	// alarm variables (alarm codes)
	AppCosmicLaggingTimestamp	int64 = 0
)

func send_alarm_slack(alarm_code int64,alarm_text string) {

	switch alarm_code {
		case ALARM_CODE_COSMIC_APP_LAYER_BLOCK_LAG:
		default:
			Info.Printf("Alarm code %v undefined (has to be added to source code)\n",alarm_code)
	}
	cur_ts := time.Now().Unix()
	elapsed_ts := cur_ts - alarm_status[alarm_code]
	if elapsed_ts > SEND_ALARMS_INTERVAL {
		alarm_status[alarm_code] = cur_ts
		api := slack.New(bot_token)
		channel_id_posted, timestamp, err := api.PostMessage(
			slack_channel,
			slack.MsgOptionText(alarm_text,false),
		)
		_ = channel_id_posted; _=timestamp;
		if err != nil {
			err_str := fmt.Sprintf("Error at sending slack alarm: %v\n",err.Error())
			update_global_errors(err_str)
		}
	}
}
