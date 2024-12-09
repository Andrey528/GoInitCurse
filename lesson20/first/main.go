package main

import (
	"fmt"
	"unicode/utf8"
)

type BigWord interface {
	IsBig() bool
}

type MySuperString string

func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("dkksdkkgksnfdkgsg")
	var interfaceSample BigWord
	interfaceSample = sample
	fmt.Println("Is big?", interfaceSample.IsBig())
}
