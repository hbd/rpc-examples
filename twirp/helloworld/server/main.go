package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/hbd/rpc-examples/twirp/helloworld/helloworld"
)

const port = "8080"

// HelloWorldServer is used to maintain state for this service.
type HelloWorldServer struct{}

// Hello says Hello!
func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	log.Printf("Received request:\n %#v\n", *req)
	return &pb.HelloResp{Text: "Hello" + req.Subject}, nil
}

func main() {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)

	mux := http.NewServeMux()
	mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)

	log.Printf("Listening on %s...\n", port)
	http.ListenAndServe(":"+port, mux)
}
