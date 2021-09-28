package main

import (
	"log"
	"time"
)

func printLog() {
	const t1 = time.Second * 5
	const t2 = time.Second * 10
	const t3 = time.Second * 15
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(t1)
			ch1 <- "5 seconds"
			log.Println("Log after every 5 sec")
		}
	}()

	go func() {
		for {
			time.Sleep(t2)
			ch2 <- "10 seconds"
			log.Println("Log after every 10 sec")
		}
	}()

	go func() {
		for {
			time.Sleep(t3)
			ch3 <- "15 seconds"
			log.Println("Log after every 15 sec")
		}
	}()

	for {
		select {
		case <-ch1:
		case <-ch2:
		case <-ch3:
		}
	}
}
func main() {
	printLog()
}
