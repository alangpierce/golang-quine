package main

import (
	"testing"
	"io/ioutil"
)

func TestQuine(t *testing.T) {
	fileContents, err := ioutil.ReadFile("./quine.go")
	if err != nil {
		t.Fail()
		return
	}
	expectedCode := string(fileContents)
	actualCode := getCode()

	firstDifference := findFirstDifference(expectedCode, actualCode)
	if firstDifference >= 0 {
		t.Error("Strings differ.\n>>> Common prefix:\n",
			expectedCode[:firstDifference],
			"\n\n>>> Expected remaining:\n",
			expectedCode[firstDifference:],
			"\n\n>>> Actual remaining:\n",
			actualCode[firstDifference:])
	}
}

// findFirstDifference returns the index of the first differing rune in the two
// strings, or -1 if the strings are equal.
func findFirstDifference(s1 string, s2 string) int {
	for i := 0; i < max(len(s1), len(s2)); i++ {
		if i >= len(s1) || i >= len(s2) || s1[i] != s2[i] {
			return i
		}
	}
	return -1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}