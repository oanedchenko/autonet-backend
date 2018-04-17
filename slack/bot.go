package slack

import (
	"os"
	"github.com/nlopes/slack"
	"log"
)

func init() {
	log.Printf("Starting autonet bot...")
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	api.SetDebug(true)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	go runBot()

	log.Printf("Autonet bot is listerning...")
}

func runBot() {
	// todo: https://rsmitty.github.io/Slack-Bot/
}
