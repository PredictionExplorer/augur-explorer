// Changes name of Discord channel
package main

import (
	"fmt"
	//"context"
	//"errors"
	"os"
	"strconv"
	//"time"

	"github.com/andersfylling/disgord"
	//"github.com/sirupsen/logrus"
)
const (//917879763404718100
)
func jsonbytes(format string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, args...))
}
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v [channel_id]\n",os.Args[0])
		os.Exit(1)
	}
	chid,err := strconv.ParseUint(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error converting channel id: %v\n",err)
		os.Exit(1)
	}
	ChannelID := disgord.Snowflake(chid)
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISGORD_TOKEN"),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	ch_obj, err := client.Channel(ChannelID).Get()
	if err != nil {
		fmt.Printf("Error getting channel object: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("User limit: %v\n",ch_obj.UserLimit)

}
