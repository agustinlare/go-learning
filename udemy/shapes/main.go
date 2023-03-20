package main

import "fmt"

type triangle struct {
	hight float64
	base  float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 3.15}
	printArea(s)
	t := triangle{hight: 3.15, base: 5.32}
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	r := s.sideLength * s.sideLength
	return r
}

func (t triangle) getArea() float64 {
	r := 0.5 * t.hight * t.base
	return r
}
