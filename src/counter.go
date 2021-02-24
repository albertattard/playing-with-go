package main

func intCounter() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}
