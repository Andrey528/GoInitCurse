package main

import "fmt"

type FugureBuilder interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {
	r1 := Rectangle{10, 20}
	c1 := Circle{5}
	r2 := Rectangle{10, 20}
	c2 := Circle{5}
	r3 := Rectangle{10, 20}
	c3 := Circle{5}

	figures := []FugureBuilder{r1, r2, r3, c1, c2, c3}
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("total area:", total)

}
