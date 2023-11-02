package main

import (
	"LFS/apps/task"
	"LFS/initialize"
	"log"
)

func main() {
	err := initialize.InitKafka([]string{"localhost:9092"})
	if err != nil {
		log.Panicf("InitKafka fail: %s", err.Error())
	}
	task.DoTask()
}
