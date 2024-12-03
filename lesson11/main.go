package main

import "fmt"

func main() {
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs:", mapper)

	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}

	newMapper["Jo"] = 3000
	fmt.Println("New mapper:", newMapper)

	testPerson := "Alice"
	fmt.Println("Salary of testPerson", newMapper[testPerson])

	testPerson = "Derek"
	fmt.Println("Salary of new testPersoon:", newMapper[testPerson])

	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}

	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den does not exists in map")
	}

	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo and value:", value)
	} else {
		fmt.Println("Jo does not exists in map")
	}

	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}

	fmt.Println("Before deleting")
	delete(employee, "Den")
	fmt.Println("After first deleting")

	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna")
	}
	fmt.Println("After second deleting", employee)

	fmt.Println("Pair amount in map:", len(employee))

	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words map:", words)
	fmt.Println("newWords map:", newWords)

	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not equal")
	}

	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero val of map bMap")
	}
}
