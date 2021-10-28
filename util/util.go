package main

import (
	"fmt"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"context"
)

type Juego struct {
	Id int64 `json: "ID"`
	Juego string `json: "juego"`
	Ganador int64 `json: "max"`
}

func main() {
	client := redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	var ctx = context.Background()

	json, err := json.Marshal(Juego{Id: 2, Juego: "Juego 1", Ganador: 20})

	if err != nil {
		fmt.Println(err)
	}

	

	err1 := client.Set(ctx, "key", json, 0).Err()
    if err1 != nil {
        panic(err1)
    }

	val, err := client.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }

	fmt.Println(val)
}