package main

import (
	"LFS/apps/api"
	"LFS/initialize"
	_ "LFS/initialize"
	"log"
)

func main() {
	api.RouterInit()
	err := initialize.InitKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Panicf("InitKafkaProducer fail: %s", err.Error())
	}
}
