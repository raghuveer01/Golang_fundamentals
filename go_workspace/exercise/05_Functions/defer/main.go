package main

import (
	"fmt"
)

func main() {
	defer one()
	other()
	defer another()
	brother()
}

func one() {
	fmt.Println("one")
}

func other() {
	fmt.Println("other")
}

func another() {
	fmt.Println("another")
}

func brother() {
	fmt.Println("brother")
}
