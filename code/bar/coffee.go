package bar

import (
	"fmt"
	"time"
)

type Coffee interface {
	Brew(orderID int) string
	Name() string
}

type Espresso struct{}

func NewEspresso() Espresso {
	return Espresso{}
}

func (e Espresso) Name() string { return "Espresso" }

func (e Espresso) Brew(orderID int) string {
	time.Sleep(200 * time.Millisecond)
	return fmt.Sprintf("Order #%d ready - %s ☕", orderID, e.Name())
}

type Cappuccino struct{}

func NewCappuccino() Cappuccino {
	return Cappuccino{}
}

func (c Cappuccino) Name() string { return "Cappuccino" }

func (c Cappuccino) Brew(orderID int) string {
	time.Sleep(300 * time.Millisecond)
	return fmt.Sprintf("Order #%d ready - %s ☕", orderID, c.Name())
}
