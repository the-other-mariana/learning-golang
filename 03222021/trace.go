package main                                           

// https://making.pusher.com/go-tool-trace/

import (
        "os"
        "runtime/trace"
        "time"
        "fmt"
)

func fib(n int) int {
        if n <= 1 {
                return n
        }
        return fib(n-1) + fib(n-2)
}

func main() {
	f, err := os.Create("trace.out")
        if err != nil {
                panic(err)
	}
        defer f.Close()

        err = trace.Start(f)
	if err != nil {
		panic(err)
	}
        defer trace.Stop() // defer means 'at the end', but you can trace.Stop() in any line

	fmt.Printf("Hello World I'm starting")

	go fib(45)
	go fib(43)

	time.Sleep(10 * time.Second)

}
