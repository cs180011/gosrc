// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	ch <- "Start"
	go fetch("first", ch)
	go fetch("second", ch)
	go fetch("third", ch)
	go fetch("fourth", ch)
	fmt.Println(<-ch) // receive from channel ch
	fmt.Println(<-ch) // receive from channel ch
	fmt.Println(<-ch) // receive from channel ch
	fmt.Println(<-ch) // receive from channel ch
	fmt.Println(<-ch) // receive from channel ch
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	ch <- url
	time.Sleep(100 * time.Millisecond)
}
