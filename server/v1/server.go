/*
 * First version of grpc helloworld service
 * We will keep server always in v1 and change the version to check what will happen
 */
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
	version = "0.1.0"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Got SayHello request: %s", in.String())
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func (s *server) Version(ctx context.Context, in *pb.VersionRequest) (*pb.VersionResponse, error) {
	log.Printf("Got version request: %s", in.String())
	return &pb.VersionResponse{
		Version: version,
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
