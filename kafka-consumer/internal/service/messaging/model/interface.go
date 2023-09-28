package model

import (
	"kafka/golang-kafka/kafka-consumer/internal/log"

	"context"
)

type Repository interface {
	Get(ctx context.Context, log log.Logger) ([]*Message, error)
	Add(ctx context.Context, log log.Logger, records Message) error
}

type Service interface {
	Get(ctx context.Context, log log.Logger) ([]*Message, error)
	Add(ctx context.Context, log log.Logger, records Message) error
}
