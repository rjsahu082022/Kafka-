package messaging

import (
	"context"

	"kafka/golang-kafka/kafka-consumer/internal/log"
	"kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"
)

type Service struct {
	repo model.Repository
}

func NewService(repo model.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(ctx context.Context, log log.Logger) ([]*model.Message, error) {
	messages, err := s.repo.Get(ctx, log)
	if err != nil {
		log.Errorf("failed to fetch records %v", err)
		return nil, err
	}
	return messages, nil
}

func (s *Service) Add(ctx context.Context, log log.Logger, records model.Message) error {
	err := s.repo.Add(ctx, log, records)
	if err != nil {
		log.Errorf("failed to add records %v", err)
		return err
	}
	return nil
}
