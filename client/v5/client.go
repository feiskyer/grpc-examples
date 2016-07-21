/*
 * Version 5 of grpc helloworld service
 * This version adds SayNewHello which returns a stream resp
 * To keep compatible with old versions (e.g. v2), it remains SayHello and will
 * calls SayHello when server's version is below 0.5.
 */
package main

import (
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	pb "github.com/feiskyer/grpc-examples/api/v5/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	version     = "0.5.0"
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

	serverVersion, err := semver.NewVersion(v.Version)
	if err != nil {
		log.Fatalf("Can not parse server version: %v", err)
	}
	minVersion, err := semver.NewVersion(version)
	if err != nil {
		log.Fatalf("Can not parse client version: %v", err)
	}

	// For old versioned server, call original method
	if serverVersion.LessThan(*minVersion) {
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{
			Name: name,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Greeting: %s", r.Message)
		return
	}

	r, err := c.SayNewHello(context.Background(), &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	resp, err := r.Recv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
