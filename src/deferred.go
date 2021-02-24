package main

import "fmt"

func runDeferredDemo() {
	deferFunctionCallByValue()
}

func deferFunctionCallByValue() {
	i := 0
	defer fmt.Printf("The value of i when called by the deferred print is %v\n", i)
	i++
	fmt.Printf("The value of i is %v\n", i)
}
