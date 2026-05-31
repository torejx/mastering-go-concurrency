package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println("Start")

	i := 0

	for j := 0; j < 1000; j++ {
		go increment(&i)
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("i = %d\n", i)
	fmt.Printf("End in %v\n", time.Since(t0).Seconds())
}

func increment(i *int) {
	*i = *i + 1
}
