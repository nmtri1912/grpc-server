package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server interface {
	SayHello(context.Context, *RequestMessage) (*ResponseMessage, error)
}

type server struct {
}

func (s *server) SayHello(ctx context.Context, request *RequestMessage) (*ResponseMessage, error) {
	log.Printf("Received message body from client: %s", request.Body)
	return &ResponseMessage{
		Body: "hello from server hello",
	}, nil
}

func NewServer() Server {
	return &server{}
}
