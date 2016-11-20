package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var emptyChar = regexp.MustCompile("^\\s+")

func extract(path string, line int) string {

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 1
	for scanner.Scan() {
		if i == line {
			return removeWhiteSpace(scanner.Bytes())
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return ""
	}

	return ""
}

func removeWhiteSpace(s []byte) string {
	r := emptyChar.ReplaceAll(s, []byte(""))
	return string(r)
}
