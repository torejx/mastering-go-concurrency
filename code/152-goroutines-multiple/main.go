package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	names := []string{"Mike", "Mark", "Anne", "Eve", "Bob", "Alice"}

	go sayHi(names)
	go sayHowAreYou(names)

	time.Sleep(4 * time.Second)

	fmt.Println("End")
}

func sayHi(names []string) {
	for _, name := range names {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Hi %v!\n", name)
	}
}

func sayHowAreYou(names []string) {
	for _, name := range names {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("\tHow are you, %v?\n", name)
	}
}
