package kafkaconsumer

import (
	"kafka/golang-kafka/kafka-consumer/internal/consumer"
)

func (a *Application) BuildReader() (*consumer.Consumer, error) {

	consumer, err := consumer.NewConsumer(a.log, a.services.messageSvc, []string{a.cfg.Consumer.Address}, a.cfg.Consumer.Topic)
	if err != nil {
		return nil, err
	}
	return consumer, err
}
