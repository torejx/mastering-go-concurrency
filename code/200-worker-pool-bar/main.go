package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/torejx/mastering-go-concurrency/code/bar"
)

func main() {
	t0 := time.Now()
	fmt.Printf("Start\n")
	defer func() {
		fmt.Printf("End in %v ms\n", time.Now().Sub(t0).Milliseconds())
	}()

	in := make(chan bar.Order, 10) // choose a proper buffer size
	out := make(chan string, 10)

	workers := 3 // number of concurrent executions

	wg := sync.WaitGroup{}

	// consume output
	go func() {
		for o := range out {
			fmt.Printf("-- %v --\n", o)
		}
	}()

	// start workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		// worker function
		go func() {
			defer wg.Done()
			for v := range in {
				out <- fmt.Sprintf("Machine #%d: %v", i, v.Coffee.Brew(v.ID))
			}
		}()
	}

	// feed the input
	for i := 0; i < 999; i++ { // what happens if I increase the limit?
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(2)
		if k == 0 {
			in <- bar.NewOrder(i, bar.NewEspresso())
		} else {
			in <- bar.NewOrder(i, bar.NewCappuccino())
		}
	}

	// clone the input channel
	close(in)

	// wait for workers
	wg.Wait()

	// close output
	close(out)
}
