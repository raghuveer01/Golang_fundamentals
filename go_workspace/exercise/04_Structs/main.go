package main

import (
	"fmt"
)

//declaring struct
type person struct {
	fname  string
	lname  string
	gender string
}

func main() {
	//defining struct
	person1 := person{
		fname:  "A",
		lname:  "B",
		gender: "M",
	}

	person2 := person{
		fname:  "C",
		lname:  "D",
		gender: "F",
	}

	//Print struct person
	fmt.Println("Struct Person")
	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.fname, person1.lname, person1.gender)
	fmt.Println(person2.fname, person2.lname, person2.gender)

	//Anonymous Struct
	fmt.Println("\nAnonymous Struct")
	anon := struct {
		name    string
		friends []string
	}{
		name: "Rupesh",
		friends: []string{
			"Prem",
			"Ajit",
		},
	}

	fmt.Println(anon)
}
