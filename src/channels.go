package main

import (
	"fmt"
	"sync"
)

func runBasicChannelDemo() {
	bufferSize := 1
	channel := make(chan int, bufferSize)

	fmt.Println("Sending 1 to the channel")
	channel <- 1

	value := <-channel
	fmt.Printf("Got %v from the channel\n", value)
}

func runReadFromCloseChannelDemo() {
	bufferSize := 1
	channel := make(chan int, bufferSize)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	/* Write one number and close the channel */
	go func() {
		channel <- 42
		/* The channel can be closed once we are done writing to it.  We will still be able to read from the channel. */
		close(channel)
		waitGroup.Done()
	}()

	/* Wait for the channel to be closed first */
	fmt.Println("Waiting for the channel to be closed")
	waitGroup.Wait()

	for i := 0; i < 2; i++ {
		value, ok := <-channel
		if ok {
			fmt.Printf("Got %v from the channel\n", value)
		} else {
			fmt.Printf("Failed to read from channel!! %v\n", ok)
		}
	}
}

func runReadWriteChannelDemo() {
	bufferSize := 1
	channel := make(chan int, bufferSize)

	var readChannel <-chan int
	var writeChannel chan<- int

	readChannel = channel
	writeChannel = channel

	fmt.Println("Sending 1 to the (write) channel")
	writeChannel <- 1

	value := <-readChannel
	fmt.Printf("Got %v from the (read) channel\n", value)
}

func runChannelWithTwoGoroutinesDemo() {
	bufferSize := 1
	channel := make(chan int, bufferSize)

	go func() {
		sendNumberToChannel(channel, 1)
		sendNumberToChannel(channel, 2)
		sendNumberToChannel(channel, 3)
		/* This will block in the line above and will never print the following, as we are taking just one from the channel and we have a buffer of 1 */
		fmt.Println("Done sending things to the channel")
	}()

	value := readNumberFromChannel(channel)
	fmt.Printf("Got %v from the channel\n", value)
}

func runChannelWithManyGoroutinesDemo() {
	bufferSize := 1
	channel := make(chan int, bufferSize)

	go func() {
		numberOfGoroutines := 100
		var waitGroup sync.WaitGroup
		waitGroup.Add(numberOfGoroutines)

		for i := 0; i < 100; i++ {
			number := i + 1
			go func() {
				sendNumberToChannel(channel, number)
				waitGroup.Done()
			}()
		}

		/* Wait for all the other Goroutines to finish writing to the channel */
		waitGroup.Wait()
		close(channel)
	}()

	readNumbersUntilChannelIsClosed(channel)
}

/* The function takes a write channel which means it can only write to the channel and cannot read from the channel */
func sendNumberToChannel(channel chan<- int, number int) {
	fmt.Printf("Sending %v to the channel\n", number)
	channel <- number
}

/* The function takes a read channel which means it can only read from the channel and cannot write to the channel */
func readNumberFromChannel(channel <-chan int) int {
	return <-channel
}

/* The function takes a read channel which means it can only read from the channel and cannot write to the channel */
func readNumbersUntilChannelIsClosed(channel <-chan int) {
	fmt.Println("Reading from the (read) channel until this is closed")
	for number := range channel {
		fmt.Printf("Got %v from the channel\n", number)
	}
	fmt.Println("Read all numbers form the (read) channel")
}
