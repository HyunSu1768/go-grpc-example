package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"context"
	pb "go-grpc-example/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Print("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRuntimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Version(ctx, &pb.VersionRequest{Version: "1"})
	if err != nil {
		fmt.Print("could not request: %v", err)
	}

	fmt.Printf("Config: %v", r)
}
