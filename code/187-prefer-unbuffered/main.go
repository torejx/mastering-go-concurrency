package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println("Start")
	defer fmt.Printf("End in %v\n", time.Since(t0).Seconds())

	wg := sync.WaitGroup{}
	ch := make(chan int) // what if I set 99999?

	for i := range 1000 {
		ch <- i
	}

	close(ch)

	wg.Add(1)
	go printFromChannel(ch, &wg)
	wg.Wait()

}

func printFromChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
