package main

import (
	"fmt"
	"testing"
)

func TestTestValidity(t *testing.T) {
	got := testValidity("12-ab-13-bc")
	desired := true

	if got != desired {
		t.Error(fmt.Printf("got %v, desired %v", got, desired))
	}
}
