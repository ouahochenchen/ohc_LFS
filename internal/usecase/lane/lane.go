package lane

import (
	"LFS/internal/domain/lane"
	"LFS/protocol/admin"
	"LFS/protocol/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LaneUseCase interface {
	CreateLane(ctx *gin.Context, info interface{}) (interface{}, error)
	UpdateLane(ctx *gin.Context)
	PageSelect(ctx *gin.Context)
}

type laneUseCaseImpl struct {
	laneService lane.LaneDomain
}

func NewLaneUseCase(laneService lane.LaneDomain) LaneUseCase {
	return &laneUseCaseImpl{
		laneService: laneService,
	}
}

func (l *laneUseCaseImpl) CreateLane(ctx *gin.Context, info interface{}) (interface{}, error) {
	// todo
	var req = info.(*admin.CreateLaneRequest)
	return l.laneService.Create(req)
}

func (l *laneUseCaseImpl) UpdateLane(ctx *gin.Context) {
	var req admin.UpdateLaneRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    "get param fail",
			//Data:
		})
	}
	update, err := l.laneService.Update(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    "get param fail",
			//Data: &admin.CreateLaneResponse{
			//
			//},
		})
		return
	}
	ctx.JSON(http.StatusOK, common.HttpCommonResponse{
		ReturnCode: 0,
		Message:    "OK",
		Data: &admin.UpdateLaneResponse{
			LaneId: update,
		},
	})
}
func (l *laneUseCaseImpl) PageSelect(ctx *gin.Context) {
	defer func() {
		er := recover()
		if er != nil {
			fmt.Println(er)
		}
	}()
	var req admin.PageSelectLaneRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    fmt.Sprintf("get param fail: %s", err.Error()),
		})
		return
	}
	pageSelect, err := l.laneService.PageSelect(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    "get param fail",
		})
		return
	}
	ctx.JSON(http.StatusOK, common.HttpCommonResponse{
		ReturnCode: 0,
		Message:    "OK",
		Data:       pageSelect,
	})
}
