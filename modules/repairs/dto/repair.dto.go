package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateRepairDTO struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID bson.ObjectID `bson:"customerId,omitempty" json:"customerId"`
	Problem    string        `bson:"problem" json:"problem"`
	Solution   []string      `bson:"solution" json:"solution"`
	Price      int           `bson:"price" json:"price"`
}

type UpdateRepairDTO struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID bson.ObjectID `bson:"customerId,omitempty" json:"customerId"`
	Problem    string        `bson:"problem" json:"problem"`
	Solution   []string      `bson:"solution" json:"solution"`
	Price      int           `bson:"price" json:"price"`
}
