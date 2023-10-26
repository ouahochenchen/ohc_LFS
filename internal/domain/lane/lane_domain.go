package lane

import (
	"LFS/internal/dal/repositry/lane_repo"
	"LFS/protocol/admin"
	"database/sql"
	"errors"
)

type LaneDomain interface {
	Create(request *admin.CreateLaneRequest) (uint64, error)
	Update(request *admin.UpdateLaneRequest) (uint64, error)
	PageSelect(request *admin.PageSelectLaneRequest) (*admin.PageSelectLaneResponse, error)
}

type laneDomainImpl struct {
}

func NewLaneDomain() LaneDomain {
	return new(laneDomainImpl)
}

func (l *laneDomainImpl) Create(req *admin.CreateLaneRequest) (uint64, error) {
	var laneP = lane_repo.LaneResourceTab{LaneName: req.LaneName, LaneType: req.LaneType, Operator: req.Operator, LaneComposition: req.LaneComposeSl} //
	create, err := lane_repo.NewLaneRepo().Create(&laneP)
	return uint64(create), err
}
func (l *laneDomainImpl) Update(request *admin.UpdateLaneRequest) (uint64, error) {
	lp := &lane_repo.LaneResourceTab{LaneId: request.LaneId, LaneName: request.LaneName, LaneType: request.LaneType,
		LaneComposition: request.LaneComposeSl, Operator: request.Operator, IsOk: sql.NullInt32{Int32: request.IsOk, Valid: true}}
	if lp.LaneId == 0 {
		return 0, errors.New("更新时laneid不能为0,嘻嘻")
	}
	lptab, err := lane_repo.NewLaneRepo().SelectById(lp.LaneId)
	if err != nil {
		return 0, err
	}
	if lptab == nil {
		return 0, errors.New("没有该条数据，无法更新")
	}
	b := lp.LaneType != 0
	if b && lp.LaneType != lptab.LaneType {
		lptab.LaneType = lp.LaneType
	}
	if lp.IsOk != lptab.IsOk {
		lptab.IsOk = lp.IsOk
	}
	if lp.Operator != lptab.Operator && lp.Operator != "" {
		lptab.Operator = lp.Operator
	}
	if lp.LaneName != lptab.LaneName && lp.LaneName != "" {
		lptab.LaneName = lp.LaneName
	}
	if lp.LaneComposition != nil {
		lptab.LaneComposition = lp.LaneComposition
	}
	err1 := lane_repo.NewLaneRepo().Update(lptab)
	if err1 != nil {
		return 0, err1
	}
	return lptab.LaneId, err1
}
func (l *laneDomainImpl) PageSelect(request *admin.PageSelectLaneRequest) (*admin.PageSelectLaneResponse, error) {
	record, total, err := lane_repo.NewLaneRepo().SelectWithPage(request.Page, request.PageSize)

	if err != nil {
		return nil, err
	}
	var recordResponse = new(admin.PageSelectLaneResponse)
	//tabs := make([]*lane_repo.LaneResourceTab, 0)
	//tabs = append(tabs, record)
	recordResponse = &admin.PageSelectLaneResponse{ResponseTab: record, Page: request.Page, PageSize: request.PageSize, Total: total}
	return recordResponse, err
}
