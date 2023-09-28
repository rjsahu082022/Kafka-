package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	From    string             `json:"from,omitempty" bson:"from"`
	Message string             `json:"message,omitempty" bson:"message"`
	TO      string             `json:"to,omitempty" bson:"to"`
}
