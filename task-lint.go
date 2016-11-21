package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/nlopes/slack"
)

type task func(*slack.RTM, *slack.MessageEvent, []string)

func taskLint(rtm *slack.RTM, ev *slack.MessageEvent, params []string) {

	if params[0] == "-help" {
		text := "`lint repository#branch`, branch is facultative, `master` is the default branch"
		rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
		return
	}

	branchSplit := strings.Split(params[0], "#")

	branch := "master"
	repo := params[0]

	if len(branchSplit) == 2 {
		branch = branchSplit[1]
		if branch == "" {
			branch = "master"
		}
		repo = branchSplit[0]
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf(":rocket: Roger that, clone the `%s` repository and set HEAD to `%s`", repo, branch), ev.Channel))

	repository, err := NewRepository("git@bitbucket.org:eliocity", repo)
	repository.currentBranch = branch

	if err != nil {
		text := ":warning: Holy crap, an error occured"
		rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
		return
	}

	repository.clone(branch)
	defer repository.clean()

	if err != nil {
		text := ":warning: Holy crap, an error occured while trying to clone the repo, please check that the syntax is correct"
		rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
		return
	}

	prm := slack.NewPostMessageParameters()
	prm.Markdown = true

	results := lint(repository.localPath)

	for i, r := range results {

		at := slack.Attachment{}
		at.Color = "warning"
		at.AuthorName = r.linter
		at.Title = "See the code on bitbucket.org"
		at.TitleLink = repository.getBitbucketLink(r.path, r.line)
		at.Text = r.message
		at.Footer = fmt.Sprintf("%s - line %d", r.path, r.line)

		at.Fields = []slack.AttachmentField{
			{
				Title: "Level",
				Value: r.level,
				Short: true,
			},
		}

		prm.Attachments = append(prm.Attachments, at)

		if i == 9 {
			break
		}
	}

	msg := "Well done ! All clear."

	if len(results) > 0 {
		msg = fmt.Sprintf(":checkered_flag: Done, please find the results below. Found %d points", len(results))
	}

	if len(results) > 10 {
		msg = fmt.Sprintf("%s, only the 10 first one will be shown", msg)
	}

	rtm.PostMessage(ev.Channel, msg, prm)

}

func lint(s string) []*lintResult {

	buffer := []*lintResult{}

	cmd := exec.Command("gometalinter",
		s+"/...",
		"--deadline=30s",
		"--vendor",
		"--disable-all",
		"--enable=errcheck",
		"--enable=vet",
		"--enable=vetshadow",
		"--enable=gosimple",
		"--enable=golint",
		"--enable=gofmt",
		"--enable=goimports",
	)

	var out bytes.Buffer
	var errb bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &errb

	// @TODO CHECK ERROR
	fmt.Printf("error :\n %s \n", errb.String())

	err := cmd.Run()

	if err != nil {
		fmt.Print(errb.String())
		fmt.Print(out.String())
	}

	lines := parseResult(out.String())

	for i, line := range lines {
		if i > 20 {
			break
		}
		result, _ := parseLine(line)
		if result != nil {
			//result.extractCode()
			buffer = append(buffer, result)
		}

	}

	return buffer
}
