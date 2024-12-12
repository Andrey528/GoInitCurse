package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name %s and age %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer
	stud1 := Student{"Alex", 23}
	descr1 = stud1
	descr1.Describe()

	stud2 := &Student{"Bob", 21}
	descr1 = stud2
	descr1.Describe()

	var descr2 Describer
	prof1 := &Professor{"victor Wann", 72}
	descr2 = prof1
	descr2.Describe()

	prof2 := Professor{"Bob Brown", 65}
	prof2.Describe()

}
