package admin

import (
	"LFS/internal/dal/repository/lane_repo"
	"LFS/internal/dal/repository/ls_connect_repo"
	connect2 "LFS/internal/domain/connect"
	lane2 "LFS/internal/domain/lane"
	"LFS/internal/usecase/connect"
	"LFS/internal/usecase/lane"
)

var admin adminApp

type adminApp struct {
	laneUseCase    lane.LaneUseCase
	connectUseCase connect.ConnectUseCase
}

func NewAdminApp(laneUseCase lane.LaneUseCase, connectUseCase connect.ConnectUseCase) *adminApp {
	return &adminApp{
		laneUseCase:    laneUseCase,
		connectUseCase: connectUseCase,
	}
}

func init() {
	repo := lane_repo.NewLaneRepo()
	repo2 := ls_connect_repo.NewConnectRepo()
	laneService := lane2.NewLaneDomain(repo)
	laneUsecase := lane.NewLaneUseCase(laneService)
	connectService := connect2.NewConnectDomain(repo2)
	connetUsecase := connect.NewConnectUseCase(connectService)
	admin = *NewAdminApp(laneUsecase, connetUsecase)
}
