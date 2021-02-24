package main

import "fmt"

type shape interface {
	hasArea
	hasPerimeter
}

type hasArea interface {
	Area() float64
}

type hasPerimeter interface {
	Perimeter() float64
}

func printShape(shape shape) {
	printArea(shape)
	printPerimeter(shape)
}

func printArea(shape hasArea) {
	fmt.Println("Area", shape.Area(), "(units squared)")
}

func printPerimeter(shape hasPerimeter) {
	fmt.Println("Perimeter", shape.Perimeter(), "(units)")
}

func runShapesDemo() {
	a := square{5}
	b := circle{5}

	/* The method sets in square, variable a, are using pointers */
	printShape(&a)
	printShape(b)
}
