package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and %d y.o\n", std.name, std.age)
}

func Typefinder(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer:
		fmt.Println("I'm describer")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 23}
	Typefinder(10)
	Typefinder("Hello")
	Typefinder(nil)
	Typefinder(std)
}
