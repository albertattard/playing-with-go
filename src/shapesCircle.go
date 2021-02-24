package main

import "math"

type circle struct {
	radius uint
}

func (c circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func (c circle) Perimeter() float64 {
	return math.Pi * 2 * float64(c.radius)
}
