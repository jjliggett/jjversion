package main

import (
	"flag"
	"os"
)

func main() {

	shouldDisplayVersionFlag := flag.Bool("version", false, "Displays version of jjversion")

	flag.Parse()

	shouldDisplayVersion := *shouldDisplayVersionFlag

	if shouldDisplayVersion {
		printVersion()
		os.Exit(0)
	}

	calculateVersion()
}
