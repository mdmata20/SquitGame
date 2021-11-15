package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"io/ioutil"
	//"strconv"
	"github.com/gorilla/mux"

	"context"

	"github.com/go-redis/redis/v8"

	//"log"
	"os"
	//"io/ioutil"
	//"net/http"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome API")
}

type JuegoMongo struct {
	Identficador            primitive.ObjectID `bson:"_id,omitempty"`
	ID int                `json: "ID"`
	juego         string             `json: "juego"`
	max       int                    `json: "max"`
	players       int                `json: "players"`
	worker        string             `json: "worker"`
}

type Juego struct {
	Id    int    `json: "ID"`
	Juego string `json: "juego"`
	Max   int    `json: "max"`
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

	keytime := time.Now()

	err1 := client.Set(ctx, keytime.String(), json, 0).Err()
	if err1 != nil {
		panic(err1)
	}

	val, err := client.Get(ctx, keytime.String()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
/*
func JuegoMongo1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")


	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(body)
	nuevo := string(data)
	fmt.Println(nuevo)

	var juego JuegoMongo
	
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

	result, insertErr := col.InsertOne(ctx, data)
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

*/

func JuegoMongo1(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")

    var body map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        fmt.Println(err)
    }
    data, err := json.Marshal(body)

    //observar la info
    nuevo := string(data)
    fmt.Println(nuevo)

    var juego JuegoMongo
	
	_ = json.NewDecoder(r.Body).Decode(&juego)

    clintOptions := options.Client().ApplyURI("mongodb://34.125.189.71:27017")
    client, err := mongo.Connect(context.TODO(), clintOptions)
    if err != nil {
        fmt.Println("Mongo.connect() ERROR: ", err)
        os.Exit(1)
    }
    col := client.Database("squid-game").Collection("games")

    fmt.Println("ClientOptions Type: ", reflect.TypeOf(clintOptions), "\n")

    ctx, _ := context.WithTimeout(context.Background(), 15 * time.Second)

    fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

    result, insertErr := col.InsertOne(ctx, bson.D{
        {Key: "ID", Value: body["ID"]},
        {Key: "juego", Value: body["juego"]},
        {Key: "max", Value: body["max"]},
        {Key: "players", Value: body["players"]},
        {Key: "worker", Value: body["worker"]},
    })

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
