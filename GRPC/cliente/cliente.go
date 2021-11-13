package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	gamepb "squidGameC/squid.game" // debe tener el mismo cliente

	"google.golang.org/grpc"
)

type juegoStruct struct {
	Id    int64
	Juego string
	Max   int64
}

func insertGame(id int64, juego string, max int64) {
	server_host := "Cserver:50051" //debe ser la de cliente

	fmt.Println("Sending petition . . .")

	conn, err := grpc.Dial(server_host, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		fmt.Println("Error enviando peticion :  %v", err)

	}

	defer conn.Close()

	fmt.Println("aquiii vaaa")

	c := gamepb.NewGameServiceClient(conn)

	request := &gamepb.GameRequest{
		Game: &gamepb.Juego{
			Id:    id,
			Juego: juego,
			Max:   max,
		},
	}
	fmt.Println(request)
	fmt.Println("sending data to server...")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.RegGame(ctx, request)

	if err != nil {
		log.Fatal("Error al enviar petición: %v", err)
	}

	fmt.Println("Working, data received: ", res.Result)

}

//               quien responde              lo que trae
func http_server(w http.ResponseWriter, r *http.Request) {
	//instance_Name:= "grpcInstancia"
	//fmt.Printf("Http request from client: ", instance_Name)
	fmt.Println("Welcome to api")

	instance_name := "grpcInstancia"

	fmt.Println(instance_name)

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	switch r.Method {
	case "GET":
		fmt.Println("Http for client")
		http.StatusText(http.StatusAccepted)
	case "POST":
		fmt.Println("Sending...")
		decoder := json.NewDecoder(r.Body)

		var squidgame juegoStruct

		err := decoder.Decode(&squidgame)

		if err != nil {
			fmt.Printf("El error está en la decodificacion %v", err)
		}

		fmt.Fprintf(w, "se recibió \n")

		insertGame(squidgame.Id, squidgame.Juego, squidgame.Max)

	default:
		fmt.Println("me perdí y me dio amsiedad")
		return
	}

}

func main() {

	instance_name := "grpcInstancia"

	fmt.Println(instance_name)

	host := ":8080"



	fmt.Println("Server started at: ", host)

	http.HandleFunc("/", http_server)

	if err := http.ListenAndServe(host, nil); err != nil {
		log.Fatal(err)
	}

}
