package order

type OrderDomain interface {
	CheckDup(ormOrderId uint64)
}
type oderDomainImpl struct {
}

func NewDomainImpl() OrderDomain {
	return &oderDomainImpl{}
}
func (*oderDomainImpl) CheckDup(ormOrderId uint64) {

}
