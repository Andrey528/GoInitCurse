package main

import (
	"fmt"
	"lesson22/rectangle"
)

func main() {
	resAdd := Add(10, 20)
	resSub := Sub(30, 40)
	resMullt := Mult(50, 4)

	fmt.Println(resAdd, resSub, resMullt)

	r := rectangle.New(10, 20, "Green")
	fmt.Println("perimeter:", r.Perimeter())

	newR := rectangle.Rectangle{
		A: 10,
		B: 7,
	}
	fmt.Println(newR)
}
