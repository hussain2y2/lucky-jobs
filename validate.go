package main

import (
	"log"
	"strings"
)

/**
 * Validate the given string format
 */
func testValidity(text string) bool {
	splitted := strings.Split(text, "-")
	if len(splitted) >= 2 && checkNumber(splitted[0]) && checkString(splitted[1]) {
		return true
	}
	log.Fatal("The format is not correct!")
	return false
}
