package main

import (
	"fmt"
)

func printer(x int){
	fmt.Println("HELLO ", x) // prints in disorder due to go scheduler
}

func main(){
	ch := make(chan struct{}) // empty struct, just notifying

	for  i := 0; i < 10; i++{
		go func(x int) {
			printer(x)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < 10; i++ {
		<- ch // waiting mechanisms
	}
	close(ch)
}