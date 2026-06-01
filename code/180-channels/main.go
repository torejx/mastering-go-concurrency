package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go writeIntoChannel(ch, 1)

	fmt.Printf("The value is: %v\n", <-ch)

}

func writeIntoChannel(ch chan int, value int) {
	ch <- value
	close(ch) // we can't write into a closed channel, but we can read from it
}
