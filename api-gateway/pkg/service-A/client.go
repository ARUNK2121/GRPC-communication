package A

import (
	"api-server/pkg/config"
	"api-server/pkg/service-A/pb"
	"api-server/pkg/service-A/procedure"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Client struct {
	Client pb.A_ServiceClient
}

func New_A_serviceClient(c *config.Config) *Client {
	cc, err := grpc.Dial(c.A_URL, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewA_ServiceClient(cc)

	return &Client{
		Client: pbClient,
	}
}

func (c *Client) FunctionA(ctx *gin.Context) {
	procedure.FunctionA(ctx, c.Client)
}
