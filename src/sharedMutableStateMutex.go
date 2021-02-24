package main

import (
	"fmt"
	"runtime"
	"sync"
)

func runMutexSharedMutableStateDemo() {
	numberOfGoroutines := 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(numberOfGoroutines)

	var sharedMutableCounter int

	var guard sync.Mutex

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			guard.Lock()
			defer guard.Unlock()
			v := sharedMutableCounter
			runtime.Gosched() /* This is like yield */
			sharedMutableCounter = v + 1
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Printf("The value of the shared mutable state is %v es expected\n", sharedMutableCounter)
}
