package main

import (
	"fmt"
)

func printer(){
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func main(){
	go printer() // we dont get anything printed, because main func finished before printer started
}