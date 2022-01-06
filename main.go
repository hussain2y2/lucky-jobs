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
		fmt.Println("Average: ", averageNumber(numbersArr))
		fmt.Println("Whole Story: ", wholeStory(stringsArr))
		shortest, longest, average := storyStats(stringsArr)
		fmt.Println("Shortest: ", shortest)
		fmt.Println("Longest: ", longest)
		fmt.Println("Average: ", average)
	}
}

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

/**
 * Calculate the average of number array
 */
func averageNumber(arr []int64) float64 {
	var sum int64 = 0
	length := len(arr)

	for i := 0; i < length; i++ {
		sum += arr[i]
	}

	return (float64(sum)) / (float64(length))
}

/**
 * Whole story of a string array
 */
func wholeStory(arr []string) string {
	return strings.Join(arr, " ")
}

/**
 * Story Status (Shortest, Longest, Average, List equal to the slice element)
 */
func storyStats(arr []string) (shortest string, longest string, average float64) {
	temp := len(arr[0])
	var wordLengths = 0
	for _, word := range arr {
		wordLength := len(word)

		// Shortest
		if wordLength <= temp {
			shortest = word
			temp = wordLength
		}

		// Longest
		if wordLength > len(longest) {
			longest = word
		}

		wordLengths += wordLength
	}

	// Calculate average word length
	average = (float64(wordLengths)) / (float64(len(arr)))

	return
}
