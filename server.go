package main

import (
	pb "github.com/inagacky/go_grpc_sample/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type greeterService struct{}

func (s *greeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("call from", req.Name)
	rsp := pb.HelloReply {
		Message:"Hello " + req.Name + ".",
	}

	return &rsp, nil
}

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterService{})
	s.Serve(l)

}
