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

	var m sync.Mutex
	var wg sync.WaitGroup

	orders := make(bar.Orders, 0, 1000)

	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go addOrder(bar.NewOrder(j, bar.NewEspresso()), &orders, &m, &wg)
	}

	wg.Wait()

	fmt.Printf("orders size = %d\n", len(orders))
	fmt.Printf("End in %v\n", time.Since(t0).Seconds())
}

func addOrder(order bar.Order, list *bar.Orders, m *sync.Mutex, wg *sync.WaitGroup) {
	// defer wg.Done()
	m.Lock()
	*list = append(*list, order)
	m.Unlock()
}
