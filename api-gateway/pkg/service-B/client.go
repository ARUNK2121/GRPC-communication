package B

import (
	"api-server/pkg/config"
	"api-server/pkg/service-B/pb"
	"api-server/pkg/service-B/procedure"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Client struct {
	Client pb.B_ServiceClient
}

func New_B_serviceClient(c *config.Config) *Client {
	cc, err := grpc.Dial(c.B_URL, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewB_ServiceClient(cc)

	return &Client{
		Client: pbClient,
	}
}

func (c *Client) FunctionB(ctx *gin.Context) {
	procedure.FunctionB(ctx, c.Client)
}

func (c *Client) FunctionB_TO_A(ctx *gin.Context) {
	procedure.FunctionB_TO_A(ctx, c.Client)
}
