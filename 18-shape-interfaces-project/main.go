package main

import (
	"fmt"
	"net/http"
)

var resp http.Response

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	triangle := triangle{base: 10, height: 10}
	square := square{sideLength: 10}
	printArea(triangle)
	printArea(square)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
