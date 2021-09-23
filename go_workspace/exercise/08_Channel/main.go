package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 2
	//fmt.Println(<-c)
	//go func() {
	c <- 111
	//c <- 3
	//}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
}
