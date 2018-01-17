package main

import (
	"context"
	"flag"
	"net/http"

	pb "github.com/hnakamur/hello-twirp"
	"github.com/twitchtv/twirp"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (hat *pb.HelloResp, err error) {
	if req.Subject == "poo" {
		return nil, twirp.InvalidArgumentError("Subject", "must not be poo")
	}
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}

// Run the implementation in a local server
func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	flag.Parse()

	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)
	http.ListenAndServe(*addr, mux)
}
