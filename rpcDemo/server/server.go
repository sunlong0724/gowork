package main

import (
	"log"
	"net"
	"net/rpc"
	"rpcDemo/pb"
)
type HelloService struct {
}

func (this *HelloService) Hello(request *go_protoc.String, reply *go_protoc.String) error{
	reply.Value = "hello " + request.GetValue()
	return nil
}

func main(){
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()

	if err != nil{
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)

}