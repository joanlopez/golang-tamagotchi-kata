all: help

##
##
## Golang Tamagotchi Kata
## ======================
##
## @author Joan López de la Franca <me@joanlopez.cat>
## @description Tamagotchi Kata written in Golang
##
## Available commands are:
##

##  help:		Help
.PHONE : help
help : Makefile
	@sed -n 's/^##//p' $<

##  tests:	Run tests
.PHONY : tests
tests:
	go test ./...

##
##