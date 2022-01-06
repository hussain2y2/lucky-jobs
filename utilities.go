package main

import (
	"regexp"
)

/**
 * Check if the given string is in the number format
 */
func checkNumber(str string) bool {
	numberP := regexp.MustCompile(`^[0-9]+$`)
	return numberP.MatchString(str)
}

/**
 * Check if the given string is in the string format
 */
func checkString(str string) bool {
	stringP := regexp.MustCompile(`^[a-zA-Z]+$`)
	return stringP.MatchString(str)
}
