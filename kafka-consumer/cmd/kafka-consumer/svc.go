package kafkaconsumer

import (
	messageService "kafka/golang-kafka/kafka-consumer/internal/service/messaging"
	messageModel "kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"
	messageRepo "kafka/golang-kafka/kafka-consumer/internal/service/messaging/repository"

	"go.mongodb.org/mongo-driver/mongo"

	"kafka/golang-kafka/kafka-consumer/internal/config"
)

type services struct {
	messageSvc messageModel.Service
}

type repos struct {
	messageRepo messageModel.Repository
}

func buildServices(cfg *config.Config, db *mongo.Collection) *services {
	svc := &services{}
	repos := &repos{}
	repos.buildRepos(db)
	svc.buildMessagingService(repos)
	return svc
}

func (r *repos) buildRepos(db *mongo.Collection) {
	r.messageRepo = messageRepo.NewRepository(db)
}

func (s *services) buildMessagingService(repo *repos) {
	s.messageSvc = messageService.NewService(repo.messageRepo)
}
