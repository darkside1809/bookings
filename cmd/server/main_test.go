package main

import (
	"testing"
)

func TestExecute(t *testing.T) {
	_, err := execute()
	if err != nil {
		t.Error("Failed execute")
	}
}