package main

import "fmt"

func CheckDBCloseConnection(a int) {
	fmt.Println("Check started... Value:", a)
	fmt.Println("Check done")
}

func OuterFunnc() int {
	defer fmt.Println("i'm deferred print function")
	fmt.Println("OuterFunction sturted")
	var result = 10
	fmt.Println("Outer function finished")
	return result
}

func main() {
	defer fmt.Println("Step 1")
	defer fmt.Println("Step 2")
	defer fmt.Println("Step 3")

	var numIp = 10
	p := &numIp
	defer CheckDBCloseConnection(numIp)
	*p++
	fmt.Println("Value numIp in main():", numIp)
	fmt.Println("Main function started")
	fmt.Println("Main function ended")
	fmt.Println("Value from Outer function:", OuterFunnc())
}
