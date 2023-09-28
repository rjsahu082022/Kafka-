package consumer_test

import (
	"context"
	"errors"
	"testing"

	consumer "kafka/golang-kafka/kafka-consumer/internal/consumer"
	"kafka/golang-kafka/kafka-consumer/internal/log"
	logMock "kafka/golang-kafka/kafka-consumer/internal/log/mocks"
	messagingModel "kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"
	messageMocks "kafka/golang-kafka/kafka-consumer/internal/service/messaging/model/mocks"

	"github.com/stretchr/testify/mock"
)

var (
	bootstrapServer = "localhost:9092"
	errService      = errors.New("messages service error")
)

func TestConsumer_Consume(t *testing.T) {
	t.Parallel()
	type fields struct {
		log             log.Logger
		messageService  messagingModel.Service
		bootstrapServer []string
		topic           []string
	}
	type args struct {
		ctx    context.Context
		topics []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "consumer stopped",
			fields: fields{
				log: func() log.Logger {
					logger := new(logMock.Logger)
					logger.On("Infof", mock.Anything, mock.Anything)
					logger.On("Info", mock.Anything, mock.Anything)
					logger.On("Errorf", mock.Anything, mock.Anything)
					return logger
				}(),
				messageService: func() messagingModel.Service {
					service := new(messageMocks.Service)
					service.On("Add", mock.Anything, mock.Anything, mock.Anything).Return(nil)
					return service
				}(),
				bootstrapServer: []string{bootstrapServer},
			},
			args: args{
				ctx: context.Background(),
				topics: func() []string {
					topics := []string{"testTopics"}
					return topics
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, err := consumer.NewConsumer(tt.fields.log, tt.fields.messageService, tt.fields.bootstrapServer, tt.args.topics)
			if err != nil {
				t.Errorf("Failed to initialize consumer: %v", err)
			}
			if err := c.Consume(tt.args.ctx, tt.args.topics); (err != nil) != tt.wantErr {
				t.Errorf("Consume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
