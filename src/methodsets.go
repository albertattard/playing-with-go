package main

import "fmt"

type device struct {
	name string
}

func (d device) String() string {
	return fmt.Sprintf("%v", d.name)
}

func (d device) valueMethodSet() {
	fmt.Printf("Device '%v' is printing...\n", d)
}

func (d *device) pointerMethodSet() {
	fmt.Printf("Device '%v' is beeping...\n", d)
}

func workWithValues(d device) {
	fmt.Printf("The argument '%v' is of type %T\n", d.name, d)
}

func workWithPointer(d *device) {
	fmt.Printf("The argument '%v' is of type %T\n", d.name, d)
}

func runMethodSetDemo() {
	/* Note that this works differently with interfaces and structs. See shapes.go for an example. */
	a := device{"a"}
	b := &device{"b (pointer)"}

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Variables")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("The variable '%v' is of type %T\n", a.name, a)
	fmt.Printf("The variable '%v' is of type %T\n", b.name, b)
	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Work with values")
	fmt.Println("--------------------------------------------------------")
	workWithValues(a)
	workWithValues(*b)
	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Work with pointers")
	fmt.Println("--------------------------------------------------------")
	workWithPointer(&a)
	workWithPointer(b)
	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Value Method Set")
	fmt.Println("--------------------------------------------------------")
	a.valueMethodSet()
	b.valueMethodSet()
	fmt.Println("--------------------------------------------------------")
	fmt.Println()

	fmt.Println("--------------------------------------------------------")
	fmt.Println("Pointer Method Set")
	fmt.Println("--------------------------------------------------------")
	a.pointerMethodSet()
	b.pointerMethodSet()
	fmt.Println("--------------------------------------------------------")
}
