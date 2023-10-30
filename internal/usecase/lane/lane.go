package lane

import (
	"LFS/internal/domain/lane"
	"LFS/protocol/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LaneUseCase interface {
	CreateLane(ctx *gin.Context, info interface{}) (interface{}, error)
	UpdateLane(ctx *gin.Context, info interface{}) (interface{}, error)
	PageSelect(ctx *gin.Context, info interface{}) (interface{}, error)
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

func (l *laneUseCaseImpl) UpdateLane(ctx *gin.Context, info interface{}) (interface{}, error) {
	var req = info.(*admin.UpdateLaneRequest)
	return l.laneService.Update(req)

}
func (l *laneUseCaseImpl) PageSelect(ctx *gin.Context, info interface{}) (interface{}, error) {
	defer func() {
		er := recover()
		if er != nil {
			fmt.Println(er)
		}
	}()
	var req = info.(*admin.PageSelectLaneRequest)
	return l.laneService.PageSelect(req)
}
