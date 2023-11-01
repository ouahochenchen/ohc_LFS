package main

import (
	"LFS/apps/api"
	"LFS/initialize"
	_ "LFS/initialize"
	_ "LFS/internal/infrastructure/snow_flake"
	"log"
)

func main() {
	api.RouterInit()
	err := initialize.InitKafka([]string{"localhost:9092"})
	if err != nil {
		log.Panicf("InitKafka fail: %s", err.Error())
	}
}
