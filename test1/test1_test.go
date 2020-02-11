package main

import (
	"fmt"
	"os"
	"testing"
)

// TestMain is called by go test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestT1(t *testing.T) {
	got := T1()
	if got != "sample" {
		t.Errorf("T1() = %s; want sample", got)
	}
}

func TestT1b(t *testing.T) {
	got := T1()
	t.Skip()
	if got != "SAMPLE" {
		t.Errorf("T1() = %s; want SAMPLE", got)
	}
}

func BenchmarkT1(t *testing.B) {
	T1()
}

func ExampleT1() {
	// Output:
	// sample output
	fmt.Printf("%s output", T1())
	// Output:
	// sample output
}
