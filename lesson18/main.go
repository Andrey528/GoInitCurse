package main

import "fmt"

type University struct {
	yearBased int
	infoShort string
	innfoLong string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	University
}

func main() {
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "cool University",
			innfoLong: "very cool university",
		},
	}

	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Year based Uni:", stud.university.yearBased)

	prof := Professor{
		firstName: "Anatoly",
		lastName:  "Smirnov",
		age:       125,
		greatWork: "Ultimate C Programming",
		University: University{
			yearBased: 1734,
			infoShort: "Short info",
		},
	}
	fmt.Println("FirstName:", prof.firstName)
	fmt.Println("Year based:", prof.yearBased)

	profLeft := Professor{}
	profRight := Professor{}

	fmt.Println(profLeft == profRight)

}
