package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/hbd/rpc-examples/twirp/helloworld/helloworld"
)

const (
	port   = "8080"
	domain = "localhost"
)

func main() {
	client := pb.NewHelloWorldProtobufClient("http://"+domain+":"+port, &http.Client{})

	// Request to `POST /twirp/<package>.<Service>/<Method>`.
	//             POST /twirp/twitch.twirp.example.helloworld.HelloWorld/Hello
	//             i.e. pb.HelloWorldPathPrefix
	req := &pb.HelloReq{Subject: "World"}
	log.Printf("Making request:\n %#v\n", req)

	resp, err := client.Hello(context.Background(), req)
	if err == nil {
		log.Printf("Response from server: %s\n", resp.Text)
	}
}
