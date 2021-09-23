package main

import (
	"fmt"
)

var y string
var i int
var f float32
var b bool

func main() {

	fmt.Println("\nZero Value:")
	fmt.Printf("%T\t %v\n", y, y)
	fmt.Printf("%T\t %v\n", i, i)
	fmt.Printf("%T\t %v\n", f, f)
	fmt.Printf("%T\t %v\n", b, b)

	//create your own type
	var aa int

	type ownType int

	var bb ownType
	fmt.Println("\nNewType")
	fmt.Printf("%T\n", aa)
	fmt.Printf("%T\n", bb)

	//var x int
	x := 42
	fmt.Println(x)

	//iota
	const (
		c0 = iota
		c1
		c2
		a     = 43
		b int = 55
	)
	const c3 = iota
	fmt.Println("\nIota")
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(a)
	fmt.Println(b)
}
