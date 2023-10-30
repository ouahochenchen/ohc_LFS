package admin

import (
	"LFS/apps/protocol_handler"
	admin2 "LFS/protocol/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	r := gin.Default()
	fmt.Println("路由初始化成功")

	r.POST("/lane", protocol_handler.SimpleGateway(admin.laneUseCase.PageSelect, &admin2.PageSelectLaneRequest{}))
	laneGroup := r.Group("/lane")
	{
		laneGroup.POST("/create", protocol_handler.SimpleGateway(admin.laneUseCase.CreateLane, &admin2.CreateLaneRequest{}))
		laneGroup.POST("/update", protocol_handler.SimpleGateway(admin.laneUseCase.UpdateLane, &admin2.UpdateLaneRequest{}))

	}

	laneGroup1 := r.Group("/order")
	{
		laneGroup1.POST("")
	}

	laneGroup2 := r.Group("/lsConnect")
	{
		laneGroup2.POST("/create", protocol_handler.SimpleGateway(admin.connectUseCase.CreateConnect, &admin2.CreateConnectRequest{}))
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
