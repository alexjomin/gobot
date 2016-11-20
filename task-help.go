package main

import (
	"strings"

	"github.com/nlopes/slack"
)

func taskHelp(rtm *slack.RTM, ev *slack.MessageEvent, params []string) {

	intro := []string{
		"*I'm here to help*, you can talk to me directly via `DM` or in a channel via `@gobot`",
		"Here are the available commands :",
	}

	lines := []string{
		"• `lint` Clone a git repository and apply a bunch of lint tools - see `lint -help`",
		"• `rules` Give you a few tips about codings rule - see `rules -help`",
		"• `joke` undocumented :-]",
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(strings.Join(intro, "\n"), ev.Channel))
	rtm.SendMessage(rtm.NewOutgoingMessage(strings.Join(lines, "\n"), ev.Channel))

}
