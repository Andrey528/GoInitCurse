package main

import "fmt"

type Empty interface {
}

type Student struct {
	name string
}

func Describer(pretendent interface{}) {
	fmt.Printf("type = %T and value %v\n", pretendent, pretendent)
}

func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int)
	fmt.Println("Value:", val, "Ok?:", ok)
}

func main() {
	strSample := "Hello world"
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vova"})

	SemiGeneric(10)
	SemiGeneric(Student{"fedya"})
	SemiGeneric("Hello world")
}
