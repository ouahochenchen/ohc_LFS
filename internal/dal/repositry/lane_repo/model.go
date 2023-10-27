package lane_repo

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type LaneResourceTab struct {
	LaneId          uint64           `gorm:"column:lane_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"lane_id"`
	IsOk            sql.NullInt32    `gorm:"column:is_ok;type:tinyint(4);default:1" json:"is_ok"`
	LaneType        int16            `gorm:"column:lane_type;type:tinyint(4);NOT NULL" json:"lane_type"`
	LaneComposition LaneComposeSlice `gorm:"column:lane_composition;type:json;comment:存为list类的json，数据为lineid，siteid，seqence;NOT NULL" json:"lane_composition"`
	LaneName        string           `gorm:"column:lane_name;type:varchar(255);NOT NULL" json:"lane_name"`
	Operator        string           `gorm:"column:operator;type:varchar(255);NOT NULL" json:"operator"`
	CreateTime      int64            `gorm:"autoCreateTime;column:create_time;NOT NULL" json:"create_time"`
	UpdateTime      int64            `gorm:"autoUpdateTime;column:update_time;NOT NULL" json:"update_time"`
}

func (m *LaneResourceTab) TableName() string {
	return "lane_resource_tab"
}

type LaneCompose struct {
	Sequence     int  `json:"sequence"`
	ResourceId   uint `json:"resource_id"`
	ResourceType int  `json:"resource_type"`
}

type LaneComposeSlice []*LaneCompose

func (lcs *LaneComposeSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid value type")
	}
	err := json.Unmarshal(bytes, lcs)
	if err != nil {
		return err
	}
	return nil
}
func (lcs LaneComposeSlice) Value() (driver.Value, error) {
	return json.Marshal(lcs)
}
