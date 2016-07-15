package main

import (
	"log"
	"net"

	pb "github.com/feiskyer/grpc-examples/api/v1/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	version = "0.1"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Got request: %s", in.String())
	return &pb.HelloReply{
		Version: version,
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server started with version %s", version)
	s.Serve(lis)
}
