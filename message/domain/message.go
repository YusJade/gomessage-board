package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID       primitive.ObjectID `bson:"_id"`
	Content  string             `bson:"content"`
	Datetime string             `bson:"datetime"`
}
