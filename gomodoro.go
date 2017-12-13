package main

import (
	"flag"
	"fmt"
	"time"
)

var minutePtr *int
var debugPtr *bool

func init() {
	minutePtr = flag.Int("m", 5, "duration in minutes")
	debugPtr = flag.Bool("d", false, "debug mode (count seconds instead of minutes)")
	flag.Parse()
}

func setTimer(minutes int) {
	for m := 1; m <= minutes; m++ {
		if *debugPtr == true {
			time.Sleep(time.Second)
		} else {
			time.Sleep(time.Minute)
		}
		if m == 1 {
			fmt.Println(" 1 minute has passed...")
		} else {
			fmt.Printf("%2d minutes have passed...\n", m)
		}
	}
	fmt.Println("and done..")
}

func main() {
	if *debugPtr == true {
		fmt.Println("DEBUG MODE: counting seconds instead of minutes")
	}
	fmt.Printf("Set timer for %d minutes\n", *minutePtr)
	minuteInt := *minutePtr
	setTimer(minuteInt)
}
