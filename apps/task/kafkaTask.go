package task

import (
	"LFS/initialize"
	"LFS/internal/constant"
	"LFS/internal/dal/invoker/lls/grpc_connect"
	_go "LFS/protocol/grpc/go"
	"LFS/protocol/task"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func DoTask() {
	service := initialize.KafkaService
	//msgSlice := new([]*task.ProduceMsg)
	err, msgChan := service.ConsumeMsg("test")
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	for {
		select {
		case msg := <-msgChan:
			//fmt.Printf("Consumer: Received message: topic=%s, partition=%d, offset=%d, value=%s\n",
			//	msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
			var pcMsg task.ProduceMsg
			err := json.Unmarshal(msg.Value, &pcMsg)
			if err != nil {
			}
			//726491287649386496
			order := taskService.taskUseCase.SelectOrder(pcMsg.OrderId)
			composeTab := taskService.taskUseCase.SelectLaneCompose(order.LaneId)
			compose := composeTab.LaneComposition
			for _, v := range compose {
				wg.Add(1)
				value := v
				go func() {
					defer wg.Done()
					if value.ResourceType == constant.LineType {
						//todo grpc request
						fmt.Println("狗儿伤", order.OrderId, value.ResourceId)
						req := &_go.LfsRequest{
							LfsOrderId:    order.OrderId,
							LineId:        value.ResourceId,
							BuyerName:     order.BuyerName.String,
							BuyerAddress:  order.BuyerAddress.String,
							BuyerPhone:    order.BuyerPhone.String,
							GoodsType:     order.GoodsType,
							SellerName:    order.SellerName.String,
							SellerAddress: order.SellerAddress.String,
							SellerPhone:   order.SellerPhone.String,
							PackageHeight: uint64(order.PackageHeight.Int32),
							PackageWeight: uint64(order.PackageWeight.Int32),
							Price:         float32(order.Price.Float64),
							OrderStatus:   order.OrderStatus,
						}
						resp, err2 := grpc_connect.LlsClientVa.LfsRpc(context.Background(), req)
						if err2 != nil {
							return
						}
						fmt.Printf("异步调用请求的返回值是:%s\n", resp)
					}
				}()
			}
		case <-signals:
			break
		}
		wg.Wait()
	}

}
