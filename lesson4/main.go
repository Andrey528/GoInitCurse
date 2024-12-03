package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var firrstBoolean bool
	fmt.Println(firrstBoolean)

	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	var a int = 32
	fmt.Printf("Type us %T\n", a)
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32 + int32(second64))

	var third64 int64 = 16123414
	var fourthint int = 156234
	fmt.Println(third64 + int64(fourthint))

	floatFirsst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirsst, floatSecond)
	sum := floatFirsst + floatSecond
	sub := floatFirsst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum: %.3f and Syb: %.3f\n", sum, sub)

	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", name, len(name))
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))

	totalString, subString := "ABCDEDFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))

	var sampleRune rune
	var anotherRune rune = 'Q'
	var thirdRune rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as cahr %c\n", sampleRune)
	fmt.Printf("Rune as cahr %c\n", anotherRune)
	fmt.Printf("Rune as cahr %c\n", thirdRune)
	fmt.Println(strings.Compare("abcd", "a"))

	var aByte byte
	fmt.Println("Byte:", aByte)
}
