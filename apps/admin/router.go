package admin

import (
	"LFS/internal/usecase/lane"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	r := gin.Default()
	fmt.Println("路由初始化成功")

	r.POST("/lane", lane.NewLaneUseCase().PageSelect)
	laneGroup := r.Group("/lane")
	{
		laneGroup.POST("/create", lane.NewLaneUseCase().CreateLane)
		laneGroup.POST("/update", lane.NewLaneUseCase().UpdateLane)

	}
	laneGroup1 := r.Group("/order")
	{
		laneGroup1.POST("")
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
