// Changes name of Discord channel
package main

import (
	"fmt"
	//"context"
	//"errors"
	"os"
	//"time"

	"github.com/andersfylling/disgord"
	//"github.com/sirupsen/logrus"
)
const (//917879763404718100
	MintChannelID_Uint			uint64 = 918642461314785290
	MintChannelID				= disgord.Snowflake(MintChannelID_Uint)
	PriceChannelID_Uint			uint64 = 918643820734869525
	PriceChannelID				= disgord.Snowflake(PriceChannelID_Uint)
	LastDateChannelID_Uint		uint = 918653298813313044
	LastDateChannelID			= disgord.Snowflake(LastDateChannelID_Uint)
	NumMintsUnicodeChar		string = "🪙 "
	LastPriceUnicodeChar	string = "💲"
	EthSign					string = "Ξ"
)
func jsonbytes(format string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, args...))
}
func main() {

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISGORD_TOKEN"),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	//client.Channel(PriceChannelID).UpdateBuilder().SetName(NumMintsUnicodeChar+" total mints : 111").Execute()
	//client.Channel(PriceChannelID).UpdateBuilder().SetName("last price "+LastPriceUnicodeChar+" : 0.0233 " + EthSign).Execute()
	client.Channel(LastDateChannelID).UpdateBuilder().SetName(" Last mint : "+" 4 Jan , 19:03").Execute()
}
