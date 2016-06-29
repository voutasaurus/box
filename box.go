package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	switch flag.Arg(0) {
	case "generate":
		if err := cmdGenerate(); err != nil {
			log.Fatal(err)
		}
	case "open":
		if err := cmdOpen(mustGetKey(), flag.Arg(1)); err != nil {
			log.Fatal(err)
		}
	case "seal":
		if err := cmdSeal(mustGetKey(), flag.Arg(1)); err != nil {
			log.Fatal(err)
		}
	default:
		cmdHelp(flag.Arg(1))
	}
}
