// code that prints the percentage of two goroutines channel content loading

package main

import (
	"fmt"
)

type Payload struct {
	ChannelID string
	Value     int
}

func aGo(pChan chan Payload) {
	inc := 0
	for {
		inc++
		pChan <- Payload{"a", inc}
	}
}

func bGo(pChan chan Payload) {
	inc := 0
	for {
		inc++
		pChan <- Payload{"b", inc}
	}
}

func main() {

	pChan := make(chan Payload)
	total := 0
	aReceives := 0
	bReceives := 0

	go aGo(pChan)
	go bGo(pChan)

	for p := range pChan {
		fmt.Printf("Channel ID: %v, Value: %5v\n", p.ChannelID, p.Value)
		if p.ChannelID == "a"{
			aReceives += 1
		}
		if p.ChannelID == "b"{
			aReceives += 1
		}
		total += 1
		if p.Value == 1000 {
			break
		}
	}
	
	fmt.Printf("Total: %v\n", total)
	fmt.Printf("A receives: %v \n", (aReceives / (total * 1.0) * 100))
	fmt.Printf("A receives: %v \n", (bReceives / (total * 1.0) * 100))

	close(pChan)
}