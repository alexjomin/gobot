package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var id string

func areYouTalkingToMeInAPrivateChannel(ev *slack.MessageEvent) (bool, bool) {

	// Check if event if from the bot or not
	fmt.Printf("message user is %s and bot id is %s on channel %s\n ", ev.User, id, ev.Channel)
	if ev.User == id || ev.User == "" {
		return false, false
	}

	if strings.HasPrefix(ev.Channel, "D") {
		return true, true
	}

	userString := fmt.Sprintf("<@%s>", id)

	f := strings.Fields(ev.Text)

	if f[0] == userString {

		// Removing user id from text
		ev.Text = strings.Join(f[1:], " ")
		fmt.Printf("test : %s\n", ev.Text)
		return true, false
	}

	return false, false
}

func main() {

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
			fmt.Printf("type :%v\n", msg.Type)
			switch ev := msg.Data.(type) {

			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				id = ev.Info.User.ID

				fmt.Println("id : ", id)

			case *slack.MessageEvent:

				// check if the message is adressed to the bot
				toMe, private := areYouTalkingToMeInAPrivateChannel(ev)

				// Not a message we need to deal with
				if !toMe {
					break
				}

				if !private {

				}

				task, args, err := parseMessage(ev.Text)

				if err != nil {
					text := "Sorry I don't understand"
					rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
				} else {
					go task(rtm, ev, args)
				}

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}

}
