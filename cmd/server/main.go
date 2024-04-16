package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.RuntimeServiceServer
}

func (s *server) Version(ctx context.Context, in *pb.VersionRequest) (*pb.VersionResponse, error) {
	log.Printf("Received version request: %v", in.Version)
	return &pb.VersionResponse{
		Version:           "1",
		RuntimeName:       "현수런타임",
		RuntimeApiVersion: "v1",
		RuntimeVersion:    "1.0.0",
	}, nil
}

func main() {
	fmt.Print("HI")
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRuntimeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
