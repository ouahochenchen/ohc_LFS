package initialize

import "LFS/internal/infrastructure/kafka"

var KafkaProducer kafka.KafkaService

func InitKafkaProducer(kafkaList []string) error {
	KafkaProducer = kafka.NewKafkaService()
	_, err := KafkaProducer.InitProduce(kafkaList)
	return err
}
