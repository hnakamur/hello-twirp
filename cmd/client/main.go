package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	pb "github.com/hnakamur/hello-twirp"
)

func main() {
	serverURL := flag.String("server", "http://127.0.0.1:8080", "server URL")
	subject := flag.String("subject", "World", "subject to say hello")
	flag.Parse()

	client := pb.NewHelloWorldProtobufClient(*serverURL, &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: *subject})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Text) // prints "Hello World"
}
