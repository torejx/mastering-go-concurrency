package bar

import (
	"time"
)

// Machine represents a bar offee machine
type Machine struct {
	ID int
}

func NewMachine(id int) Machine {
	return Machine{ID: id}
}

// autoClean run after each brew
func (m Machine) autoClean() {
	time.Sleep(100 * time.Millisecond)
}

// Brew a coffee and return the operation log
func (m Machine) Brew(order Order) LogEntry {
	t0 := time.Now()
	time.Sleep(order.Coffee.BrewTime())
	res := LogEntry{
		MachineID: m.ID,
		Start:     t0,
		End:       time.Now(),
		Product:   order.Coffee.Name(),
		Order:     order.ID,
	}

	m.autoClean()

	return res
}
