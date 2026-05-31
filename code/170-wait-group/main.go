package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()

	fmt.Println("Start")

	i := 0

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go increment(&i, &mutex, &wg)
	}

	wg.Wait()

	fmt.Printf("i = %d\n", i)
	fmt.Printf("End in %v ms\n", time.Now().Sub(t0).Milliseconds())
}

func increment(i *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*i = *i + 1
	mutex.Unlock()
}
