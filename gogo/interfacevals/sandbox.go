package main

import (
	"fmt"
	"bytes"
	"io"
	"os"
)

func main() {
	var w1 io.Writer = os.Stdout
	var w2 io.Writer

	f1, ok := w1.(*os.File) // success:  ok, f == os.Stdoutb, ok := w.(*bytes.Buffer) // failure: !ok, b == nil

	fmt.Printf("{zzz: %T}\n", f1)
	fmt.Printf("{aaa: %t}\n", ok)

	f2, ok := w1.(*bytes.Buffer) // failure: !ok, b == nil

	fmt.Printf("{zzz: %T}\n", f2)
	fmt.Printf("{aaa: %t}\n", ok)


	f3, ok := w2.(*bytes.Buffer) 

	fmt.Printf("{zzz: %T}\n", f3)
	fmt.Printf("{aaa: %t}\n", ok)

	_, err := os.Open("/no/such/file")
	fmt.Printf("{Open Error: %t}\n", err)
	fmt.Println(os.IsNotExist(err)) // "true"

	//err = err.(type)
	do(21)
	do("hello")
	do(err)



}

//func underlyingError(err error) error {
//	switch err := err.(type) {
//	case *PathError:
//		return err.Err
//	case *LinkError:
//		return err.Err
//	case *SyscallError:
//		return err.Err
//	}
//	return err
//}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
