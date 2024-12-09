package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func calcAndReturnValueFunc(command string, a int, b int) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

func recieveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

func addNums(a, b int) int {
	return a + b
}

func main() {
	var command string
	command = "addition"
	res := calcAndReturnValueFunc(command, 10, 20)
	fmt.Println("Result with command:", command, "value:", res(10, 20))

	fmt.Printf("type of func is %T\n", res)
	fmt.Printf("type of recieveFuncAndReturnValue is %T\n", recieveFuncAndReturnValue)

	fmt.Println("recieveFuncAndReturnValue(addNums):", recieveFuncAndReturnValue(addNums))
	fmt.Println(recieveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))
}
