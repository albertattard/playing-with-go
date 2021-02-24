package main

import "fmt"

func printPointers() {
	a := 42
	b := a

	fmt.Println("Created two numbers")
	fmt.Printf("The value of %T 'a' is %v\n", a, a)
	fmt.Printf("The value of %T 'b' is %v\n", b, b)

	b = 24
	fmt.Println()
	fmt.Println("Change variable 'b'")
	fmt.Printf("The value of %T 'a' is %v\n", a, a)
	fmt.Printf("The value of %T 'b' is %v\n", b, b)

	c := &a
	fmt.Println()
	fmt.Println("Create a pointer 'c' pointing to 'a'")
	fmt.Printf("The value of %T 'a' is %v\n", a, a)
	fmt.Printf("The value of %T 'b' is %v\n", b, b)
	fmt.Printf("The value of %T 'c' is %v (holding %v)\n", c, c, *c)

	*c = 72
	fmt.Println()
	fmt.Println("Changing variable 'c' (which is pointing to 'a')")
	fmt.Printf("The value of %T 'a' is %v\n", a, a)
	fmt.Printf("The value of %T 'b' is %v\n", b, b)
	fmt.Printf("The value of %T 'c' is %v (holding %v)\n", c, c, *c)
}
