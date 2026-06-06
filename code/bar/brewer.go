package bar

// Brewer represents a generic item that Brews coffee.
type Brewer interface {
	Brew(order Order) LogEntry
}
