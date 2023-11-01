package task

import (
	"LFS/internal/dal/repositry/lane_repo"
	"LFS/internal/dal/repositry/order_repo"
	"LFS/internal/domain/lane"
	"LFS/internal/domain/order"
)

type TaskUseCase interface {
	SelectOrder(ormId uint64) *order_repo.LaneOrderTab
	SelectLaneCompose(laneId uint64) *lane_repo.LaneResourceTab
}
type taskUseCaseImpl struct {
	laneService  lane.LaneDomain
	orderService order.OrderDomain
}

func NewTaskUseCase(laneService lane.LaneDomain, orderService order.OrderDomain) TaskUseCase {
	return &taskUseCaseImpl{
		laneService:  laneService,
		orderService: orderService,
	}
}
func (task *taskUseCaseImpl) SelectOrder(ormId uint64) *order_repo.LaneOrderTab {
	orderTab := task.orderService.SelectById(ormId)
	return orderTab
}
func (task *taskUseCaseImpl) SelectLaneCompose(laneId uint64) *lane_repo.LaneResourceTab {
	laneTab := task.laneService.SelectById(laneId)
	return laneTab

}
