package rectangle

import "fmt"

func init() {
	fmt.Println("Init function from rectangle package")
	fmt.Println("NName:", name, "Age:", age)
}

var (
	name string = "John"
	age  int    = 33
)

type Rectangle struct {
	A, B int
}

func New(newA, newB int) *Rectangle {
	return &Rectangle{newA, newB}
}
