package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}

	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Printf("After for loop with break")

	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Printf("After for loop with continue")

	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идее выше треугольник")

outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}

	var loopVar int = 0
	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}

	var password string
outer2:
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
		} else {
			fmt.Println("Password Accepted")
			break outer2
		}
	}

	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}
}
