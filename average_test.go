package main

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	got := averageNumber([]int64{23, 24, 25})
	want := 24

	if got != float64(want) {
		t.Error(fmt.Printf("got %v, want %v", got, want))
	}
}
