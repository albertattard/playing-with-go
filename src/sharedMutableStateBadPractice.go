package main

import (
	"fmt"
	"sync"
)

/*
 * Run this code using 'go run -race file-name.go' to see the race warning too.  See sharedMutableStateMutex.go for a better example.
 */
func runUnprotectedSharedMutableStateDemo() {
	numberOfGoroutines := 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(numberOfGoroutines)

	var sharedMutableCounter int

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			/* The following make the problem more visible */
			//v := sharedMutableCounter
			//runtime.Gosched() /* This is like yield */
			//sharedMutableCounter = v + 1

			sharedMutableCounter++
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	if numberOfGoroutines == sharedMutableCounter {
		fmt.Printf("We were lucky this time as the value of the shared mutable state is %v as expected\n", sharedMutableCounter)
	} else {
		fmt.Printf("The value of the shared mutable state is %v instead of the expected value of %v\n", sharedMutableCounter, numberOfGoroutines)
	}

}
