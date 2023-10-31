package task

type ProduceMsg struct {
	OrderId       uint64  `json:"order_id"`
	LaneId        uint64  `json:"lane_id"`
	OrderStatus   uint64  `json:"order_status"`
	BuyerName     string  `json:"buyer_name"`
	BuyerAddress  string  `json:"buyer_address"`
	BuyerPhone    string  `json:"buyer_phone"`
	GoodsType     uint64  `json:"goods_type"`
	SellerName    string  `json:"seller_name"`
	SellerAddress string  `json:"seller_address"`
	SellerPhone   string  `json:"seller_phone"`
	PackageHeight uint64  `json:"package_height"`
	PackageWeight uint64  `json:"package_weight"`
	Price         float64 `json:"price"`
}

//OrmOrderId    uint64  `json:"orm_order_id" binding:"required"`
//BuyerName     string  `json:"buyer_name" binding:"required"`
//BuyerAddress  string  `json:"buyer_address" binding:"required"`
//BuyerPhone    string  `json:"buyer_phone" binding:"required"`
//GoodsType     uint64  `json:"goods_type" binding:"required"`
//SellerName    string  `json:"seller_name" binding:"required"`
//SellerAddress string  `json:"seller_address" binding:"required"`
//SellerPhone   string  `json:"seller_phone" binding:"required"`
//PackageHeight uint64  `json:"package_height" binding:"required"`
//PackageWeight uint64  `json:"package_weight" binding:"required"`
//Price         float64 `json:"price" binding:"required"`
//LaneId        uint64  `json:"lane_id" binding:"required"`
