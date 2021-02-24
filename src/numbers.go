package main

func sum(numbers ...int) int {
	acceptAll := func(int) bool { return true }
	return sumIf(acceptAll, numbers...)
}

func sumIf(predicate func(int) bool, numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		if predicate(number) {
			sum += number
		}
	}

	return sum
}

func transform(function func(int) int, numbers []int) {
	for i, number := range numbers {
		numbers[i] = function(number)
	}
}

func apply(function func(int) int, numbers ...int) []int {
	for i, number := range numbers {
		numbers[i] = function(number)
	}
	return numbers
}

func factorial(number uint) uint {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}

func fibonacci(number uint) uint {
	if number <= 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	return fibonacci(number-1) + fibonacci(number-2)
}
