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
	Id      int    `json:"ID"`
	juego   string `json:"juego"`
	max int    `json:"max"`
	players int    `json:"players"`
	worker  string `json:"worker"`
}

type server struct {
	gamepb.UnimplementedGameServiceServer
}

func juego1(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int

	for i := 1; i <= jugadores; i++ {
		ganador = append(ganador, i)

	}

	ganador1 := rand.Intn(len(ganador))
	Campeon := ganador[ganador1]
	return Campeon
}

/* juego 2 */
func juego2(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int

	for i := 1; i <= jugadores; i++ {

		if i%2 == 1 {
			ganador = append(ganador, i)

		}

	}

	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

	return campeon2

}

/* juego 3 */
func juego3(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	for i := 1; i <= jugadores; i++ {

		if i%2 == 0 {
			ganador = append(ganador, i)

		}

	}

	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

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

    //    identificador := strconv.FormatInt(id, 10)
    //ganador := strconv.Itoa(winner)

    byt := []byte(`{"ID":0}`)

    var dat map[string]interface{}
    err := json.Unmarshal(byt, &dat); 

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(dat)

    dat["ID"] = desicion
    dat["juego"] = juego
    dat["max"] = winner
    dat["players"] = maximo
    dat["worker"] = "rabbitMq"

    data, err := json.Marshal(dat)

  /*  peticion, _ := json.Marshal(gameStruct{
        Id:      desicion,
        juego:   juego,
        max: winner,
        players: maximo,
        worker:  "kafka",

    })
*/
    petition := bytes.NewBuffer(data)

    response, err := http.Post("http://104.197.111.247:2062/", "application/json", petition)
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
	host := "0.0.0.0:50052"
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
