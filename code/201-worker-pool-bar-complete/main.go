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
	fmt.Println("Start")
	defer func() {
		fmt.Printf("End in %v ms\n", time.Now().Sub(t0).Milliseconds())
	}()

	in := make(chan bar.Order, 10) // choose a proper buffer size
	out := make(chan bar.LogEntry, 10)

	numberOfCoffeeMachines := 3

	wg := sync.WaitGroup{}
	outWg := sync.WaitGroup{}

	// consume output
	outWg.Add(1)
	go func() {
		defer outWg.Done()
		for o := range out {
			fmt.Println(o)
		}
	}()

	// start coffee machines!!!
	for i := 0; i < numberOfCoffeeMachines; i++ {
		wg.Add(1)
		// worker function
		go func(machineID int) {
			defer wg.Done()
			machine := bar.NewMachine(machineID)

			for order := range in {
				out <- machine.Brew(order)
			}
		}(i)
	}

	// feed the input
	for i := 0; i < 150; i++ { // what happens if I increase the limit?
		k := rand.Intn(2)
		if k == 0 {
			in <- bar.NewOrder(i, bar.NewEspresso())
		} else {
			in <- bar.NewOrder(i, bar.NewCappuccino())
		}
	}

	// close the input channel
	close(in)

	// wait for workers
	wg.Wait()

	// close output
	close(out)

	// wait for output worker
	outWg.Wait()
}
