package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	text := "23-ab-48-caba-56-haha"
	if testValidity(text) {
		var numbersArr = []int64{}
		var stringsArr = []string{}
		tempArray := strings.Split(text, "-")

		for _, item := range tempArray {
			if checkNumber(item) {
				num, _ := strconv.ParseInt(item, 10, 64)
				numbersArr = append(numbersArr, num)
			}
			if checkString(item) {
				stringsArr = append(stringsArr, item)
			}
		}

		fmt.Println(numbersArr, stringsArr)
	}
}

func testValidity(text string) bool {
	splitted := strings.Split(text, "-")
	if len(splitted) >= 2 && checkNumber(splitted[0]) && checkString(splitted[1]) {
		return true
	}
	log.Fatal("The format is not correct!")
	return false
}

func checkNumber(str string) bool {
	numberP := regexp.MustCompile(`^[0-9]+$`)
	return numberP.MatchString(str)
}

func checkString(str string) bool {
	stringP := regexp.MustCompile(`^[a-zA-Z]+$`)
	return stringP.MatchString(str)
}

func averageNumber() {

}

func wholeStory() {

}
