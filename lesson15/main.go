package main

import "fmt"

func main() {
	var variable int = 30
	var pointer *int = &variable
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	var zeroPinter *int
	fmt.Printf("Type %T and value %v\n", zeroPinter, zeroPinter)
	if zeroPinter == nil {
		zeroPinter = &variable
		fmt.Printf("After initialization type %T and value %v\n", zeroPinter, zeroPinter)
	}

	var numericValue int = 32
	var pointerToNumeric *int = &numericValue
	fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)

	zeroPoint := new(int)
	fmt.Printf("Value in *zeroPoint %v\nAddress is %v\n", *zeroPoint, zeroPoint)

	zeroPointerToInt := new(int)
	fmt.Printf("Address is %v and value in zeroPointerToInt is: %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Address is %v and value in zeroPointerToInt is: %v\n", zeroPointerToInt, *zeroPointerToInt)

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println(b)

	sample := 1
	samplePointer := &sample
	fmt.Println("Origin value of sample:", sample)
	chaneParam(samplePointer)
	fmt.Println("After changing sample is:", sample)

	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)
}

func chaneParam(val *int) {
	*val += 100
}

func returnPointer() *int {
	var numeric int = 321
	return &numeric
}
