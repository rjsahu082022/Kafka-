package kafkaconsumer

import "kafka/golang-kafka/kafka-consumer/api/v1/messaging"

func (a *Application) SetupHandlers() {
	messaging.RegisterHandlers(
		a.router,
		a.services.messageSvc,
		a.log,
	)
}
