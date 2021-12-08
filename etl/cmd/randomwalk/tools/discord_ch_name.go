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
	ChannelID_Uint			uint64 = 917879763404718100
	ChannelID				= disgord.Snowflake(917879763404718100)
)
func jsonbytes(format string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, args...))
}
func main() {

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISGORD_TOKEN"),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	cache := client.Cache
	//updater := cache.CacheUpdater
	data := jsonbytes(`{"id":%d,"name":"%s"}`, ChannelID_Uint, "NewChan")
	//channel_update,err := updater.ChannelUpdate(data)
	channel_update,err := cache().ChannelUpdate(data)
	if err != nil {
		fmt.Printf("Channel Update error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n",channel_update)
}
