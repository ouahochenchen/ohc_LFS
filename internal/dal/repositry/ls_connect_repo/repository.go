package ls_connect_repo

import "LFS/initialize"

var dbConnect = initialize.MasterDb

type ConnectRepo interface {
	Create(tab *LaneSiteConnectConfigurationTab) (uint64, error)
}
type connectRepoImpl struct {
}

func NewConnectRepo() ConnectRepo {
	return &connectRepoImpl{}
}

func (*connectRepoImpl) Create(tab *LaneSiteConnectConfigurationTab) (uint64, error) {
	if tx := dbConnect.Create(tab); tx.Error != nil {
		return -1, tx.Error
	}
	return tab.Id, nil
}
