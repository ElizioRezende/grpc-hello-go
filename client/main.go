package main

import (
	"context"
	"github.com/eliziorezende/grpc-hello-go/pb"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main(){
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err!=nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	for x:=0 ; x<10; x++ {
		request := &pb.HelloRequest{
			Name: strconv.Itoa(x) + "-Elizio Rezende",
		}

		res, err := client.Hello(context.Background(), request)

		if err!=nil{
			log.Fatalf("Erro durante execucao %v", err)
		}

		log.Println(res.Msg)
	}
}

