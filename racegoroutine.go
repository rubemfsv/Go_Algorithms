// Write two goroutines which have a race condition when executed concurrently.
//  Explain what the race condition is and how it can occur.

package main

import (
	"fmt"
	"time"
)

var count int

func race() {
	fmt.Println(count)
	count++
}

func main() {

	count = 1
	go race()
	go race()
	time.Sleep(1 * time.Second)
}