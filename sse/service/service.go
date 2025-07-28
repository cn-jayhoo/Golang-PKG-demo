package service

import (
	"fmt"
	"google.golang.org/grpc"
	"sse/service_api"
	"time"
)

type Person struct{}

func NewPerson() service_api.EventServiceServer {
	return &Person{}
}

func (p Person) StreamEvents(req *service_api.EventRequest, g grpc.ServerStreamingServer[service_api.EventResponse]) error {
	for i := 0; i < 10; i++ {
		if err := g.Send(&service_api.EventResponse{
			EventType: "message",
			Data:      fmt.Sprintf("Event %d for client %s", i, req.ClientId),
		}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
