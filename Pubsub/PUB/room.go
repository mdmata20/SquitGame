package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

)


func publishMessage(message []byte) {
	opt, err := redis.ParseURL("redis://34.123.108.198:6379/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)

	errs := redis.Publish(context.TODO(), "mensaje", message).Err()

	if errs != nil {
		log.Println(errs)
	}
}



func createTask(w http.ResponseWriter, r *http.Request) {

	requestAt := time.Now()
	w.Header().Set("Content-Type", "application/json")
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	log.Println("Error Parseando JSON: ", err)
	data, err := json.Marshal(body)
	log.Println("Error Reading Body: ", err)
	fmt.Println(string(data))
	publishMessage(data)
	duration := time.Since(requestAt)
	fmt.Fprintf(w, "Task scheduled in %+v", duration)
}

func main() {
	//CreateRedisClient()
	http.HandleFunc("/", createTask)
	fmt.Println("Server listening on port 3009...")
	if errors := http.ListenAndServe(":3009", nil); errors != nil {
		log.Fatal(errors)
	}
}
