package main

import (
	"flag"
	"fmt"
	"os"

	jjvercore "github.com/jjliggett/jjversion/jjvercore"
)

func main() {

	shouldDisplayVersionFlag := flag.Bool("version", false, "Displays version of jjversion")

	flag.Parse()

	shouldDisplayVersion := *shouldDisplayVersionFlag

	if shouldDisplayVersion {
		printVersion()
		os.Exit(0)
	}

	v := jjvercore.CalculateVersion()
	fmt.Println(v.Json())
}
