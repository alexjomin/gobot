package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
	"strings"

	"github.com/natefinch/pie"
	"github.com/nlopes/slack"
)

type dispatcher struct {
	plugins map[string]plugin
}

func newDispatcher() *dispatcher {
	d := &dispatcher{}
	d.plugins = map[string]plugin{}
	return d
}

func (d *dispatcher) handle(e *slack.MessageEvent) (response string, err error) {

	elements := strings.Fields(e.Text)
	keyword := elements[0]

	// global help, will display the intro and all plugin's decriptions
	if keyword == "help" {
		a := intro
		for _, p := range d.plugins {
			a = append(a, "*"+p.keyword+"* : "+p.description)
		}
		response = strings.Join(a, "\n")
		return
	}

	action := elements[1]

	// help asked for the plugin
	if action == "help" {
		response, err = d.plugins[keyword].Help()
		return response, err
	}

	// Action
	r, err := d.plugins[keyword].Run(elements)
	response = fmt.Sprintf("%v", r)
	return response, err
}

func (d *dispatcher) addPlugin(keyword, path string) error {

	log.Printf("Trying to add plugin, path: '%s' with keyword: '%s'", keyword, path)

	client, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stderr, path)

	if err != nil {
		log.Printf("Error trying to load plugin: %s", err)
		return err
	}
	//defer client.Close()

	p := plugin{client, keyword, ""}

	res, err := p.Description()

	if err != nil {
		log.Printf("Error trying to get the description of the plugin: %s", err)
		return err
	}

	p.description = res

	d.plugins[keyword] = p

	return nil
}
