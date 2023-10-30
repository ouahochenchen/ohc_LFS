package ls_connect_repo

import "LFS/initialize"

var dbConnect = initialize.MasterDb

type ConnectRepo interface {
	Create(tab *LaneSiteConnectConfigurationTab) (uint64, error)
	Delete(id uint64) error
	SelectWithAlgo(rId uint64, rType uint64) ([]*LaneSiteConnectConfigurationTab, error)
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
func (*connectRepoImpl) SelectWithAlgo(rId uint64, rType uint64) ([]*LaneSiteConnectConfigurationTab, error) {
	record := make([]*LaneSiteConnectConfigurationTab, 0)
	err := dbConnect.Where("resource_id=? and resource_type=?", rId, rType).Find(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}
