package main

import "fmt"

type University struct {
	city string
	name string
}

func (u *University) fullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

type Professor struct {
	fullName string
	age      int
	University
}

func main() {
	p := Professor{
		fullName: "Boris Petrof",
		age:      120,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}

	p.fullInfoUniversity()
}
