package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 9000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// TODO implement proto
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
