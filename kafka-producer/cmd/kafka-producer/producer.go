package kafkaproducer

import (
	"kafka/golang-kafka/kafka-producer/internal/producer"
)

func (a *Application) BuildWriter() (*producer.Producer, *producer.Producer) {
	producer1, err := producer.NewProducer(a.log, []string{a.cfg.Producer.Address}, a.cfg.Producer.Topic[0])
	if err != nil {
		panic(err)
	}
	producer2, err := producer.NewProducer(a.log, []string{a.cfg.Producer.Address}, a.cfg.Producer.Topic[1])
	if err != nil {
		panic(err)
	}
	return producer1, producer2
}

func (a *Application) CloseWriter() {
	a.producer1.Producers.Close()
	a.producer2.Producers.Close()

}
