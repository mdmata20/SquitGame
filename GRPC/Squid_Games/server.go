package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"log"
	"net"

	gamepb "squidGame/squid.game"

	"google.golang.org/grpc"
)

type gameStruct struct {
	Id      int    `json:"Id"`
	Juego   string `json:"Juego"`
	Ganador int    `json:"Ganador"`
}

type server struct {
	gamepb.UnimplementedGameServiceServer
}

func juego1(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int

	fmt.Println("--------- juego 1 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {
		ganador = append(ganador, i)
		fmt.Println(ganador)
	}

	fmt.Println("El ganador es: ")

	ganador1 := rand.Intn(len(ganador))
	Campeon := ganador[ganador1]

	fmt.Println(Campeon)
	return Campeon
}

/* juego 2 */
func juego2(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	fmt.Println("--------- juego 2 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {

		if i%2 == 1 {
			ganador = append(ganador, i)

		}

	}

	fmt.Println("El ganador es: ")
	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

	fmt.Println(campeon2)

	return campeon2

}

/* juego 3 */
func juego3(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	fmt.Println("--------- juego 2 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {

		if i%2 == 0 {
			ganador = append(ganador, i)

		}

	}

	fmt.Println("El ganador es: ")
	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

	fmt.Println(campeon2)

	return campeon2

}

func (s *server) RegGame(ctx context.Context, in *gamepb.GameRequest) (*gamepb.GameResponse, error) {

	var winner int
	result := ""

	id := in.GetGame().GetId()
	juego := in.GetGame().GetJuego()
	max := in.GetGame().GetMax()

	desicion := int(id)
	maximo := int(max)

	if desicion == 1 {
		winner = juego1(maximo)
	} else if desicion == 2 {
		winner = juego2(maximo)
	} else if desicion == 3 {
		winner = juego3(maximo)
	} else {
		winner = 0
	}

	//	identificador := strconv.FormatInt(id, 10)
	//ganador := strconv.Itoa(winner)

	peticion, _ := json.Marshal(gameStruct{
		Id:      desicion,
		Juego:   juego,
		Ganador: winner,
	})

	petition := bytes.NewBuffer(peticion)

	response, err := http.Post("http://localhost:3010/JuegoMongo1", "application/json", petition)
	if err != nil {
		log.Fatalln("Error sending info", err)
	}
	//response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// SE NECESITA OTRO ENDPOINT
	result += string(body)
	fmt.Println(result + "------------")
	//result := " ID: " + identificador + " Juego: " + juego + " Ganador!: " + ganador

	res := &gamepb.GameResponse{
		Result: result,
	}

	return res, nil
}
func main() {
	host := "0.0.0.0:8082"
	fmt.Println("Server iniciado en: " + host)

	lis, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("El server no funciono %v", err)
	}

	fmt.Println("Starting gRPC server ...")

	s := grpc.NewServer()
	fmt.Println("gRPC server is up!")
	gamepb.RegisterGameServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("el server no funcionó: %v", err)
	}

}
