package main

import (
	"flag"
	"fmt"
	"time"
)

var durationPtr *int

func init() {
	durationPtr = flag.Int("mg", 5, "duration in minutes")
	flag.Parse()
}

func main() {
	fmt.Printf("Set timer for %d minutes\n", *durationPtr)

	// channel not really needed
	c1 := make(chan string, 1)
	go func() {
		durationInt := *durationPtr
		time.Sleep(time.Minute * time.Duration(durationInt))
		c1 <- "and done.."
	}()

	res := <-c1
	fmt.Println(res)
}
