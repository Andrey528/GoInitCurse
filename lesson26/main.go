package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

func funcWithError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("this is error")
}

func panicExm(firstName *string, lastName *string) {
	defer PanicRecover()

	if firstName == nil {
		panic("runtime error: firstName can't be nil")
	}

	if lastName == nil {
		panic("runtime error: lastName can't be nil")
	}

	fmt.Println("Full name:", &firstName, &lastName)
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied:", r)
		debug.PrintStack()
	}
}

func main() {
	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		log.Fatal("Can not convert:", err)
	}
	fmt.Println("Convertion done:", num)

	ans, err := funcWithError(5)
	if err != nil {
		fmt.Println("Not use odd values", err)
		return
	}
	fmt.Println(ans)

	var name = "Bob"
	panicExm(&name, nil)
}
