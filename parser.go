package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile("([^:]+):(\\d*):(\\d*?):(\\w+): (.+) .*\\((\\w+)\\)")

type lintResult struct {
	path    string
	line    int
	column  int
	level   string
	message string
	linter  string
	code    string
}

func (l *lintResult) format() string {
	s := fmt.Sprintf("`%s` - %s:%d:%d - *%s* - %s\n", l.level, l.path, l.line, l.column, l.message, l.linter)
	return s
}

func newLintResult(a []string) (*lintResult, error) {

	if len(a) < 2 {
		return nil, errors.New("Not a valid slice")
	}

	line, err := strconv.Atoi(a[1])

	if err != nil {
		line = 0
	}

	column, err := strconv.Atoi(a[2])

	if err != nil {
		column = 0
	}

	l := &lintResult{
		path:    a[0],
		line:    line,
		column:  column,
		level:   a[3],
		message: a[4],
		linter:  a[5],
	}

	return l, nil

}

func parseResult(out string) []string {
	return strings.Split(out, "\n")
}

func parseLine(line string) (*lintResult, error) {
	l := reg.FindAllStringSubmatch(line, -1)
	if len(l) > 0 {
		a := l[0]
		return newLintResult(a[1:])
	}
	return nil, nil
}
