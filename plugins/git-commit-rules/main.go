package main

import (
	"log"
	"net/rpc/jsonrpc"
	"strings"

	"github.com/natefinch/pie"
)

type Payload struct {
	ID   int
	Test string
}

var description = "Give you a few tips about coding rules - see `help` for more details"

var help = []string{
	"`rules commit` gives you a reminder how to name your commit",
	"`rules ranges` gives you a reminder about x-ranges semver",
}

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

func main() {

	p := pie.NewProvider()

	if err := p.RegisterName("Plugin", api{}); err != nil {
		log.Fatalf("failed to register Plugin: %s", err)
	}

	p.ServeCodec(jsonrpc.NewServerCodec)
}

type api struct{}

func (api) Description(args string, response *string) error {
	*response = description
	return nil
}

func (api) Help(args string, response *string) error {
	*response = strings.Join(help, "\n")
	return nil
}

func (api) Run(args []string, response *Payload) error {
	//*response = strings.Join(commit, "\n")
	*response = Payload{1, "test"}
	return nil
}
