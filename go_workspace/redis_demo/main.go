package main

import (
	"fmt"

	"context"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`

	Age int `json:"age"`
}

func main() {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{

		Addr: "localhost:6379", // use default Addr

		Password: "", // no password set

		DB: 1, // use default DB

	})

	pong, err := client.Ping(ctx).Result()

	fmt.Println(pong, err)

	jaonData, err := json.Marshal(Author{Name: "Elliot", Age: 25})

	if err != nil {

		fmt.Println(err)

	}

	err = client.Set(ctx, "id1234", jaonData, 5000000000).Err() // key, value, expiration

	if err != nil {

		fmt.Println(err)

	}

	val, err := client.Get(ctx, "id1234").Result()

	if err != nil {

		fmt.Println(err)

	}

	fmt.Println(val)

}
