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
		print_version()
		os.Exit(0)
	}

	calculate_version()
}
