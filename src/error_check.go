package main

import "log"

func CheckIfError(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}
