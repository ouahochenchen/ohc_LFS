package order

import (
	"LFS/internal/dal/repositry/order_repo"
	"LFS/protocol/api"
)

type OrderDomain interface {
	CheckDup(req *api.CheckDuplicateRequest) (bool, error)
}
type oderDomainImpl struct {
	orderService order_repo.OrderRepo
}

func NewDomainImpl(repo order_repo.OrderRepo) OrderDomain {
	return &oderDomainImpl{
		orderService: repo,
	}
}
func (o *oderDomainImpl) CheckDup(req *api.CheckDuplicateRequest) (bool, error) {
	id, err := o.orderService.SelectById(req.OrmOrderId)
	if err != nil {
		return false, err
	}
	if id != nil {
		return false, nil
	}
	return true, nil
}
