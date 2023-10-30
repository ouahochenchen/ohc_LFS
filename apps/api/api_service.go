package api

import (
	"LFS/internal/dal/repositry/order_repo"
	order2 "LFS/internal/domain/order"
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
	orderRepo := order_repo.NewOrderRepo()
	orderDomian := order2.NewDomainImpl(orderRepo)
	orderUseCase := order.NewOrderUseCase(orderDomian)
	api = *NewApiApp(orderUseCase)
}
