package ls_connect_repo

type LaneSiteConnectConfigurationTab struct {
	Id             uint64 `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	ResourceId     uint64 `gorm:"column:resource_id;type:int(11);NOT NULL" json:"resource_id"`
	NextResourceId uint64 `gorm:"column:next_resource_id;type:int(11);NOT NULL" json:"next_resource_id"`
	ResourceType   uint64 `gorm:"column:resource_type;type:int(11);comment:1是点2是线;NOT NULL" json:"resource_type"`
	NextType       uint64 `gorm:"column:next_type;type:int(11)" json:"next_type"`
	CreateTime     uint64 `gorm:"autoCreateTime;column:create_time;NOT NULL" json:"create_time"`
	UpdateTime     uint64 `gorm:"autoUpdateTime;column:update_time;NOT NULL" json:"update_time"`
}

func (m *LaneSiteConnectConfigurationTab) TableName() string {
	return "lane_site_connect_configuration_tab"
}
