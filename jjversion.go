package main

import (
	"flag"
	"fmt"
	"os"

	jjvercore "github.com/jjliggett/jjversion/jjvercore"
)

func main() {

	flag.Usage = func() {
		printVersion()
		fmt.Println("Evaluate a SemVer version for a git repository according to configuration in versioning.yaml.")
		fmt.Println("Usage: jjversion")
		fmt.Println()

		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("For more info, see: https://github.com/jjliggett/jjversion")
	}

	shouldDisplayVersionFlag := flag.Bool("version", false, "Displays version of jjversion utility")

	flag.Parse()

	shouldDisplayVersion := *shouldDisplayVersionFlag

	if shouldDisplayVersion {
		printVersion()
		os.Exit(0)
	}

	v := jjvercore.CalculateVersion()
	fmt.Println(v.Json())
}
