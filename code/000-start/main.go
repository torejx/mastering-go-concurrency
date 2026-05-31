package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	run()

	fmt.Println("End")
}

func run() {
	fmt.Println("\tGo func start!")

	time.Sleep(1 * time.Second)

	fmt.Println("\tGo func end!")
}
