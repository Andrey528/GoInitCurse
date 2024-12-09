package main

import "fmt"

type MySuperDuperInt int

func (msdi *MySuperDuperInt) IsEven() bool {
	if *msdi%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MySuperDuperInt(202)
	num2 := MySuperDuperInt(201)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())
}
