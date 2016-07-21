/*
 * Version 2 of grpc helloworld service
 * This version deprecated HelloRequest.Name and replaced it with HelloRequest.User
 * Since `name` is not removed yet, client could continue setting value for name in
 * order to be compatible with old versions.
 */
package main

import (
	"log"
	"os"

	pb "github.com/feiskyer/grpc-examples/api/v2/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	version     = "0.2.0"
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
		// Notes: name is depracated and maybe removed in later versions.
		// Still setting value for name to work with old versions
		Name: name,
		User: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
