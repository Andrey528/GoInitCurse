package main

import "fmt"

func main() {
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2]
	fmt.Println("Slice[0:2]:", startSlice)

	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4]
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

	wordSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordSlice, "Length:", len(wordSlice), "Capacity:", cap(wordSlice))
	wordSlice = append(wordSlice, "four")
	fmt.Println("Slice:", wordSlice, "Length:", len(wordSlice), "Capacit:", cap(wordSlice))

	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // перестает ссылаться на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	sl := make([]int, 10, 15) // тип, длина, емкость
	fmt.Println(sl)

	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	fmt.Println("myWords:", myWords)
}
