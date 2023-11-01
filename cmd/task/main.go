package main

import (
	"LFS/initialize"
	"LFS/protocol/task"
)

func main() {
	service := initialize.KafkaService
	msgSlice := new([]*task.ProduceMsg)
	err := service.ConsumeMsg("LFS", msgSlice)
	if err != nil {
		return
	}

}
