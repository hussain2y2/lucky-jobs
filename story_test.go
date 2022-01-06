package main

import (
	"fmt"
	"testing"
)

func TestWholeStory(t *testing.T) {
	result := wholeStory([]string{"Hello", "world"})
	desired := "Hello world"

	if result != desired {
		t.Error(fmt.Printf("result %v, desired %v", result, desired))
	}
}

func TestStoryStats(t *testing.T) {
	// Got
	shortest, longest, average, list := storyStats([]string{"ab", "caba", "haha"})

	// Desired
	short, long, avg, li := "ab", "caba", 3.3333333333333335, []string{"caba", "haha"}

	if shortest != short {
		t.Error(fmt.Printf("result %v, desired %v", shortest, short))
	}

	if longest != long {
		t.Error(fmt.Printf("result %v, desired %v", longest, long))
	}

	if average != avg {
		t.Error(fmt.Printf("result %v, desired %v", average, avg))
	}

	if len(list) != len(li) {
		t.Error(fmt.Printf("result %v, desired %v", list, li))
	}
}
