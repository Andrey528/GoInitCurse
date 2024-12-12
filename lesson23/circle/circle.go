package circle

import "fmt"

func init() {
	fmt.Println("Init circle")
}

type Circle struct {
	Radius float64
}

func New(newRadius float64) *Circle {
	return &Circle{newRadius}
}
