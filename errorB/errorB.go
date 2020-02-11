package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// ReallyNil demonstrates a value that is really returned as nil
func ReallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)
	return e
}

// MyError struct imlments the error interface in the error package
type MyError struct {
	A       int
	B       int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("Values %d and %d produced error %s", me.A, me.B, me.Message)
}

// NotReallyNil demonstrates the interface nil versus not-nil
func NotReallyNil() error {
	var me *MyError
	fmt.Println("me is nil:", me == nil)
	return me
}

func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		//return 0, 0, errors.New("Cannot Divide By Zero")
		return 0, 0, &MyError{A: a, B: b, Message: "Cannot Divide By Zero"}
	}
	return a / b, a % b, nil
}

func main() {

	e := ReallyNil()
	me := NotReallyNil()
	fmt.Println("in main(), e is nil:", e == nil)
	fmt.Println("in main(), me is nil:", me == nil)

	meV := reflect.ValueOf(me)
	meVT := meV.Type()
	meT := reflect.TypeOf(me)
	fmt.Println("meV and meT :", meV, meVT, meT)
	fmt.Printf("meVT = %T\n", meVT)

	eV := reflect.ValueOf(e)
	eVT := eV.Type()
	eT := reflect.TypeOf(e)
	fmt.Println("eV and eT :", eV, eT)
	fmt.Printf("eVT = %T\n", eVT)

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
