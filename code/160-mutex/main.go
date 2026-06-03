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

	orders := make(bar.Orders, 0, 1000)

	for j := 0; j < 1000; j++ {
		go addOrder(bar.NewOrder(j, bar.NewEspresso()), &orders, &m)
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("orders size = %d\n", len(orders))
	fmt.Printf("End in %v\n", time.Since(t0).Seconds())
}

func addOrder(order bar.Order, list *bar.Orders, m *sync.Mutex) {
	m.Lock()
	*list = append(*list, order)
	m.Unlock()
}
