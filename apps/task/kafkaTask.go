package task

import (
	"LFS/initialize"
	"LFS/internal/constant"
	"LFS/protocol/task"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func doTask() {
	service := initialize.KafkaService
	//msgSlice := new([]*task.ProduceMsg)
	err, msgChan := service.ConsumeMsg("LFS")
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("Consumer: Received message: topic=%s, partition=%d, offset=%d, value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
			var pcMsg task.ProduceMsg
			err := json.Unmarshal(msg.Value, &pcMsg)
			if err != nil {
			}
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
					}
				}()
			}
		case <-signals:
			break
		}
		wg.Wait()
	}

}
