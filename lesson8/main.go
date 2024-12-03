package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("This is my araay:", arr)

	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720

	fmt.Println("After elements init:", arr)

	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declaration and init:", newArr)

	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))

	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First arr:", first)
	fmt.Println("Second arr:", second)

	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}

	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element on arr is %.2f\n", id, val)
		sum += val
	}
	fmt.Println("Total sum is:", sum)

	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}

	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimensional array:", words)

	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s", val2)
		}
		fmt.Println()
	}

	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200)
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	emptySlice = append(emptySlice, 200)
}
