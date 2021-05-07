package kfk

import (
	"github.com/Shopify/sarama"
	"github.com/ducksoso/kitex/queue"
	"go.uber.org/zap"
)

type KafkaConfig struct {
	Brokers       []string
	RetryMax      int
	ReturnSuccess bool
}

/**
生产者分为：同步生产者、异步生产者

*/

// KafkaSyncProducer sync producer
type KafkaSyncProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(config *KafkaConfig) (*KafkaSyncProducer, error) {
	conf := sarama.NewConfig()

	conf.Producer.Retry.Max = config.RetryMax
	conf.Producer.Return.Successes = config.ReturnSuccess

	producer, err := sarama.NewSyncProducer(config.Brokers, conf)
	if err != nil {
		zap.S().Errorf("init producer failed -> %v\n", err)
		return nil, err
	}
	zap.S().Infof("producer init success")

	return &KafkaSyncProducer{producer: producer}, nil
}

func (p *KafkaSyncProducer) Send(topic, msg string) {
	msgx := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	zap.S().Infof("SendMsg -> %v\n", queue.DumpString(msgx))

	partition, offset, err := p.producer.SendMessage(msgx)
	if err != nil {
		zap.S().Errorf("send msg error: %s\n", err)
	}
	zap.S().Infof("msg send success, message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}

// KafkaAsyncProducer async producer
type KafkaAsyncProducer struct {
	producer sarama.AsyncProducer
}

func NewAsyncProducer(config *KafkaConfig) (*KafkaAsyncProducer, error) {
	conf := sarama.NewConfig()

	client, err := sarama.NewClient(config.Brokers, conf)
	if err != nil {
		panic(err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	producer.Input()

	return &KafkaAsyncProducer{producer: producer}, nil
}

func (p *KafkaAsyncProducer) Send() {

}
