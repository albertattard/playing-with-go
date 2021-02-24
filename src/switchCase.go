package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runSwitchCaseDemo() {
	number := randomInt(10)
	fmt.Printf("The random number was %v\n", number)

	/* By default, the switch has a 'true' value to be matched */
	switch {
	case 1 == number:
		fmt.Println("We were lucky!! We got a 1")
	case number <= 5:
		fmt.Println("We got a number between 2 and 5.  Fallthrough is not enabled!!")
	case number == 6:
		fmt.Println("We got a 6!!")
		fallthrough
	case number <= 8:
		fmt.Println("We got a number between 6 and 8.  Fallthrough is enabled!!")
	default:
		fmt.Println("We got a 9")
	}
}

func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
