package procedure

import (
	"api-server/pkg/service-B/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleRequest struct {
	Place string `json:"place"`
}

func FunctionB(ctx *gin.Context, c pb.B_ServiceClient) {
	sample := SampleRequest{}

	if err := ctx.BindJSON(&sample); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FunctionB(context.Background(), &pb.SampleRequest{
		Place: sample.Place,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func FunctionB_TO_A(ctx *gin.Context, c pb.B_ServiceClient) {
	sample := SampleRequest{}

	if err := ctx.BindJSON(&sample); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FunctionB_TO_A(context.Background(), &pb.SampleRequest{
		Place: sample.Place,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
