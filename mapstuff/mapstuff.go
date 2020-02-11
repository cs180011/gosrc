// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ages := make(map[string]int) // mapping from strings to ints
	fmt.Printf("%d\n", len(ages))

	ages["a"] = 1
	fmt.Printf("%d\n", len(ages))

	newages := map[string]int{
		"alice": 31, "charlie": 34}
	fmt.Printf("%d\n", len(newages))

	_, ok := ages["bob"]
	if !ok {
		fmt.Printf("Bob Not There.\n")
		/* "bob" is not a key in this map; age == 0. */
	}

	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("Bob Not There; age=%d\n", age)
		/* "bob" is not a key in this map; age == 0. */
	}

	_, _ = ages["bob"]

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
