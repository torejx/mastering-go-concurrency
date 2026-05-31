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

	for j := 0; j < 100; j++ {
		go increment(&i, &mutex)
	}

	time.Sleep(4 * time.Second)

	fmt.Printf("i = %d\n", i)
	fmt.Printf("End in %v\n", time.Now().Sub(t0).Seconds())
}

func increment(i *int, mutex *sync.Mutex) {
	mutex.Lock()
	*i = *i + 1
	mutex.Unlock()
}
