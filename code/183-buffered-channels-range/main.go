package main

import "fmt"

func main() {
	ch := make(chan int, 3) // buffer size => how many items I can write without blocking

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for val := range ch {
		fmt.Printf("Value: %v\n", val)
	}
}
