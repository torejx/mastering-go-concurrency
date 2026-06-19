package main

import (
	"fmt"

	"github.com/torejx/mastering-go-concurrency/code/bar"
)

func main() {

	order := bar.NewOrder(1, bar.NewEspresso())

	machine := bar.NewMachine(1)
	moka := bar.NewMoka(2)

	mokaCh := make(chan bar.LogEntry, 1) // why 1?
	machineCh := make(chan bar.LogEntry, 1)

	go brew(mokaCh, moka, order)
	go brew(machineCh, machine, order)
	select {
	case <-mokaCh:
		fmt.Println("Moka is faster!")
	case <-machineCh:
		fmt.Println("Machine is faster!")
	}
}

func brew(ch chan bar.LogEntry, br bar.Brewer, order bar.Order) {
	ch <- br.Brew(order)
}
