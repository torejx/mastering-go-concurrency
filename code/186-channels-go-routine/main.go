package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/torejx/mastering-go-concurrency/code/bar"
)

func main() {
	t0 := time.Now()
	fmt.Println("Start")

	var wg sync.WaitGroup

	orders := make(bar.Orders, 0, 1000)
	orderCh := make(chan bar.Order)

	wg.Add(1)
	go addOrder(orderCh, &orders, &wg)

	for j := 0; j < 1000; j++ {
		orderCh <- bar.NewOrder(j, bar.NewEspresso())
	}

	close(orderCh)

	wg.Wait()

	fmt.Printf("orders size = %d\n", len(orders))
	fmt.Printf("End in %v\n", time.Since(t0).Seconds())
}

func addOrder(ch chan bar.Order, list *bar.Orders, wg *sync.WaitGroup) {
	defer wg.Done()
	for o := range ch {
		*list = append(*list, o)
	}
}
