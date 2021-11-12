package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "mytopic"
	brokerAddress = "34.125.131.17:19092"
)

func produce(ctx context.Context, resultado string) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte(resultado),
	})
	if err != nil {
		resultado = "error al escribir"
	}

}

func manejador(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", string(body))

	value := string(body)
	//	mesage, _ := sjson.Set(value, "way", "Kafka")

	//	resultado := mesage
	go produce(context.Background(), value)
	w.Write([]byte(value))
}

func main() {
	http.HandleFunc("/", manejador)
	fmt.Println("El servidor se encuentra en ejecución")
	fmt.Println(http.ListenAndServe(":2062", nil))
}
