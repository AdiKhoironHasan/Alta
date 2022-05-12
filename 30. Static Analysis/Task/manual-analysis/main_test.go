package main

import "testing"

func TestPassed(t *testing.T) {
	if passed() != true {
		t.Error("passed test incorrect")
	}
}
