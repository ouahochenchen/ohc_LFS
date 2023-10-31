package order

import (
	"LFS/initialize"
	"LFS/internal/dal/repositry/order_repo"
	"LFS/internal/infrastructure/algo"
	"LFS/internal/infrastructure/err_code"
	"LFS/internal/infrastructure/snow_flake"
	"LFS/protocol/api"
	"LFS/protocol/task"
	"database/sql"
)

type OrderDomain interface {
	CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error)
}
type oderDomainImpl struct {
	orderService order_repo.OrderRepo
	algoService  algo.AlgoService
}

func NewDomainImpl(repo order_repo.OrderRepo, algoRepo algo.AlgoService) OrderDomain {
	return &oderDomainImpl{
		orderService: repo,
		algoService:  algoRepo,
	}
}
func (o *oderDomainImpl) CheckOrder(req *api.CheckDuplicateRequest) (*api.CheckDuplicateResponse, error) {
	tab, err := o.orderService.SelectById(req.OrmOrderId)
	if err != nil {
		return nil, err
	}
	if tab.OrderId != 0 {
		return nil, &err_code.MyError{Msg: "已有重复订单"}
	}
	canDeliver, err := o.algoService.IsLaneCanDeliver(req.LaneId)
	if err != nil {
		return nil, err
	}
	if canDeliver == false {
		return nil, &err_code.MyError{Msg: "链路不可达"}
	}
	height := req.PackageHeight
	weight := req.PackageWeight
	if err != nil {
		return nil, err
	}
	if height > 50 || weight > 200 {
		return nil, &err_code.MyError{Msg: "包裹尺寸或重量过大"}
	}

	orderId, err := snow_flake.GetId(req.LaneId)
	if err != nil {
		return nil, err
	}
	orderTab := order_repo.LaneOrderTab{
		OrderId:       uint64(orderId),
		BuyerName:     sql.NullString{String: req.BuyerName, Valid: true},
		BuyerAddress:  sql.NullString{String: req.BuyerAddress, Valid: true},
		BuyerPhone:    sql.NullString{String: req.BuyerPhone, Valid: true},
		GoodsType:     req.GoodsType,
		SellerName:    sql.NullString{String: req.SellerName, Valid: true},
		SellerAddress: sql.NullString{String: req.SellerAddress, Valid: true},
		SellerPhone:   sql.NullString{String: req.SellerPhone, Valid: true},
		PackageHeight: sql.NullInt32{Int32: int32(req.PackageHeight), Valid: true},
		PackageWeight: sql.NullInt32{Int32: int32(req.PackageWeight), Valid: true},
		Price:         sql.NullFloat64{Float64: req.Price, Valid: true},
		OrderStatus:   0,
		LaneId:        req.LaneId,
		OmsOrderId:    req.OrmOrderId,
	}
	order, err := o.orderService.CreateOrder(&orderTab)
	order = orderTab.OrderId
	if err != nil {
		return nil, err
	}
	resp := api.CheckDuplicateResponse{
		OrmOrderId: orderTab.OmsOrderId,
		OrderId:    order,
		//IsOk:        true,
		OrderStatus: orderTab.OrderStatus,
	}
	sendMssg := task.ProduceMsg{
		OrderId:       order,
		LaneId:        orderTab.LaneId,
		OrderStatus:   orderTab.OrderStatus,
		BuyerName:     orderTab.BuyerName.String,
		BuyerAddress:  orderTab.BuyerAddress.String,
		BuyerPhone:    orderTab.BuyerPhone.String,
		GoodsType:     orderTab.GoodsType,
		SellerName:    orderTab.SellerName.String,
		SellerAddress: orderTab.SellerAddress.String,
		SellerPhone:   orderTab.SellerPhone.String,
		PackageHeight: uint64(orderTab.PackageHeight.Int32),
		PackageWeight: uint64(orderTab.PackageWeight.Int32),
		Price:         orderTab.Price.Float64,
	}
	//kafkaList := []string{"localhost:9092"}
	//service := kafka.NewKafkaService()
	//_, err1 := service.InitProduce(kafkaList)
	////_ = produce
	//if err1 != nil {
	//	return nil, err1
	//}
	err2 := initialize.KafkaProducer.ProduceMsg(sendMssg)
	if err2 != nil {
		return nil, err2
	}
	resp.IsOk = true
	return &resp, nil
}
