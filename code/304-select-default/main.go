package main

import (
	"fmt"
	"time"

	"github.com/torejx/mastering-go-concurrency/code/bar"
)

func main() {
	in := make(chan bar.Order, 5)
	machine := bar.NewMachine(1)

	// feed some orders with a delay to simulate a real scenario
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * 300 * time.Millisecond)
			in <- bar.NewOrder(i, bar.NewEspresso())
		}
		close(in)
	}()

	for {
		select {
		case order, ok := <-in:
			if !ok {
				fmt.Println("no more orders, shutting down 🔌")
				return
			}
			fmt.Println(machine.Brew(order))
		default:
			fmt.Println("no orders at the moment, machine is idle 💤")
			time.Sleep(100 * time.Millisecond) // avoid busy loop
		}
	}
}
