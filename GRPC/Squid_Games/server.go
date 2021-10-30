package main

import (
	"context"
	"fmt"
	"strconv"

	"log"
	"net"

	gamepb "squidGame/squid.game"

	"google.golang.org/grpc"
)


type server struct{gamepb.UnimplementedGameServiceServer}


func (s *server)RegGame(ctx context.Context, in *gamepb.GameRequest)(*gamepb.GameResponse, error){
	fmt.Println("todo ok")

	id := in.GetGame().GetId()
	juego:= in.GetGame().GetJuego()	
	winner := in.GetGame().GetMax()

	identificador:=strconv.FormatInt(id,10)
	ganador:= strconv.FormatInt(winner,10)

	result := " ID: " + identificador + " Juego: "+ juego + " Ganador!: "+ ganador

	res := &gamepb.GameResponse{
		Result:result,
	}

	return res,nil 
}


func main(){
	host:= "0.0.0.0:8082"
	fmt.Println("Server iniciado en: "+ host)

	lis, err := net.Listen("tcp",host)

	if err != nil{
		log.Fatalf("El server no funciono %v",err)
	}

	fmt.Println("Starting gRPC server ...")

	s:= grpc.NewServer()

	gamepb.RegisterGameServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil{
		log.Fatalf("el server no funcionó: %v",err)
	}
}