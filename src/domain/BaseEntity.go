package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseEntity struct {
	Id         primitive.ObjectID `bson:"_id"`
	InsertedAt time.Time          `bson:"inserted_at"`
	UpdatedAt  time.Time          `bson:"last_update"`
	CreatedBy  primitive.ObjectID `bson:"created_by"`
}


func NewBaseEntity(createdBy primitive.ObjectID) *BaseEntity{
	return &BaseEntity{
		Id:         primitive.NewObjectID(),
		InsertedAt: time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		CreatedBy:  createdBy,
	}
}