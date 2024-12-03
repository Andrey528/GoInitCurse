package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		name = input.Text()
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	fmt.Println(numInt64, numFloat32)
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32))
}
