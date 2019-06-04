package main

import (
	"github.com/inagacky/go_grpc_sample/pb/cat"
	"github.com/inagacky/go_grpc_sample/pb/helloworld"
	ser "github.com/inagacky/go_grpc_sample/service"
	"google.golang.org/grpc"
	"log"
	"net"
)



func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &ser.GreeterService{})
	cat.RegisterCatServiceServer(s, &ser.CatService{})
	s.Serve(l)

}
