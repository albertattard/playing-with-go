package main

import "fmt"

func runChannelSelectDemo() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	completeChannel := make(chan int)

	fmt.Println("Sending numbers to channels")
	go sendToChannels(evenChannel, oddChannel, completeChannel)

	fmt.Println("Reading numbers form channels using the select")
	printNumbersReadFromChannels(evenChannel, oddChannel, completeChannel)
}

func runChannelFanInExample() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	completeChannel := make(chan int)

	fmt.Println("Sending numbers to channels")
	go sendToChannels(evenChannel, oddChannel, completeChannel)

	fanInChannel := make(chan int)

	fmt.Println("Fan in from the other channels")
	go fanInChannels(evenChannel, oddChannel, completeChannel, fanInChannel)

	/* Trying to read from a drained channel will yield an error.  I don't understand why!!  It seems that the writer blocks, but the reader seems to fails when reading from a drained channel. */
	for number := range fanInChannel {
		fmt.Printf("Got %v from the fan in channel\n", number)
	}
}

func printNumbersReadFromChannels(evenChannel, oddChannel, completeChannel <-chan int) {
	for {
		select {
		case number := <-evenChannel:
			fmt.Printf("Recevied %v from the even channel\n", number)
		case number := <-oddChannel:
			fmt.Printf("Recevied %v from the odd channel\n", number)
		case _ = <-completeChannel:
			fmt.Println("Complete!!")
			return
		}
	}
}

func fanInChannels(evenChannel, oddChannel, completeChannel <-chan int, fanInChannel chan<- int) {
	/* Close the channel when done writing.  We can still read values from this channel. */
	defer close(fanInChannel)

	for {
		select {
		case number := <-evenChannel:
			fanInChannel <- number
		case number := <-oddChannel:
			fanInChannel <- number
		case _ = <-completeChannel:
			return
		}
	}
}

func sendToChannels(evenChannel, oddChannel, completeChannel chan<- int) {
	/* Close the channels when done writing.  We can still read values from these channels. */
	defer close(evenChannel)
	defer close(oddChannel)
	defer close(completeChannel)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}
	completeChannel <- 0
}
