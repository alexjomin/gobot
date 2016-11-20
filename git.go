package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"path"
	"strings"
	"time"
)

type repository struct {
	remoteBaseURL string
	remoteURL     string
	name          string
	localPath     string
	currentBranch string
	currentHash   string
	OnlineCodeURL string
}

func NewRepository(baseURL, name string) (*repository, error) {

	path := "/tmp/" + randStringRunes(6)

	// Remvoe trailing slash if needed
	b := strings.TrimSuffix(baseURL, "/")

	return &repository{
		localPath:     path,
		remoteBaseURL: baseURL,
		remoteURL:     b + "/" + name,
		name:          name,
		currentBranch: "master",
	}, nil

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (r *repository) getBitbucketLink(p string, line int) string {
	segments := strings.Split(p, r.localPath)

	file := path.Base(segments[1])

	return fmt.Sprintf(
		"https://bitbucket.org/eliocity/%s/src/%s/%s?at=%s&fileviewer=file-view-default#%s-%d", r.name, r.currentHash, segments[1], r.currentBranch, file, line)
}

func (r *repository) clone(branch string) error {

	cmd := exec.Command("git", "clone", r.remoteURL, r.localPath, "-b", r.currentBranch, "--single-branch")

	var out bytes.Buffer
	var errb bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		fmt.Print(errb.String())
		return err
	}

	fmt.Println(errb.String())

	err = r.getCurrentHash()

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) clean() error {
	cmd := exec.Command("rm", "-rf", r.localPath)
	err := cmd.Run()
	return err
}

func (r *repository) getCurrentHash() error {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = r.localPath

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return err
	}

	r.currentHash = strings.TrimSpace(out.String())

	return nil
}
