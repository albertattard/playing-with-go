package main

import (
	"fmt"
	"runtime"
	"sync"
)

func runGoroutinesDemo() {
	printGoroutinesEnv()

	//var waitGroup sync.WaitGroup
	waitGroup := sync.WaitGroup{}
	runInGoroutine(&waitGroup, func() { printNTimes("A", 10) })
	runInGoroutine(&waitGroup, func() { printNTimes("B", 10) })

	printGoroutinesEnv()

	fmt.Println("Waiting for all goroutines to finish")
	waitGroup.Wait()
	fmt.Println("Done")
}

func runInGoroutine(waitGroup *sync.WaitGroup, fn func()) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		fn()
	}()
}

func printGoroutinesEnv() {
	fmt.Printf("OS:        %v\n", runtime.GOOS)
	fmt.Printf("Arch:      %v\n", runtime.GOARCH)
	fmt.Printf("CPU:       %v\n", runtime.NumCPU())
	fmt.Printf("Goroutine: %v\n", runtime.NumGoroutine())
}

func printNTimes(s string, times uint) {
	for i := uint(0); i < times; i++ {
		fmt.Printf("%v (%d)\n", s, i+1)
	}
}
