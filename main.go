package main

import (
	"log"
	"net"

	"github.com/sazap10/grpc-hello-world/pkg/chat"
	pb "github.com/sazap10/grpc-hello-world/grpcgen/chat-service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Go gRPC Beginners Tutorial!")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
