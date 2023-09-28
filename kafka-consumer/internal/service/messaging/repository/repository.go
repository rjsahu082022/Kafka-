package messaging

import (
	"kafka/golang-kafka/kafka-consumer/internal/log"
	"kafka/golang-kafka/kafka-consumer/internal/service/messaging/model"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context, log log.Logger) ([]*model.Message, error) {
	cur, err := r.db.Database().Collection("Messages").Find(ctx, bson.D{{}}, nil)
	if err != nil {
		log.Errorf("failed to add records into db %v", err)
		return nil, err
	}
	var messages []*model.Message
	if err := cur.All(context.Background(), &messages); err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	return messages, nil
}

func (r *Repository) Add(ctx context.Context, log log.Logger, records model.Message) error {
	_, err := r.db.InsertOne(ctx, &records)
	if err != nil {
		log.Errorf("failed to add records into db %v", err)
		return err
	}
	return nil
}
