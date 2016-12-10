package main

import "github.com/nlopes/slack"

type tasK interface {
	register(string) error
	run() task
}

func taskJoke(rtm *slack.RTM, ev *slack.MessageEvent, params []string) {

	if len(params) == 0 {
		return
	}

	switch params[0] {
	case "navette":
		rtm.SendMessage(rtm.NewOutgoingMessage("https://www.youtube.com/watch?v=trvqZZrKQqQ", ev.Channel))
	default:
		rtm.SendMessage(rtm.NewOutgoingMessage("sorry, no joke.", ev.Channel))
	}
}
