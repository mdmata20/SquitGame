package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"io/ioutil"

	"github.com/gorilla/mux"

	"context"

	"github.com/go-redis/redis/v8"

	//"log"
	"os"
	//"io/ioutil"
	//"net/http"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome API")
}

type JuegoMongo struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Identificador int                `json: "ID"`
	Juego         string             `json: "juego"`
	Ganador       int                `json: "max"`
}

type Juego struct {
	Id      int    `json: "ID"`
	Juego   string `json: "juego"`
	Ganador int    `json: "ganador"`
}

func JuegoRedis(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "34.125.131.17:6379",
		Password: "",
		DB:       0,
	})
	var juego Juego
	_ = json.NewDecoder(r.Body).Decode(&juego)
	var ctx = context.Background()

	//	json, err := json.Marshal(Juego{Id: 2, Juego: "Juego 1", Ganador: 20})
	json, err := json.Marshal(juego)
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

func JuegoMongo1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var juego Juego
	_ = json.NewDecoder(r.Body).Decode(&juego)

	clintOptions := options.Client().ApplyURI("mongodb://34.125.189.71:27017")
	client, err := mongo.Connect(context.TODO(), clintOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	col := client.Database("squid-game").Collection("games")

	fmt.Println("ClientOptions Type: ", reflect.TypeOf(clintOptions), "\n")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	result, insertErr := col.InsertOne(ctx, juego)
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

	json.NewEncoder(w).Encode(result)

	//fmt.Fprintf(w, "Se Funcion")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/JuegoMongo1", JuegoMongo1)
	router.HandleFunc("/JuegoRedis", JuegoRedis)

	log.Fatal(http.ListenAndServe(":3010", router))

}
