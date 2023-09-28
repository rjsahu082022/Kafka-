package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"kafka/golang-kafka/kafka-consumer/internal/log"
	messageModel "kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	log            log.Logger
	messageService messageModel.Service
	consumers      *kafka.Reader
}

func NewConsumer(log log.Logger, messageService messageModel.Service, brokerAddress, topic []string) (*Consumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerAddress,
		//	Topic:    "my-topic",
		GroupTopics: topic,
		GroupID:     "my-group",
		MinBytes:    10e3,
		MaxBytes:    10e6,
	})

	r.SetOffset(kafka.LastOffset)

	return &Consumer{
		log:            log,
		messageService: messageService,
		consumers:      r,
	}, nil
}

func (con *Consumer) Consume(ctx context.Context, topics []string) error {
	for {
		m, err := con.consumers.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		pay, err := bytesToMessage(m.Value)
		if err != nil {
			return err
		}
		if err := con.messageService.Add(ctx, con.log, pay); err != nil {
			return fmt.Errorf("error while adding data %s", err)
		}
		fmt.Println("Received message: ", pay)
	}

}

func bytesToMessage(b []byte) (messageModel.Message, error) {
	var p messageModel.Message
	err := json.Unmarshal(b, &p)
	return p, err
}
