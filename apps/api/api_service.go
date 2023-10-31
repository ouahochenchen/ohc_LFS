package api

import (
	"LFS/internal/dal/repositry/lane_repo"
	"LFS/internal/dal/repositry/ls_connect_repo"
	"LFS/internal/dal/repositry/order_repo"
	order2 "LFS/internal/domain/order"
	"LFS/internal/infrastructure/algo"
	"LFS/internal/usecase/order"
)

var api apiApp

type apiApp struct {
	orderUseCase order.OrderUseCase
}

func NewApiApp(o order.OrderUseCase) *apiApp {
	return &apiApp{
		orderUseCase: o,
	}
}
func init() {
	repo := lane_repo.NewLaneRepo()
	connectRepo := ls_connect_repo.NewConnectRepo()
	orderRepo := order_repo.NewOrderRepo()
	algoService := algo.NewAlgoService(repo, connectRepo)
	orderDomian := order2.NewDomainImpl(orderRepo, algoService)
	orderUseCase := order.NewOrderUseCase(orderDomian)
	api = *NewApiApp(orderUseCase)
}
