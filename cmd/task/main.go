package main

import (
	"LFS/apps/task"
	"LFS/initialize"
	"LFS/internal/dal/invoker/lls_invoker/grpc_connect"
	"log"
)

func main() {
	err := initialize.InitKafka([]string{"localhost:9092"})
	if err != nil {
		log.Panicf("InitKafka fail: %s", err.Error())
	}
	grpc_connect.InitLLS("localhost:50051")
	task.DoTask()
}
