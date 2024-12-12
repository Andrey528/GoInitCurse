package main

import "fmt"

type PerimeterCalculator interface {
	Perimeter() int
}

type AreaCalculator interface {
	Area() int
}

type FigureCalculator interface {
	PerimeterCalculator
	AreaCalculator
}

type Rectangle struct {
	a, b  int
	color string
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

func (r Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	var pCalc PerimeterCalculator
	var aCalc AreaCalculator
	rect := Rectangle{10, 20, "green"}
	pCalc = rect
	fmt.Println("Perimeter:", pCalc.Perimeter())

	aCalc = rect
	fmt.Println("Area:", aCalc.Area())

	var fCalc FigureCalculator
	fCalc = rect
	fmt.Println(fCalc.Area())
	fmt.Println(fCalc.Perimeter())
}
