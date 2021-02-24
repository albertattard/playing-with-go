package main

import "fmt"

func runArraysAndSlicesDemo() {
	a := []string{"a", "b", "c"}
	fmt.Printf("a=%v\n", a)

	b := append(a, "d")
	fmt.Printf("b=%v\n", b)

	a[0] = "A"
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
}
