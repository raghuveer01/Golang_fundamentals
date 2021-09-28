package main

import (
	"log"
	"time"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {

	const t1 = time.Second * 5
	const t2 = time.Second * 4
	const t3 = time.Second * 12

	i := 0
	for {
		i += 1
		time.Sleep(time.Second)
		if i%5 == 0 {
			// time.Sleep(t1)
			log.Println(t1)
		}
		if i%4 == 0 {
			log.Println(t2)
		}
		if i%12 == 0 {
			log.Println(t3)
		}
		if i == 12 {
			i = 0
		}
	}
}
