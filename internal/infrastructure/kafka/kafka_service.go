package kafka

import (
	"LFS/protocol/task"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type KafkaService interface {
	InitProduce(brokerList []string) (sarama.SyncProducer, error)
	ProduceMsg(msg task.ProduceMsg) error
}
type kafkaServiceImpl struct {
	producer sarama.SyncProducer
}

func NewKafkaService() KafkaService {
	return &kafkaServiceImpl{}
}
func (kafka *kafkaServiceImpl) InitProduce(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 等待所有 ISR（in sync replicas）确认
	config.Producer.Retry.Max = 5                      // 最大重试次数
	config.Producer.Return.Successes = true            // 是否返回成功的消息
	//brokerList := []string{"localhost:9092"}
	// 创建生产者
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}
	kafka.producer = producer
	return producer, nil
}
func (kafka *kafkaServiceImpl) ProduceMsg(msg task.ProduceMsg) error {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	topic := "test"
	sendMsg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(1),
		Value:     sarama.StringEncoder(msgJson),
	}
	partition, offset, err := kafka.producer.SendMessage(sendMsg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
