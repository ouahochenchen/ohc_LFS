package admin

import (
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
	laneService := lane2.NewLaneDomain()
	laneUsecase := lane.NewLaneUseCase(laneService)
	connectService := connect2.NewConnectDomain()
	connetUsecase := connect.NewConnectUseCase(connectService)
	admin = *NewAdminApp(laneUsecase, connetUsecase)
}
