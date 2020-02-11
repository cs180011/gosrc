package main

import (
	"os"
	"testing"
)

// TestMain is called by go test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}