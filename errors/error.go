package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot Divide By Zero")
	}
	return a / b, a % b, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Expected 2 Input Parameters")
		os.Exit(1)
	}

	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	div, mod, err := divAndMod(int(a), int(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(div, mod)
	fmt.Println(div, mod)

}