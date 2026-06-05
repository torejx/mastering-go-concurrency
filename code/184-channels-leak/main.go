package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go double(2, ch)
	go double(4, ch)

	fmt.Printf("value: %d\n", <-ch)

	close(ch)
}

func double(i int, ch chan int) {
	ch <- i * 2
}
