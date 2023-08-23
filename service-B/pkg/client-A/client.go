package client_to_a

import (
	"fmt"
	"service-B/pkg/client-A/pb"
	"service-B/pkg/config"

	"google.golang.org/grpc"
)

type SampleRequest struct {
	Place string `json:"place"`
}

type Client struct {
	Client pb.A_ServiceClient
}

func New_B_to_A_serviceClient(c *config.Config) *Client {
	cc, err := grpc.Dial(c.A_URL, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewA_ServiceClient(cc)

	return &Client{
		Client: pbClient,
	}
}
