package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main(){
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil{
		log.Fatal("dialing ....")
	}
	defer conn.Close()

	client := go_protoc.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &go_protoc.String{Value: "hello-sadly"})
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(reply)
}
