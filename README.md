# gobot

## Purpose

> As a lead developer I don't wan't to waste my time to check PR of the team so I could spend more time drinking belgian beers.

Ok, calm down. That was a joke, PR are really great.

`gobot` is a slack bot to help you and your team to produce better golang code.

It relies on [gometalinter](https://github.com/alecthomas/gometalinter) to perform :
+ errcheck
+ vet
+ gosimple
+ golint
+ gofmt
+ goimports

## Install
This project uses [Glide](https://github.com/Masterminds/glide) a Vendor Package Management
	
	glide install
	make build

## Start
	gobot your-slack-token