package main

import "net/rpc"

type payload struct {
	ID   int
	Test string
}

type plugin struct {
	client      *rpc.Client
	description string
	keyword     string
}

func (p plugin) Help() (result string, err error) {
	err = p.client.Call("Plugin.Help", "", &result)
	return result, err
}

func (p plugin) Description() (result string, err error) {
	err = p.client.Call("Plugin.Description", "", &result)
	return result, err
}

func (p plugin) Run(args []string) (result payload, err error) {
	err = p.client.Call("Plugin.Run", args, &result)
	return result, err
}
