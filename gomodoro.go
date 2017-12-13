package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"time"
)

const shortWorkMinutes int = 15
const longWorkMinutes int = 25
const shortPauseMinutes int = 5
const longPauseMinutes int = 10

var minutePtr *int
var debugPtr *bool
var shortWorkPtr *bool
var longWorkPtr *bool
var shortPausePtr *bool
var longPausePtr *bool

var minuteInt int

func init() {
	minutePtr = flag.Int("m", 5, "duration in minutes")
	debugPtr = flag.Bool("d", false, "debug mode (count seconds instead of minutes)")
	shortWorkPtr = flag.Bool("w", false, "short work (15 min)")
	longWorkPtr = flag.Bool("W", false, "long work (25 min)")
	shortPausePtr = flag.Bool("p", false, "short pause (5 min)")
	longPausePtr = flag.Bool("P", false, "long pause (10 min)")

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
	err := beeep.Notify("Gomodoro", "and done...")
	if err != nil {
		panic(err)
	}
}

func main() {
	if *debugPtr == true {
		fmt.Println("DEBUG MODE: counting seconds instead of minutes")
	}

	if *shortWorkPtr == true {
		minuteInt = shortWorkMinutes
	} else if *longWorkPtr == true {
		minuteInt = longWorkMinutes
	} else if *shortPausePtr == true {
		minuteInt = shortPauseMinutes
	} else if *longPausePtr == true {
		minuteInt = longPauseMinutes
	} else {
		minuteInt = *minutePtr
	}

	fmt.Printf("Set timer for %d minutes\n", minuteInt)

	setTimer(minuteInt)
}
