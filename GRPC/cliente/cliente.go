package main

import (
	"context"
	"fmt"
	"log"
	"time"

	gamepb "squidGame/squid.game"

	"google.golang.org/grpc"
)

func insertGame(id int64, juego string, max int64){
	server_host := "0.0.0.0:8082"

	fmt.Println("Enviando peticion . . .")

	conn,err:= grpc.Dial(server_host, grpc.WithInsecure(),grpc.WithBlock())

	if err != nil{
		fmt.Println("Error enviando peticion :  %v",err)

	}

	defer conn.Close()

	c:= gamepb.NewGameServiceClient(conn)
	
	fmt.Println("Todo bien hasta aqui")

	request := &gamepb.GameRequest{
		Game: &gamepb.Juego{
			Id:    id,
			Juego: juego,
			Max:   max,
			
		},
	}
	
	
	fmt.Println("Enviando datos al servidor")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,err:= c.RegGame(ctx, request)
	
	if err != nil{
		log.Fatal("Error al enviar petición: %v",err)
	}

	fmt.Println("okas: ", res.Result)


}

func main(){
	insertGame(2,"A",20)
}