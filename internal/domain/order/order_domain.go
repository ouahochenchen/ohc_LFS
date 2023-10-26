package order

import "LFS/internal/dal/repositry/order_repo"

type OrderDomain interface {
	CheckDup(ormOrderId uint64) (bool, error)
}
type oderDomainImpl struct {
}

func NewDomainImpl() OrderDomain {
	return &oderDomainImpl{}
}
func (*oderDomainImpl) CheckDup(ormOrderId uint64) (bool, error) {
	id, err := order_repo.NewOrderRepo().SelectById(ormOrderId)
	if err != nil {
		return false, err
	}
	if id != nil {
		return false, nil
	}
	return true, nil
}
