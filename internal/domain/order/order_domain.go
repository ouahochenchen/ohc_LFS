package order

import (
	"LFS/internal/dal/repositry/order_repo"
	"LFS/internal/infrastructure/algo"
	"LFS/internal/infrastructure/err_code"
	"LFS/protocol/api"
)

type OrderDomain interface {
	CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error)
}
type oderDomainImpl struct {
	orderService order_repo.OrderRepo
	algoService  algo.AlgoService
}

func NewDomainImpl(repo order_repo.OrderRepo) OrderDomain {
	return &oderDomainImpl{
		orderService: repo,
	}
}
func (o *oderDomainImpl) CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error) {
	tab, err := o.orderService.SelectById(req.OrmOrderId)
	if err != nil {
		return nil, err
	}
	if tab != nil {
		return nil, &err_code.MyError{Msg: "已有重复订单"}
	}
	canDeliver, err := o.algoService.IsLaneCanDeliver(req.LaneId)
	if err != nil {
		return nil, err
	}
	if canDeliver == false {
		return nil, &err_code.MyError{Msg: "链路不可达"}
	}
	return nil, nil
}
