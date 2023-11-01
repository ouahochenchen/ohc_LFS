package initialize

import "LFS/internal/infrastructure/kafka"

var KafkaService kafka.KafkaService

func InitKafka(kafkaList []string) error {
	KafkaService = kafka.NewKafkaService()
	_, err := KafkaService.InitProduce(kafkaList)
	_, err = KafkaService.InitConsumer(kafkaList)
	return err
}
