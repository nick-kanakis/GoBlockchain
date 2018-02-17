package utils

import "testing"

func TestConcatStringsMultipleInputs(t *testing.T) {
	concatString := ConcatStrings("one", "two", "tree")
	expectedResult := "onetwotree"
	if concatString != expectedResult {
		t.Error("Concatenation of strings failed")
	}
}

func TestConcatStringsSingleInputs(t *testing.T) {
	concatString := ConcatStrings("one")
	expectedResult := "one"
	if concatString != expectedResult {
		t.Error("Concatenation of one string failed")
	}
}
