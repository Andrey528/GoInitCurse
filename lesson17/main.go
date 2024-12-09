package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func PrintStudent(std Student) {
	fmt.Println("+++++++++++")
	fmt.Println("Firstname:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func main() {
	stud1 := Student{
		firstName: "Fedya",
		age:       21,
		lastName:  "Petrov",
	}
	PrintStudent(stud1)
	stud2 := Student{"Petya", "Ivanov", 19}
	PrintStudent(stud2)

	stud3 := Student{
		firstName: "Vasya",
	}
	PrintStudent(stud3)

	anonStudent := struct {
		age           int
		groupId       int
		proffesorName string
	}{
		age:           23,
		groupId:       2,
		proffesorName: "Alexeev",
	}
	fmt.Println("AnonStudent", anonStudent)

	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("FirstName:", studVova.firstName)
	studVova.age += 2
	fmt.Println("New age:", studVova.age)

	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer:", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value afterPointerModify:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	fmt.Println("Age via (*...).age:", (*secondPointer).age)
	fmt.Println("Age via .age:", secondPointer.age)

	human := &Human{
		firstName: "Bob",
		lastName:  "Johnson",
		string:    "Additional Info",
		int:       -1,
		bool:      true,
	}

	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}
