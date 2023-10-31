package kafka

type KafkaService interface {
	initProduce()
}
type kafkaServiceImpl struct {
}

func (kafka *kafkaServiceImpl) initProduce() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 等待所有 ISR（in sync replicas）确认
	config.Producer.Retry.Max = 3                      // 最大重试次数
	config.Producer.Return.Successes = true            // 是否返回成功的消息

	// 创建生产者
	producer, err := sarama.NewSyncProducer(bootstrapServers, config)
	if err != nil {
		return nil, err
	}
}
