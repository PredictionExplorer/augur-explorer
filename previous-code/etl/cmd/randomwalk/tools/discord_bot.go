package main

import (
	//"context"
	//"errors"
	"os"
	//"time"

	"github.com/andersfylling/disgord"
	//"github.com/sirupsen/logrus"
)
const (
	ChannelID				= disgord.Snowflake(915980333537718306)
)
func main() {

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISGORD_TOKEN"),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	f1, err := os.Open("/tmp/token.png")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	_, err = client.Channel(ChannelID).CreateMessage(
			&disgord.CreateMessageParams{
				Content: "This is a bot test http:://google.com",
				Files: []disgord.CreateMessageFileParams{
					{f1, "token.png", false},
				},
				Embed: &disgord.Embed{
					Description: "http://google.com",
					URL: "http://google.com",
				},
			},
	)
	if err != nil {
		client.Logger().Error("unable to upload images.", err)
	}
}
