package api

type CheckDuplicateRequest struct {
	OrmOrderId    uint64 `json:"orm_order_id" binding:"required"`
	BuyerName     string `json:"buyer_name" binding:"required"`
	BuyerAddress  string `json:"buyer_address" binding:"required"`
	BuyerPhone    string `json:"buyer_phone" binding:"required"`
	GoodsType     uint64 `json:"goods_type" binding:"required"`
	SellerName    string `json:"seller_name" binding:"required"`
	SellerAddress string `json:"seller_address" binding:"required"`
	SellerPhone   string `json:"seller_phone" binding:"required"`
	PackageHeight string `json:"package_height" binding:"required"`
	Price         string `json:"price" binding:"required"`
	LaneId        uint64 `json:"lane_id" binding:"required"`
}
type CheckDuplicateResponse struct {
	OrmOrderId  uint64 `json:"orm_order_id"`
	OrderId     uint64 `json:"order_id"`
	IsOk        bool   `json:"is_ok"`
	OrderStatus uint64 `json:"order_status"`
}
