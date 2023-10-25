package lane_repo

import (
	"LFS/initialize"
	"fmt"
)

var dbLane = initialize.MasterDb

// LaneRepo 写增删改查方法
type LaneRepo interface {
	Create(laneResource *LaneResourceTab) (uint64, error)
	CreateBatch(laneResources []*LaneResourceTab) error
	Update(laneResource *LaneResourceTab) error
	SelectById(i uint64) (*LaneResourceTab, error)
	SelectWithPage(page uint64, pageSize uint64) ([]*LaneResourceTab, error)
}

type laneRepoImpl struct {
}

func NewLaneRepo() LaneRepo {
	return &laneRepoImpl{}
}

func (l *laneRepoImpl) CreateBatch(laneResources []*LaneResourceTab) error {
	result := dbLane.Model(&LaneResourceTab{}).CreateInBatches(laneResources, 100)
	return result.Error
}
func (l *laneRepoImpl) Create(laneResource *LaneResourceTab) (uint64, error) {
	re := dbLane.Create(laneResource)
	if re.Error != nil {
		return 0, re.Error
	}
	return laneResource.LaneId, nil
}
func (l *laneRepoImpl) Update(tab *LaneResourceTab) error {
	if tab.LaneId == 0 {
		return fmt.Errorf("没主键更毛线")
	}
	err := dbLane.Updates(tab).Error
	return err
}
func (l *laneRepoImpl) SelectById(i uint64) (*LaneResourceTab, error) {
	var lane LaneResourceTab
	err := dbLane.Where("lane_id=?", i).Find(&lane).Error
	return &lane, err
}
func (l *laneRepoImpl) SelectWithPage(page uint64, pageSize uint64) ([]*LaneResourceTab, error) {
	var laneRecord = make([]*LaneResourceTab, 0)
	//var tabs *LaneResourceTab
	//var totalRecords int64
	err2 := dbLane.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&laneRecord).Error
	return laneRecord, err2
}
