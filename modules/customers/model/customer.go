package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Customer struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Phone        string        `bson:"phone" json:"phone"`
	VehiclePlate string        `bson:"vehiclePlate" json:"vehiclePlate"`
	VehicleBrand string        `bson:"vehicleBrand" json:"vehicleBrand"`
	VehicleModel string        `bson:"vehicleModel" json:"vehicleModel"`
	CreatedAt    time.Time     `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt,omitempty" json:"updatedAt"`
}
