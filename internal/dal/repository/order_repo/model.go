package order_repo

import (
	"database/sql"
)

type LaneOrderTab struct {
	OrderId       uint64          `gorm:"column:order_id;type:bigint(20) unsigned;primary_key" json:"order_id"`
	BuyerName     sql.NullString  `gorm:"column:buyer_name;type:varchar(255)" json:"buyer_name"`
	BuyerAddress  sql.NullString  `gorm:"column:buyer_address;type:varchar(255)" json:"buyer_address"`
	BuyerPhone    sql.NullString  `gorm:"column:buyer_phone;type:varchar(255)" json:"buyer_phone"`
	GoodsType     uint64          `gorm:"column:goods_type;type:int(11);default:0;NOT NULL" json:"goods_type"`
	SellerName    sql.NullString  `gorm:"column:seller_name;type:varchar(255)" json:"seller_name"`
	SellerAddress sql.NullString  `gorm:"column:seller_address;type:varchar(255)" json:"seller_address"`
	SellerPhone   sql.NullString  `gorm:"column:seller_phone;type:varchar(255)" json:"seller_phone"`
	PackageHeight sql.NullInt32   `gorm:"column:package_height;type:int(11);default:0" json:"package_height"`
	PackageWeight sql.NullInt32   `gorm:"column:package_weight;type:int(11);default:0" json:"package_weight"`
	Price         sql.NullFloat64 `gorm:"column:price;type:decimal(10,4)" json:"price"`
	OrderStatus   uint64          `gorm:"column:order_status;type:tinyint(4);default:0;NOT NULL" json:"order_status"`
	LaneId        uint64          `gorm:"column:lane_id;type:int(11);NOT NULL" json:"lane_id"`
	OmsOrderId    uint64          `gorm:"column:oms_order_id;type:int(11);NOT NULL" json:"oms_order_id"`
	CreateTime    uint64          `gorm:"autoCreateTime;column:create_time;NOT NULL" json:"create_time"`
	UpdateTime    uint64          `gorm:"autoUpdateTime;column:update_time;NOT NULL" json:"update_time"`
}

func (m *LaneOrderTab) TableName() string {
	return "lane_order_tab"
}
