package kafka

import (
	"LFS/protocol/task"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type KafkaService interface {
	InitProduce(brokerList []string) (sarama.SyncProducer, error)
	ProduceMsg(msg task.ProduceMsg, topic string) error
	InitConsumer(brokerList []string) (sarama.Consumer, error)
	ConsumeMsg(topic string) (error, <-chan *sarama.ConsumerMessage)
}
type kafkaServiceImpl struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
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
func (kafka *kafkaServiceImpl) ProduceMsg(msg task.ProduceMsg, topic string) error {
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	//topic := "test"
	sendMsg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(0),
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
func (kafka *kafkaServiceImpl) InitConsumer(brokerList []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 连接到 Kafka broker
	// 创建消费者
	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		fmt.Println("Failed to connect to Kafka:", err.Error())
		return nil, err
	}
	kafka.consumer = consumer
	return consumer, nil
}

//func (kafka *kafkaServiceImpl) ConsumeMsg(topic string, msgSlice *[]*task.ProduceMsg) (error, <-chan *sarama.ConsumerMessage) {

func (kafka *kafkaServiceImpl) ConsumeMsg(topic string) (error, <-chan *sarama.ConsumerMessage) {
	//var wg sync.WaitGroup
	partition := int32(0)
	partitionConsumer, err := kafka.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

	if err != nil {
		fmt.Println("Failed to create partition consumer:", err.Error())
		return err, nil
	}
	return err, partitionConsumer.Messages()
}
