package main

import "fmt"

func main() {
	ch := make(chan int, 3) // buffer size => how many items I can write without blocking

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	val, ok := <-ch
	fmt.Printf("Value: %v, is valid: %v, len: %v, cap: %v\n", val, ok, len(ch), cap(ch))

	val, ok = <-ch
	fmt.Printf("Value: %v, is valid: %v, len: %v, cap: %v\n", val, ok, len(ch), cap(ch))

	val, ok = <-ch
	fmt.Printf("Value: %v, is valid: %v, len: %v, cap: %v\n", val, ok, len(ch), cap(ch))

	// no more values
	val, ok = <-ch
	// ok = false if channel is closed and all values have already been received
	fmt.Printf("Value: %v, is valid: %v, len: %v, cap: %v\n", val, ok, len(ch), cap(ch))
}
