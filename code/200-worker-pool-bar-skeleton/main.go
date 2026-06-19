package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println("Start")
	defer func() {
		fmt.Printf("End in %v ms\n", time.Now().Sub(t0).Milliseconds())
	}()

	/**
	Input: lots of orders
	Output: lots of ☕and some logEntries of our coffee machines (string)

	Deliverable: a log flow for each machine with start-end timestamp for each order

	Constraints: we have only 3 coffee machines (workers)
	*/

}
