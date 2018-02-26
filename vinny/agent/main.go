package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	Fname = flag.String("name", "", "")
)

func main() {
	flag.Parse()
	for {
		fmt.Printf("%s: Sleeping 20 seconds...\n", *Fname)
		time.Sleep(20 * time.Second)
	}
}
