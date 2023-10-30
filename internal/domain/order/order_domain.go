package order

import (
	"LFS/internal/dal/repositry/order_repo"
	"LFS/internal/util"
	"LFS/protocol/api"
)

type OrderDomain interface {
	CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error)
}
type oderDomainImpl struct {
	orderService order_repo.OrderRepo
}

func NewDomainImpl(repo order_repo.OrderRepo) OrderDomain {
	return &oderDomainImpl{
		orderService: repo,
	}
}
func (o *oderDomainImpl) CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error) {
	id, err := o.orderService.SelectById(req.OrmOrderId)
	if err != nil {
		return nil, err
	}
	if id != nil {
		return nil, &util.MyError{"已有重复订单"}
	}
	return nil, nil
}
