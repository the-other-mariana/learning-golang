package main

import (
	"fmt"
)

// counter squares and main are running at the same time

func counter(naturals chan int) {
	for x := 0; x < 100; x++ {
		naturals <- x
	}
	close(naturals)
}

func squarer(naturals, squares chan int) {
	// at the moment the natural channel closes, the for stops
	for x := range naturals { // consume the queue elements, if u have two funcs reading naturals, values dont get repeated
		squares <- x * x
	}
	close(squares)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go counter(naturals)

	// Squarer
	go squarer(naturals, squares)

	// Printer (in main goroutine)
	// at the moment the squares  channel is closed (in squarer) the for stops
	for x := range squares {
		fmt.Println(x)
	}
	// if u put close channel here it doesnt work
}