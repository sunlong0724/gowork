package main

import (
	"fmt"
	"log"
	"net/rpc"
	go_protoc "rpcDemo/pb"
)

func main(){
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil{
		log.Fatal("dialing....", err)
	}

	var reply = &go_protoc.String{}
	var param = &go_protoc.String{
		Value: "jd .......",
	}
	err = client.Call("HelloService.Hello", &param, &reply)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
}
