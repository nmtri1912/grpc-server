package main

import (
	"log"
	"net"

	"grpc-server/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Cannot to listen port 9090")
	}

	server := chat.NewServer()
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Fail to Serve grpc Server")
	}

}
