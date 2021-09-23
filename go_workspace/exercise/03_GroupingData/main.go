package main

import (
	"fmt"
)

func main() {

	//Array
	fmt.Println("Array")
	var x [5]int
	fmt.Println(x)
	x[3] = -1
	fmt.Println(x)
	fmt.Println(len(x))

	//Slice
	fmt.Println("\nSlice")
	y := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(y)
	for i, v := range y {
		fmt.Println(i, v)
	}

	//slice using make([]T, length, capacity)
	make := make([]int, 10, 12)
	fmt.Println(make)

	//slicing a slice
	fmt.Println("\nSlicing a slice y")
	fmt.Println(y[1:4])

	//append to slice
	fmt.Println("\nAppend to slice y")
	y = append(y, 77, 88, 99)
	fmt.Println(y)

	z := []int{234, 456, 678, 987}
	y = append(y, z...)
	fmt.Println(y)

	//delete a slice
	fmt.Println("\nDeleting a slice")
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

}
