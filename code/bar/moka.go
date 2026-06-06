package bar

import (
	"time"
)

// Moka represents a moka, of course!
type Moka struct {
	ID int
}

func NewMoka(id int) Moka {
	return Moka{ID: id}
}

func (m Moka) setup() {
	time.Sleep(1500 * time.Millisecond)
}

// Brew a coffee and return the operation log
func (m Moka) Brew(order Order) LogEntry {
	m.setup()

	t0 := time.Now()
	time.Sleep(order.Coffee.BrewTime())
	res := LogEntry{
		MachineID: m.ID,
		Start:     t0,
		End:       time.Now(),
		Product:   order.Coffee.Name(),
		Order:     order.ID,
	}

	return res
}
