package main

import (
	"fmt"
	"testing"
)

func TestCheckNumber(t *testing.T) {
	result := checkNumber("123")
	desired := true

	if result != desired {
		t.Error(fmt.Printf("result %v, desired %v", result, desired))
	}
}

func TestCheckString(t *testing.T) {
	result := checkString("abc")
	desired := true

	if result != desired {
		t.Error(fmt.Printf("result %v, desired %v", result, desired))
	}
}
