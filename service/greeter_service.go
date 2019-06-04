package service

import (
	"github.com/inagacky/go_grpc_sample/pb/helloworld"
	"golang.org/x/net/context"
	"log"
)

type GreeterService struct{}

func (s *GreeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println("call from", req.Name)
	rsp := helloworld.HelloReply{
		Message:"Hello " + req.Name + ".",
	}

	return &rsp, nil
}
