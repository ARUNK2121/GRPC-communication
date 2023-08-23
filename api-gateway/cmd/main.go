package main

import (
	"log"

	"api-server/pkg/config"
	A "api-server/pkg/service-A"
	B "api-server/pkg/service-B"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	A_client := A.New_A_serviceClient(&config)
	B_client := B.New_B_serviceClient(&config)

	r.POST("/function_a", A_client.FunctionA)
	r.POST("/function_b", B_client.FunctionB)
	r.POST("/b_to_a", B_client.FunctionB_TO_A)

	r.Run(config.PORT)
}
