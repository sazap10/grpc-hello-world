package chat

import (
	"log"

	"golang.org/x/net/context"
	pb "github.com/sazap10/grpc-hello-world/grpcgen/chat-service"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
