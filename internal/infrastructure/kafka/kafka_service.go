package kafka

import (
	"LFS/internal/infrastructure/err_code"
	"LFS/protocol/task"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
)

type KafkaService interface {
	InitProduce(brokerList []string) (sarama.SyncProducer, error)
	ProduceMsg(msg task.ProduceMsg, topic string) error
	InitConsumer(brokerList []string) (sarama.Consumer, error)
	ConsumeMsg(topic string, msgSlice *[]*task.ProduceMsg) (error, *task.ProduceMsg)
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
func (kafka *kafkaServiceImpl) ConsumeMsg(topic string, msgSlice *[]*task.ProduceMsg) (error, *task.ProduceMsg) {
	//defer func(consumer sarama.Consumer) {
	//	err := consumer.Close()
	//	if err != nil {
	//	}
	//}(kafka.consumer)
	partition := int32(0)
	partitionConsumer, err := kafka.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Failed to create partition consumer:", err.Error())
		return err, nil
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		err := partitionConsumer.Close()
		if err != nil {

		}
	}(partitionConsumer)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	//for {
	//	select {
	//	case msg := <-partitionConsumer.Messages():
	//		fmt.Printf("Received message: %+v\n", string(msg.Value))
	//	case <-signals:
	//		return
	//	}
	//}
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Consumer: Received message: topic=%s, partition=%d, offset=%d, value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
			var pcMsg task.ProduceMsg
			err := json.Unmarshal(msg.Value, &pcMsg)
			_ = append(*msgSlice, &pcMsg)
			if err != nil {
				return err, &pcMsg
			}
		case <-signals:
			return &err_code.MyError{Msg: "消费中止"}, nil
		}
	}

}
