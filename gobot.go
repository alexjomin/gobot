package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"

	"github.com/natefinch/pie"
	"github.com/nlopes/slack"
)

var intro = []string{
	"*I'm here to help*, you can talk to me directly via `DM` or in a channel via `@gobot`",
	"Here are the available commands :",
	"",
}

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

type plug struct {
	client      *rpc.Client
	description string
	keyword     string
}

func (p plug) Help() (result string, err error) {
	err = p.client.Call("Plugin.Help", "", &result)
	return result, err
}

func (p plug) Description() (result string, err error) {
	err = p.client.Call("Plugin.Description", "", &result)
	return result, err
}

func (p plug) Run(args []string) (result string, err error) {
	fmt.Println("run ! :")
	err = p.client.Call("Plugin.Run", args, &result)
	return result, err
}

type plugin struct {
	keyword string
}

type dispatcher struct {
	plugins map[string]plug
}

func NewDispatcher() *dispatcher {
	d := &dispatcher{}
	d.plugins = map[string]plug{}
	return d
}

func (d *dispatcher) handle(e *slack.MessageEvent) (response string, err error) {

	elements := strings.Fields(e.Text)
	keyword := elements[0]

	// global help, will display the intro and all plugin's decriptions
	if keyword == "help" {
		a := intro
		for _, p := range d.plugins {
			a = append(a, "*"+p.keyword+"* : "+p.description)
		}
		response = strings.Join(a, "\n")
		return
	}

	action := elements[1]

	// help asked for the plugin
	if action == "help" {
		response, err = d.plugins[keyword].Help()
		return response, err
	}

	// Action
	response, err = d.plugins[keyword].Run(elements)
	return response, err
}

func (d *dispatcher) addPlugin(keyword, path string) error {

	client, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stderr, path)

	if err != nil {
		fmt.Printf("Error running plugin: %s", err)
		return err
	}
	//defer client.Close()

	p := plug{client, "", ""}

	res, err := p.Description()

	if err != nil {
		return err
	}

	p.keyword = keyword
	p.description = res

	d.plugins[keyword] = p

	return nil
}

func main() {

	d := NewDispatcher()
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

	// TODO Create a dispatcher :
	// 1. First create the dispatcher
	// 2. Register tasks with a key word (with config ?)
	// 3. Pass the message to the dispatcher

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Printf("type :%v\n", msg)
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
