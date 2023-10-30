package api

import (
	"LFS/apps/protocol_handler"
	api2 "LFS/protocol/api"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	r := gin.Default()
	r.POST("/order", protocol_handler.SimpleGateway(api.orderUseCase.CheckOrder, &api2.CheckDuplicateRequest{}))
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
