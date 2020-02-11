// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%d\n", len(a))

	//var []int b := nil
	var b []int
	fmt.Printf("%d\n", len(b))

	b = []int{}
	fmt.Printf("%d\n", len(b))

	c := a[3:6]
	fmt.Printf("c= %d\n", len(c))

	slice := make([]int, 30)
	fmt.Printf("%d\n", len(slice))
	fmt.Printf("%d\n", cap(slice))

	slice = append(slice, 30)
	fmt.Printf("%d\n", len(slice))
	fmt.Printf("%d\n", cap(slice))

	fmt.Println(slice)

	//type thisslice slice(10,10)
	s := make([]int, 20, 40)
	fmt.Printf("%d\n", len(s))

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
