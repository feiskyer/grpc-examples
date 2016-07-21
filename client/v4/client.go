/*
 * Version 4 of grpc helloworld service
 * This version changes SayHello to stream response.
 * All clients/servers of old versions will not work any more (can not compile actually)
 * because of introduced stream.
 */
package main

import (
	"log"
	"os"

	pb "github.com/feiskyer/grpc-examples/api/v4/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	version     = "0.4.0"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Request with client version: %s", version)
	v, err := c.Version(context.Background(), &pb.VersionRequest{
		Version: version,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Got server's version: %s", v.Version)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		User: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
