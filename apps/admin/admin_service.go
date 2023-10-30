package admin

import (
	"LFS/internal/dal/repositry/lane_repo"
	connect2 "LFS/internal/domain/connect"
	lane2 "LFS/internal/domain/lane"
	"LFS/internal/usecase/connect"
	"LFS/internal/usecase/lane"
)

var admin AdminApp

type AdminApp struct {
	laneUseCase    lane.LaneUseCase
	connectUseCase connect.ConnectUseCase
}

func NewAdminApp(laneUseCase lane.LaneUseCase, connectUseCase connect.ConnectUseCase) *AdminApp {
	return &AdminApp{
		laneUseCase:    laneUseCase,
		connectUseCase: connectUseCase,
	}
}

func init() {
	repo := lane_repo.NewLaneRepo()
	laneService := lane2.NewLaneDomain(repo)
	laneUsecase := lane.NewLaneUseCase(laneService)
	connectService := connect2.NewConnectDomain()
	connetUsecase := connect.NewConnectUseCase(connectService)
	admin = *NewAdminApp(laneUsecase, connetUsecase)
}
