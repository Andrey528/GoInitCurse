package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Second line")
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)

	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after asignment", age)

	var height int = 183
	fmt.Println("My height is:", height)

	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)

	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)
	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUID)

	var a, b = 30, "Vova"
	fmt.Println(a, b)

	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)

	count, cArg := 300, 400
	fmt.Println(count, cArg)

	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.3f\n", math.Min(width, length))
}
