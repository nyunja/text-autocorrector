package main

import (
	"testing"

	"goreloaded/functions"
)

/*
Unit test for the FinalizedText function
*/

func TestFixText(t *testing.T) {
	strSlice := "I was thinking ... You were right"
	expected := "I was thinking... You were right"
	result := functions.FinalizedText([]byte(strSlice))

	if string(result) != expected {
		t.Errorf("Error: got [%s] wanted [%s]", result, expected)
	}
}
