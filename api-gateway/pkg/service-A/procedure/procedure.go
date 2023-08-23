package procedure

import (
	"api-server/pkg/service-A/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FunctionA(ctx *gin.Context, c pb.A_ServiceClient) {
	sample := SampleRequest{}

	if err := ctx.BindJSON(&sample); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FunctionA(context.Background(), &pb.SampleRequest{
		Id:   int64(sample.ID),
		Name: sample.Name,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
