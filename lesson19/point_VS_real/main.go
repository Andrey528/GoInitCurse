package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call function for rectangleAsValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter())

	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter())

	fmt.Println("Before call method SetLength:", rectangleAsValue)
	rectangleAsValue.SetLength(9999)
	fmt.Println("After call methhod SetLegth:", rectangleAsValue)

	rectangleAsPointer.SetLength(99999999999)
	fmt.Println("After call method SetLength for &rectangle:", *rectangleAsPointer)
}
