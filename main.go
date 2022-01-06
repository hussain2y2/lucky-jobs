package main

import (
	"fmt"
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

		shortest, longest, average, list := storyStats(stringsArr)

		fmt.Println("Shortest: ", shortest)
		fmt.Println("Longest: ", longest)
		fmt.Println("Average: ", average)
		fmt.Println("List of elements equal to the average length: ", list)
	}
}
