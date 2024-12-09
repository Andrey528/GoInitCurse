package main

import "fmt"

func mutation(arr *[3]int) {
	(*arr)[1] = 909
	(*arr)[2] = 10000
}

func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation:", arr)

	newArr := [3]int{1, 2, 4}
	fmt.Println("Arr before mutation:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("Arr after mutation:", newArr)
}
