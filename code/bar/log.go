package bar

import (
	"fmt"
	"time"
)

// LogEntry is the output produced by the coffee machine after a brew
type LogEntry struct {
	MachineID int
	Start     time.Time
	End       time.Time
	Product   string
	Order     int
}

func (le LogEntry) String() string {
	msg := fmt.Sprintf("Order #%d ready - %s ☕", le.Order, le.Product)
	return fmt.Sprintf("-- Machine #%d \t %s - %s \t[%v] --",
		le.MachineID, le.Start.Format("15:04:05.000"), le.End.Format("15:04:05.000"), msg)
}
