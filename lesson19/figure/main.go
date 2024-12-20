package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

func PerimeterCircle(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", PerimeterCircle(c))
	fmt.Println("Call nethod:", c.Perimeter())

	r := Rectangle{10, 20}
	fmt.Println("Call function:", PerimeterRectangle(r))
	fmt.Println("Call nethod:", r.Perimeter())
}
