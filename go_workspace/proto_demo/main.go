package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(elliot)
	fmt.Println(data)
	time.Sleep(time.Second * 3)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(newElliot)

}
