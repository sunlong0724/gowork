package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)
import "../pb"

type HelloServiceImpl struct{}

func (this *HelloServiceImpl)Hello(ctx context.Context, args *go_protoc.String)(*go_protoc.String, error){
	reply := &go_protoc.String{Value: "grpc hello " + args.GetValue()}
	return reply, nil
}

func main(){
	grpcServer := grpc.NewServer()
	go_protoc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
