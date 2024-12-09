package main

import "fmt"

func DoSomething(pretendent interface{}) {
	switch pretendent.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("this is int")
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hello world")
	DoSomething(func(a, b int) int { return a + b })
}
