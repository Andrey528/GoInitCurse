package main

import "fmt"

func main() {
	var price int
	fmt.Scan(&price)

	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("third case")
	case 130:
		fmt.Println("Another case")
	default:
		fmt.Println("Default case")
	}

	var ageGroup string = "b"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are to yong/old")
	}

	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("too yong")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("too old")
	default:
		fmt.Println("Who are you")
	}

	var number int
	fmt.Scan(&number)
outer3:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 1--\n", number)
		if number%2 == 0 {
			break outer3
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d greater then 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d greater then 1000\n", number)
		fallthrough
	default:
		fmt.Printf("%d default\n", number)
	}

	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}

	fmt.Println("END")
}
