package main

import (
	"fmt"
	"testing"
)

func TestNewRepository(t *testing.T) {
	r, err := NewRepository("git@bitbucket.org:eliocity", "api-gateway")

	if err != nil {
		t.Error(err)
	}

	e := "git@bitbucket.org:eliocity/api-gateway"
	if r.remoteURL != e {
		t.Errorf("remoteURL is not correct, %s expected but get %s", r.remoteURL, e)
	}
}

func TestGetBitbucketLink(t *testing.T) {

	r, err := NewRepository("git@bitbucket.org:eliocity", "api-gateway")

	if err != nil {
		t.Error(err)
	}

	r.currentHash = "a1b2c3d6"

	p := r.localPath + "foo/bar"

	fmt.Println(r.getBitbucketLink(p, 0))

}
