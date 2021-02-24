package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func runChannelWithContextDemo() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("Error check before run %v\n", ctx.Err())
	fmt.Printf("Number of Goroutines %v\n", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 500)
				fmt.Printf("[%v] Working...\n", n)
			}
		}
	}()

	fmt.Println("Waiting few seconds...")
	time.Sleep(time.Second * 5)

	fmt.Printf("Error check while running %v\n", ctx.Err())
	fmt.Printf("Number of Goroutines %v\n", runtime.NumGoroutine())

	fmt.Println("Cancelling context")
	cancel()

	fmt.Printf("Error check after cancelling %v\n", ctx.Err())
	fmt.Printf("Number of Goroutines %v\n", runtime.NumGoroutine())
	fmt.Println("Done")
}
