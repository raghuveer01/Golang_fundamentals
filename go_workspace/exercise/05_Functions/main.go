package main

import (
	"fmt"
)

func main() {
	//x1:=[]int{1,2,3,4,5,6,7}
	//x:=sum(x1...)
	x := add(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("The total is", x)

	//anonymous function
	func() {
		fmt.Println("\nAnonymous function\n")
	}()

	f := func(x int) {
		fmt.Println("Return:", x)
	}
	f(000)
}

func add(x ...int) int {

	add := 0
	for _, v := range x {
		add += v
	}
	fmt.Println("total = ", add)
	return add
}
