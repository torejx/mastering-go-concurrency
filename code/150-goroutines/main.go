package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	go run() // notice the "go" keyword

	time.Sleep(1 * time.Second) // guess why

	fmt.Println("End")
}

func run() {
	fmt.Println("\tGo routine is running!")
}
