package main

import (
	"log"

	notification "github.com/senvardarsemih/slack-notifier/notifier"
)

func main() {

	sc := notification.SlackClient{
		WebHookUrl: "https://YOUR_SLACK_WEB_HOOK_URL", //your slack webhook url
		UserName:   "Go Notifier Bot",//the name of the bot in a slack message
		Channel:    "#YOUR_SLACK_CHANNEL_NAME",// your slack channel name
	}
	sr := notification.SlackRequest{
		Text:      "Houston, we have a problem!",
		IconEmoji: ":hot-desking:", //the name of the icon image, otherwise it is a default image of incoming webhook.
	}
	err := sc.SendSlackNotification(sr)
	if err != nil {
		log.Fatal(err)
	}

	//Job Notification Sample
	// sr := SlackJobNotification{
	//     Text:  "This is attachment message",
	//     Details: "details of the jobs",
	//     Color: "warning",
	//     IconEmoji: ":hammer_and_wrench",
	// }
	// err := sc.SendJobNotification(sr)
	// if err != nil {
	//     log.Fatal(err)
	// }
}
