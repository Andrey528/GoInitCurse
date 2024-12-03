package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello world"
	fmt.Println(name)

	word := "Sample word"
	fmt.Printf("String @s\n", word)

	fmt.Printf("Bytes:")

	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i])
	}
	fmt.Println()

	fmt.Printf("Characters:")

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
	fmt.Println()

	runeSlice := []rune(word)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()

	for id, runeVal := range word {
		fmt.Printf("%c starts at position %d\n", runeVal, id)
	}

	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103}
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrfromRune := string(runeHexSlice)
	fmt.Println("form runes(hex):", myStrfromRune)

	runeLiteralSlice := []rune{'v', 'a', 's', 'y', 'a'}
	myStrFromRuneLiterals := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFromRuneLiterals)

	fmt.Printf("%s and %T\n", myStrFromRuneLiterals, myStrFromRuneLiterals)

	fmt.Println("Length of Вася:", len("Вася"), "bytes")

	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	fmt.Println(cap([]rune("Вася")))

	word1, word2 := "Вася", " Петя"
	if word1 != word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	word3 := word1 + word2
	fmt.Println(word3)

	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println("FullName:", fullName)

	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	fullName = string(fullNameSlice)

	if 'Q' == 'Q' {
		fmt.Println("Runes equal")
	} else {
		fmt.Println("runes not equal")
	}
}
