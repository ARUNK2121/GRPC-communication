package server

import (
	"context"
	"errors"
	"fmt"
	client_to_a "service-B/pkg/client-A"
	client_pb "service-B/pkg/client-A/pb"
	"service-B/pkg/pb"

	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Client *client_to_a.Client
	pb.UnimplementedB_ServiceServer
}

func NewServer(db *gorm.DB, c *client_to_a.Client) Server {
	return Server{
		DB:     db,
		Client: c,
	}
}

func (s *Server) FunctionB(c context.Context, req *pb.SampleRequest) (*pb.SampleResponse, error) {

	a := fmt.Sprintf("Service B got a request from the place: %s", req.Place)

	return &pb.SampleResponse{
		Message: a,
	}, nil
}

func (s *Server) FunctionB_TO_A(c context.Context, req *pb.SampleRequest) (*pb.SampleResponse, error) {

	fmt.Println("it was here atleast")

	d := client_pb.SampleRequest2{
		Place: req.Place,
	}

	result, err := s.Client.Client.FunctionB_TO_A(c, &d)
	if err != nil {
		return &pb.SampleResponse{}, errors.New("error")
	}

	return &pb.SampleResponse{
		Message: result.Message,
	}, nil
}
