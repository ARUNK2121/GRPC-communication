package server

import (
	"context"
	"fmt"
	"service-A/pkg/pb"

	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	pb.UnimplementedA_ServiceServer
}

func NewServer(db *gorm.DB) Server {
	return Server{
		DB: db,
	}
}

func (s *Server) FunctionA(c context.Context, req *pb.SampleRequest) (*pb.SampleResponse, error) {

	a := fmt.Sprintf("hii mr: %s, This message is from the service A", req.Name)

	return &pb.SampleResponse{
		Message: a,
	}, nil
}

func (s *Server) FunctionB_TO_A(c context.Context, req *pb.SampleRequest2) (*pb.SampleResponse, error) {

	fmt.Println("reached here")

	a := "the message reached service B and then reached service A"

	return &pb.SampleResponse{
		Message: a,
	}, nil
}
