package main

import (
	"context"
	"fmt"
	"grpccliente/greetpb"
	"log"

	"google.golang.org/grpc"
)

func enviarMensajes(firstName string, message string) {
	server_host := "0.0.0.0:50051"

	fmt.Println("Enviado peticion")

	cc, err := grpc.Dial(server_host, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Println("Conexion ok")

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: firstName,
			Message:   message,
		},
	}

	res, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatal("Error: %v", err)
	}

	fmt.Println("Todo bien, ", res.Result)

}

func main() {
	fmt.Println("LISTO")

	enviarMensajes("YO", "mimido")
}

// go run client.go
// go get github.com/golang/protobuf/proto
// go get google.golang.org/grpc
// go get google.golang.org/protobuf/reflect/protoreflect@1.25.0
