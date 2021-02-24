package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func runAtomicSharedMutableStateDemo() {
	numberOfGoroutines := 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(numberOfGoroutines)

	var sharedMutableCounter int64

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			atomic.AddInt64(&sharedMutableCounter, 1)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Printf("The value of the shared mutable state is %v es expected\n", sharedMutableCounter)
}
