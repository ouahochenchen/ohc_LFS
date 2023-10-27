package ls_connect_repo

type LaneSiteConnectConfigurationTab struct {
	Id             uint  `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	ResourceId     int   `gorm:"column:resource_id;type:int(11);NOT NULL" json:"resource_id"`
	NextResourceId int   `gorm:"column:next_resource_id;type:int(11);NOT NULL" json:"next_resource_id"`
	CreateTime     int64 `gorm:"autoCreateTime;column:create_time;NOT NULL" json:"create_time"`
	UpdateTime     int64 `gorm:"autoUpdateTime;column:update_time;NOT NULL" json:"update_time"`
}

func (m *LaneSiteConnectConfigurationTab) TableName() string {
	return "lane_site_connect_configuration_tab"
}
