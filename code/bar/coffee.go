package bar

import (
	"time"
)

type Coffee interface {
	Name() string
	BrewTime() time.Duration
}

type Espresso struct{}

func NewEspresso() Espresso {
	return Espresso{}
}

func (e Espresso) Name() string { return "Espresso" }

func (e Espresso) BrewTime() time.Duration {
	return 200 * time.Millisecond
}

type Cappuccino struct{}

func NewCappuccino() Cappuccino {
	return Cappuccino{}
}

func (c Cappuccino) Name() string { return "Cappuccino" }

func (c Cappuccino) BrewTime() time.Duration {
	return 300 * time.Millisecond
}
