package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repair struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID bson.ObjectID `bson:"customerId,omitempty" json:"customerId"` 
	Problem    string        `bson:"problem" json:"problem"`
	Solution   []string      `bson:"solution" json:"solution"` 
	Price      int           `bson:"price" json:"price"`
	CreatedAt  time.Time     `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt  time.Time     `bson:"updatedAt,omitempty" json:"updatedAt"`
}