package main

import (
	"fmt"
	"log"
	"net"
	"service-A/pkg/config"
	"service-A/pkg/db"
	"service-A/pkg/pb"
	"service-A/pkg/server"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := db.Init(c.DB_URL)

	lis, err := net.Listen("tcp", c.PORT)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("service A on", c.PORT)

	s := server.NewServer(db)

	grpcServer := grpc.NewServer()

	pb.RegisterA_ServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
