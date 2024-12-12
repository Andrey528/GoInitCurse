package main

import (
	"fmt"
	"lesson23/circle"
	_ "lesson23/circle"
	"lesson23/rectangle"
)

func init() {
	fmt.Println("Init function for main package")
	fmt.Println("Name:", name, "Age:", age)
}

var (
	name string = "bob"
	age  int    = 99
)

func main() {
	r := rectangle.New(10, 10)
	fmt.Println(r)
	c := circle.New(22.5)
	fmt.Println(c)
}
