package main

import (
	"log"
)

//App wide error handling
func handleError(err error) {
	if err != nil {
		log.Print(err)
	}
}

//
