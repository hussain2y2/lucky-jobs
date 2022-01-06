package main

import (
	"math"
	"strings"
)

/**
 * Whole story of a string array
 */
func wholeStory(arr []string) string {
	return strings.Join(arr, " ")
}

/**
 * Story Status (Shortest, Longest, Average, List equal to the slice element)
 */
func storyStats(arr []string) (shortest string, longest string, average float64, list []string) {
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

	// List the same length as with average round up/round down
	roundUp := math.Ceil(average)
	roundDown := math.Floor(average)

	for _, word := range arr {
		wordLength := len(word)
		if wordLength == int(roundUp) || wordLength == int(roundDown) {
			list = append(list, word)
		}
	}

	return
}
