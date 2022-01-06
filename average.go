package main

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
