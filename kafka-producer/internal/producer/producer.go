package producer

import (
	"context"
	"time"

	"kafka/golang-kafka/kafka-producer/internal/log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	log       log.Logger
	Producers *kafka.Writer
	Topic     string
}

func NewProducer(log log.Logger, address []string, topic string) (*Producer, error) {
	c, err := kafka.Dial("tcp", address[0])
	if err != nil {
		log.Fatal("unable to build producer:", err)
		return nil, err
	}
	kt := kafka.TopicConfig{Topic: topic, NumPartitions: 2, ReplicationFactor: 1}

	if err := c.CreateTopics(kt); err != nil {
		log.Fatal("unable to create topic:", err)
		return nil, err
	}

	w := kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: address,
			// Topic:        topic,
			RequiredAcks: int(kafka.RequireAll),
			MaxAttempts:  5,
		},
	)
	return &Producer{
		log:       log,
		Producers: w,
		Topic:     topic,
	}, nil
}

func (prod *Producer) Retry(value []byte) error {
	var err error
	for i := 0; i < 3; i++ {
		err = prod.Produce(value)
		if err == nil {
			prod.log.Infof("Produced message after retry %d", i)
			return nil
		}
		prod.log.Infof("Failed to produce message, retrying in 2 seconds...")
		time.Sleep(time.Second * 2)
	}

	prod.log.Error("Failed to produce message after retrying times")
	return err
}

func (prod *Producer) Produce(value []byte) error {

	err := prod.Producers.WriteMessages(context.Background(),
		kafka.Message{Value: value, Topic: prod.Topic},
	)
	if err != nil {
		prod.log.Fatal("Failed to send message:", err)
		if err := prod.Retry(value); err != nil {
			prod.log.Fatal("Failed to send message:", err)
			return err
		}

	}
	return nil
}
