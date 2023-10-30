package order

import (
	"LFS/internal/domain/order"
	"LFS/protocol/api"
	"github.com/gin-gonic/gin"
)

type OrderUseCase interface {
	CheckOrder(ctx *gin.Context, info interface{}) (interface{}, error)
}
type orderUseCaseImpl struct {
	orderService order.OrderDomain
}

func NewOrderUseCase(orderService order.OrderDomain) OrderUseCase {
	return &orderUseCaseImpl{
		orderService: orderService,
	}
}
func (o *orderUseCaseImpl) CheckOrder(ctx *gin.Context, info interface{}) (interface{}, error) {
	request := info.(*api.CheckDuplicateRequest)
	return o.orderService.CheckOrder(request)
}
