package main

import "fmt"

func add(a int, b int) int {
	result := a + b
	return result
}

func main() {
	fmt.Println("Hello world")

	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)

	fmt.Println("Result of mult(1, 2, 3, 4):", mult(1, 2, 3, 4))

	per, area := rectangleParams(10.5, 10.5)
	newPer, _ := rectangleParams(10, 10)
	fmt.Println("Area of rect(10.5, 10.5):", area)
	fmt.Println("Perimeter of rect(10.5, 10.5):", per)
	fmt.Println("Perimeter of rect(10, 10):", newPer)

	secondPer, secondArea := namedreturn(10, 20)
	fmt.Println(secondArea, secondPer)

	helloVariadic(10, 20, 30)
	helloVariadic()

	someStrings(10, 20, "Bob", "Alex")
	someStrings(10, 20)

	sum1 := sumVariadic(10, 20, 30)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)

	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon:", anon(20, 30))

	fmt.Println("BigFunction(10, 20):", bigFunction(10, 20))
}

func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b }(aArg, bArg)
}

func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

func rectangleParams(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width

	return perimeter, area
}

func namedreturn(a, b int) (perimeter, area int) {
	perimeter = 2 * (a + b)
	area = a * b
	return
}

func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

func emptyReturn(a int) {
	fmt.Println("func with empty return. a:", a)
	return
}

func helloVariadic(a ...int) {
	fmt.Println(a)
	fmt.Printf("value %v and type %T\n", a, a)
}

func someStrings(a, b int, words ...string) {
	fmt.Println("Parametr", a)
	fmt.Println("Parametr", b)
	var result string
	for _, word := range words {
		result += word
	}

	fmt.Println("Result concat:", result)
}

func sumVariadic(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}
