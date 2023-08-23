package main

import (
	"fmt"
	"log"
	"net"
	client_to_a "service-B/pkg/client-A"
	"service-B/pkg/config"
	"service-B/pkg/db"
	"service-B/pkg/pb"
	"service-B/pkg/server"

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

	fmt.Println("service B on Svc on", c.PORT)

	client := client_to_a.New_B_to_A_serviceClient(&c)

	s := server.NewServer(db, client)

	grpcServer := grpc.NewServer()

	pb.RegisterB_ServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
