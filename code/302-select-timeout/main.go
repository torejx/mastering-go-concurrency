package main

import (
	"fmt"
	"time"

	"github.com/torejx/mastering-go-concurrency/code/bar"
)

func main() {

	order := bar.NewOrder(1, bar.NewEspresso())

	moka := bar.NewMoka(1)

	ch := make(chan bar.LogEntry, 1)

	go brew(ch, moka, order)

	select {
	case <-ch:
		fmt.Println("Coffee is ready!")
	case <-time.After(1800 * time.Millisecond):
		fmt.Println("Too slow, the customer has gone :(")
	}

}

func brew(ch chan bar.LogEntry, br bar.Brewer, order bar.Order) {
	ch <- br.Brew(order)
}
