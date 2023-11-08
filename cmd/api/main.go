package main

import (
	"LFS/apps/api"
	"LFS/initialize"
	_ "LFS/initialize"
	_ "LFS/internal/infrastructure/snow_flake"
	"log"
)

func main() {
	err := initialize.InitKafka([]string{"localhost:9092"})
	if err != nil {
		log.Panicf("InitKafka fail: %s", err.Error())
	}
	//grpc_connect.Init("localhost:50051")
	api.RouterInit()
}
