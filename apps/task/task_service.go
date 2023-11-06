package task

import (
	"LFS/internal/dal/repository/lane_repo"
	"LFS/internal/dal/repository/order_repo"
	lane2 "LFS/internal/domain/lane"
	order2 "LFS/internal/domain/order"
	"LFS/internal/usecase/task"
)

var taskService taskApp

type taskApp struct {
	taskUseCase task.TaskUseCase
}

func NewApiApp(task task.TaskUseCase) *taskApp {
	return &taskApp{
		taskUseCase: task,
	}
}
func init() {
	repo := lane_repo.NewLaneRepo()
	orderRepo := order_repo.NewOrderRepo()
	laneDomain := lane2.NewLaneDomain(repo)
	orderDomain := order2.NewDomainImpl(orderRepo, nil)
	taskUseCase := task.NewTaskUseCase(laneDomain, orderDomain)
	taskService = *NewApiApp(taskUseCase)
}
