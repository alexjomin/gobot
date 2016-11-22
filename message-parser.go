package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseMessage(content string) (task, []string, error) {

	args := []string{}

	parts := strings.Fields(content)

	fmt.Printf("Parsing message `%s` \n", content)

	if parts[0] != "help" && len(parts) < 2 {
		return nil, args, errors.New("Can't parse message :" + content)
	}

	action := parts[0]

	switch action {
	case "lint":
		args = append(args, parts[1])
		return taskLint, args, nil
	case "joke":
		args = append(args, parts[1])
		return taskJoke, args, nil
	case "help":
		return taskHelp, args, nil
	case "rules":
		args = append(args, parts[1])
		return taskRule, args, nil
	}

	return nil, args, errors.New("Can't parse message :" + content)

}
