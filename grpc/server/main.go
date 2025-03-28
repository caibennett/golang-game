package main

import (
	"log"
	"net"
	"server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMathServer(grpcServer, proto.NewServer())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
