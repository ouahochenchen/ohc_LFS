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
	CreateLane(ctx *gin.Context)
	UpdateLane(ctx *gin.Context)
	PageSelect(ctx *gin.Context)
}

type laneUseCaseImpl struct {
}

func NewLaneUseCase() LaneUseCase {
	return &laneUseCaseImpl{}
}

func (l *laneUseCaseImpl) CreateLane(ctx *gin.Context) {
	// todo
	var req admin.CreateLaneRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    fmt.Sprintf("get param fail: %s", err.Error()),
			//Data: &admin.CreateLaneResponse{
			//
			//},
		})
		return
	}
	//var laneP = lane_repo.LaneResourceTab{LaneName: req.LaneName, LaneType: req.LaneType, Operator: req.Operator, LaneComposition: req.LaneComposeSl}
	//create, err := lane_repo.NewLaneRepo().Create(&laneP)
	create, err := lane.NewLaneDomain().Create(&req)
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
		Data: &admin.CreateLaneResponse{
			LaneId: create,
		},
	})
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
	update, err := lane.NewLaneDomain().Update(&req)
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
	var req *admin.PageSelectLaneRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.HttpCommonResponse{
			ReturnCode: -1,
			Message:    "get param fail",
		})
		return
	}
	pageSelect, err := lane.NewLaneDomain().PageSelect(req)
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
