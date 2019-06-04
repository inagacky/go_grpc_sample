package main

import (
	"github.com/inagacky/go_grpc_sample/pb/cat"
	"github.com/inagacky/go_grpc_sample/pb/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	catSearch(conn)

	// HelloWorldのコール
//	callHelloWorld(conn)


}

func callHelloWorld(conn *grpc.ClientConn) {
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}


func catSearch(conn *grpc.ClientConn) {

	c := cat.NewCatServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.Search(ctx, &cat.CatSearchRequest{Type:cat.Type_SCOTTISH_FOLD})
	if err != nil {
		log.Fatalf("could not Cat: %v", err)
	}
	log.Printf("Cat: %s", r.Cat)
}
