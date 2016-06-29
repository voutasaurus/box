package main

import "log"

const defaultHelp = `box is a tool for encrypting and decrypting files

Usage:

	box command [arguments]
	
The commands are:

	generate	generate a secretbox key
	seal		encrypt a file to stdout
	open		decrypt a file to stdout

Use "box help [command]" for more information about a command.	

`

var help = map[string]string{}

func registerHelp(topic, helpStr string) {
	help[topic] = helpStr
}

func cmdHelp(topic string) {
	h := help[topic]
	if h == "" {
		h = defaultHelp
	}
	log.Fatal(h)
}
