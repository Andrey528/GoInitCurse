package main

import "fmt"

type Rectangle struct {
	length, wwidth int
}

func (r *Rectangle) Area() int {
	return r.length * r.wwidth
}

func Area(r *Rectangle) int {
	return r.length * r.wwidth
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.wwidth = newWidth
}

func main() {
	rectangleAsValue := Rectangle{10, 10}
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call Area function for &rectangle:", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &rectangle:", rectangleAsPointer.Area())

	fmt.Println("Call Area method for rectangle", rectangleAsValue.Area())

	fmt.Println("Before changing width:", rectangleAsValue)

	rectangleAsPointer.SetWidth(999)
	fmt.Println("After change via method SetWidth for &rectangle:", rectangleAsValue)

	rectangleAsValue.SetWidth(888)
	fmt.Println("After change via method SetWidth for rectangle:", rectangleAsValue)
}
