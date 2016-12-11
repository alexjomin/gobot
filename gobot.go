package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var intro = []string{
	"*I'm here to help*, you can talk to me directly via `DM` or in a channel via `@gobot`",
	"Here are the available commands :",
	"",
}

var botID string

func areYouTalkingToMeInAPrivateChannel(ev *slack.MessageEvent) (bool, bool) {

	// Check if event if from the bot or not
	log.Printf("Message received, user is %s and bot id is %s on channel %s\n ", ev.User, botID, ev.Channel)

	if ev.User == botID || ev.User == "" {
		return false, false
	}

	if strings.HasPrefix(ev.Channel, "D") {
		return true, true
	}

	userString := fmt.Sprintf("<@%s>", botID)

	f := strings.Fields(ev.Text)

	if f[0] == userString {
		// Removing user id from text
		ev.Text = strings.Join(f[1:], " ")
		return true, false
	}

	return false, false
}

func main() {

	d := newDispatcher()
	err := d.addPlugin("rules", "/Users/alexandrejomin/go/src/github.com/alexjomin/gobot/plugins/git-commit-rules/plugin")

	if err != nil {
		log.Fatalf("Error while adding plugin : %s", err)
	}

	if len(os.Args) != 2 {
		log.Fatalf("usage: mybot slack-bot-token\n")
	}

	api := slack.New(os.Args[1])
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	//api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {

			case *slack.HelloEvent:

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID

			case *slack.MessageEvent:

				// check if the message is adressed to the bot
				toMe, private := areYouTalkingToMeInAPrivateChannel(ev)

				// Not a message we need to deal with
				if !toMe {
					break
				}

				if !private {
				}

				r, _ := d.handle(ev)

				rtm.SendMessage(rtm.NewOutgoingMessage(r, ev.Channel))

				// Dispatcher
				//task, args, err := parseMessage(ev.Text)
				/*
					if err != nil {
						text := "Sorry I don't understand"
						rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
					} else {
						go task(rtm, ev, args)
					}
				*/

			case *slack.PresenceChangeEvent:
				log.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				log.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				log.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				log.Printf("Invalid credentials")
				break Loop

			default:

			}
		}
	}

}
