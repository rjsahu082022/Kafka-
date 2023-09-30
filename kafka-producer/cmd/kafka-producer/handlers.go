package kafkaproducer

import (
	"kafka/golang-kafka/kafka-producer/api/v1/message"
)

func (a *Application) SetupHandlers() {
	message.RegisterHandlers(
		a.router,
		a.producer1,
		a.producer2,
		a.log,
	)
}
