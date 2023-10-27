package ls_connect_repo

import "LFS/initialize"

var dbConnect = initialize.MasterDb

type ConnectRepo interface {
	Create(tab *LaneSiteConnectConfigurationTab)
}
type connectRepoImpl struct {
}

func NewConnectRepo() ConnectRepo {
	return &connectRepoImpl{}
}

func (*connectRepoImpl) Create(tab *LaneSiteConnectConfigurationTab) {
	dbConnect.Create(tab)
}
