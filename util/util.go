package main

import (
	"fmt"
	//"encoding/json"
	//"github.com/go-redis/redis/v8"
	
	//"github.com/gorilla/mux"
	"context"
    
    //"log"
	"os"
    //"io/ioutil"
    //"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	
)

type Juego struct {
	Id int64 `json: "ID"`
	Juego string `json: "juego"`
	Ganador int64 `json: "max"`
}

type JuegoMongo struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Identificador int `json: "ID"`
	Juego string `json: "juego"`
	Ganador int `json: "max"`
}


func JuegoMongo1(){
	clintOptions := options.Client().ApplyURI("mongodb://34.125.189.71:27017")
	fmt.Println("ClientOptions Type: ", reflect.TypeOf(clintOptions), "\n")
	client, err := mongo.Connect(context.TODO(), clintOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("squid-game").Collection("games")

	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := JuegoMongo{
		Identificador:  525,
		Juego:  "prueba",
		Ganador: 45,
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}
}


func main() {
	/*
	client := redis.NewClient(&redis.Options {
		Addr: "34.125.230.217:6379",
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
	*/
	JuegoMongo1()
}