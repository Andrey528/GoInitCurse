package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e Employee) SetName(newName string) {
	e.name = newName
}

func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting parameters:", e)
	e.SetName("Bob")
	fmt.Println("After SetName call:", e)
	(&e).SetSalary(4500)
	fmt.Println("After SetSalary call:", e)
}
