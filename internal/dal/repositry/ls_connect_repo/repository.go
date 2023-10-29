package ls_connect_repo

import "LFS/initialize"

var dbConnect = initialize.MasterDb

type ConnectRepo interface {
	Create(tab *LaneSiteConnectConfigurationTab) (uint64, error)
	Delete(id uint64) error
	/*继续。。。。。*/
}
type connectRepoImpl struct {
}

func NewConnectRepo() ConnectRepo {
	return &connectRepoImpl{}
}
func (*connectRepoImpl) Create(tab *LaneSiteConnectConfigurationTab) (uint64, error) {
	if tx := dbConnect.Create(tab); tx.Error != nil {
		return 0, tx.Error
	}
	return tab.Id, nil
}

func (*connectRepoImpl) Delete(id uint64) error {
	var tab LaneSiteConnectConfigurationTab
	db := dbConnect.Delete(&tab).Where("id=?", id)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
