package main

type square struct {
	side uint
}

func (a *square) Area() float64 {
	return float64(a.side * a.side)
}

func (a *square) Perimeter() float64 {
	return float64(a.side * 4)
}
