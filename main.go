package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	text := "23-ab-48-caba-56-haha"
	fmt.Println(text)
	testValidity(text)
}

func testValidity(text string) bool {
	splitted := strings.Split(text, "-")
	if len(splitted) >= 2 {
		numberP := regexp.MustCompile(`^[0-9]+$`)
		stringP := regexp.MustCompile(`^[a-zA-Z]+$`)

		if numberP.MatchString(splitted[0]) && stringP.MatchString(splitted[1]) {
			return true
		}
		log.Fatal("The format is not correct!")
	} else {
		log.Fatal("The format is not correct!")
	}
	return false
}
