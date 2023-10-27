package order_repo

import "LFS/initialize"

var dbLane = initialize.MasterDb

type OrderRepo interface {
	CreateOrder(laneOrder *LaneOrderTab) (uint64, error)
	UpdateOrder(laneOrder *LaneOrderTab) (uint64, error)
	SelectById(i uint64) (*LaneOrderTab, error)
	SelectWithPage(page uint64, pageSize uint64, laneId *uint64, laneName *string) ([]*LaneOrderTab, uint64, error)
}
type orderRepoImpl struct {
}

func NewOrderRepo() OrderRepo {
	return &orderRepoImpl{}
}
func (l *orderRepoImpl) CreateOrder(laneOrder *LaneOrderTab) (uint64, error) {
	r := dbLane.Create(laneOrder)
	if r.Error != nil {
		return -1, r.Error
	}
	return laneOrder.OrderId, nil
}
func (l *orderRepoImpl) UpdateOrder(laneOrder *LaneOrderTab) (uint64, error) {
	r := dbLane.Updates(laneOrder)
	if r.Error != nil {
		return -1, r.Error
	}
	return laneOrder.OrderId, nil
}
func (l *orderRepoImpl) SelectById(i uint64) (*LaneOrderTab, error) {
	var lane LaneOrderTab
	err := dbLane.Where("order_id=?", i).Find(&lane).Error
	return &lane, err
}
func (l *orderRepoImpl) SelectWithPage(page uint64, pageSize uint64, laneId *uint64, laneName *string) ([]*LaneOrderTab, uint64, error) {
	var orderRecord = make([]*LaneOrderTab, 0)
	//var tabs *LaneResourceTab
	var totalRecords uint64
	i := int64(totalRecords)
	var err2 error
	if laneId != nil && laneName != nil {
		err2 = dbLane.Where("order_id=?", *laneId).Where("order_id=?", *laneName).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&orderRecord).Error
	} else if laneId != nil {
		err2 = dbLane.Where("order_id=?", *laneId).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&orderRecord).Error
	} else if laneName != nil {
		err2 = dbLane.Where("order_id=?", *laneName).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&orderRecord).Error
	} else {
		err2 = dbLane.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&orderRecord).Error
	}
	dbLane.Model(&LaneOrderTab{}).Count(&i)
	return orderRecord, uint64(i), err2
}
