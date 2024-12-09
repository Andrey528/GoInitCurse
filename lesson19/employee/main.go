package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary :%d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "Bob",
		position: "Senior Go developer",
		salary:   3000,
		currency: "USD",
	}
	emp.DisplayInfo()

}
