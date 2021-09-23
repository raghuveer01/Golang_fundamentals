package main

import (
	"fmt"
)

func main() {
	//if-else
	fmt.Println("If-else")
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if 1 == 1 {
		fmt.Println("1")
	}
	x := 2
	if x == 0 {
		fmt.Println("value is 0")
	} else {
		fmt.Println("value not 0")
	}

	fmt.Println("\nInitialization statement")
	if x := 4; x == 4 {
		fmt.Println("00")
	}

	// for init; condition; post {}
	fmt.Println("\nFor loop implementation")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)

	}

	//infinite loop
	fmt.Println("\nInfinite loop and break using loop")
	x = 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}

	// while loop
	fmt.Println("\nFor used as if While loop")
	x = 6
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//nested loop
	fmt.Println("\nNested For loop")
	for i := 0; i <= 2; i++ {
		fmt.Printf(" %d\n", i)
		for j := 5; j < 8; j++ {
			fmt.Printf(" %d\n", j)

		}
	}

	fmt.Println("\nSwitch case")
	switch {
	case false:
		fmt.Println("false")
	case (2 == 4):
		fmt.Println("2 equals 4")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, fallthrough")
		fallthrough
	default:
		fmt.Println("default")
	}
}
