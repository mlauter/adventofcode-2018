package main

import (
	"testing"
)

func TestReducePolymer(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	expected := "dabCBAcaDA"

	got := reducePolymer(input)

	if got != expected {
		t.Errorf("day05 reduce polymer failed. Expected %s got %s", expected, got)
	}
}
