package main

import (
	"strings"

	"github.com/nlopes/slack"
)

var commit = []string{
	"```",
	"<type>(<scope>): <subject>",
	"or:",
	"<type>: <subject>",
	"",
	"Type is mandatory and the scope is optional.",
	"",
	"• feat: A new feature",
	"• fix: A bug fix",
	"• docs: Documentation only changes",
	"• style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
	"• refactor: A code change that neither fixes a bug nor adds a feature",
	"• perf: A code change that improves performance",
	"• test: Adding missing tests",
	"• chore: Changes to the build process or auxiliary tools and libraries such as documentation generation",
	"```",
	"",
	"https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md",
}

func taskRule(rtm *slack.RTM, ev *slack.MessageEvent, params []string) {

	if len(params) == 0 {
		return
	}

	if params[0] == "-help" {
		text := "`rules commit`, gives a reminder how to name your commit"
		rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
		return
	}

	switch params[0] {
	case "commit":
		rtm.SendMessage(rtm.NewOutgoingMessage(strings.Join(commit, "\n"), ev.Channel))
	default:
		rtm.SendMessage(rtm.NewOutgoingMessage("sorry, i don't understand.", ev.Channel))
	}
}
