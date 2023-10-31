package algo

import (
	"LFS/internal/dal/repositry/lane_repo"
	"LFS/internal/dal/repositry/ls_connect_repo"
	"LFS/internal/infrastructure/err_code"
	"sort"
)

type AlgoService interface {
	IsLaneCanDeliver(laneId uint64) (bool, error)
}

type algoServiceImpl struct {
	laneService    lane_repo.LaneRepo
	connectService ls_connect_repo.ConnectRepo
}

func NewAlgoService(laneService lane_repo.LaneRepo, connectService ls_connect_repo.ConnectRepo) AlgoService {
	return &algoServiceImpl{
		laneService:    laneService,
		connectService: connectService,
	}
}

func (a *algoServiceImpl) IsLaneCanDeliver(laneId uint64) (bool, error) {
	//根据链路id查出链路，并取出链路中的点线组成
	result, err := a.laneService.SelectById(laneId)
	if err != nil {
		return false, err
	}
	if result.LaneComposition == nil {
		return false, &err_code.MyError{Msg: "没有此链路"}
	}
	composition := result.LaneComposition

	sort.Slice(composition, func(i, j int) bool {
		return composition[i].Sequence > composition[j].Sequence
	})
	//循环判断点线资源，判断当前资源是否在上一个资源的可达资源组中，根据id和type取出资源下一个可达资源组
	// 点1-线1-点2-线2-点3
	canDeliver := new([]*ls_connect_repo.LaneSiteConnectConfigurationTab)

	for _, v := range composition {
		if canDeliver != nil {
			var flag bool = false
			for _, v1 := range *canDeliver {
				if v1.NextResourceId == v.ResourceId && v1.NextType == v.ResourceType {
					flag = true
					break
				}
			}
			if flag == false {
				return false, &err_code.MyError{Msg: "链路不可达"}
			}
		}
		algo, err := a.connectService.SelectWithAlgo(v.ResourceId, v.ResourceType)
		if err != nil {
			return false, err
		}
		canDeliver = &algo
	}
	return true, nil

}
