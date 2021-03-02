package main

import (
        "fmt"
        "os"
		"flag"
)

func main() {

        timezone := os.Getenv("TZ")
        fmt.Printf("Timezone: %v\n", timezone)
		
		var port = flag.Int("port", 9000, "Port number")
		flag.Parse()
		fmt.Printf("Port: %v\n", *port) // port is a pointer
		// NewYork tags are useless
}
