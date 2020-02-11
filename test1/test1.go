package main

import (
	"fmt"
)

// Main is the main entry point
func main() {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("Hello World")

	fmt.Printf("Returned String = %s.\n", T1())
}

// T1 is my first routine to test!
func T1() string {
	t1str := "sample"
	return t1str
}
