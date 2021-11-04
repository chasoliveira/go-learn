package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type tringle struct {
	height float64
	base   float64
}

func main() {
	sq := square{
		sideLength: 2,
	}
	tg := tringle{
		base:   3,
		height: 4,
	}

	printArea(sq)
	printArea(tg)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t tringle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
