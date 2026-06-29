package repository

import (
	"context"

	"github.com/poonyawat/vehicle-ac-backend/config"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/dto"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RepairRepository struct {
	col *mongo.Collection
}

func NewRepairRepository() *RepairRepository {
	return &RepairRepository{
		col: config.DB.Collection("repairs"),
	}
}

func (r *RepairRepository) Create(repair model.Repair) error {
	_, err := r.col.InsertOne(context.Background(), repair)
	return err
}

func (r *RepairRepository) FindAll() ([]model.Repair, error) {
	cursor, err := r.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var result []model.Repair
	err = cursor.All(context.Background(), &result)
	return result, err
}

func (r *RepairRepository) FindById(id bson.ObjectID) (model.Repair, error) {
	var result model.Repair
	err := r.col.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return model.Repair{}, err
	}

	return result, nil
}

func (r *RepairRepository) UpdateById(id bson.ObjectID, updateData model.Repair) (model.Repair, error) {
	var result model.Repair
	update := bson.M{
		"$set": updateData,
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.col.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update, opts).Decode(&result)
	if err != nil {
		return model.Repair{}, err
	}
	return result, nil
}

func (r *RepairRepository) DeleteById(id bson.ObjectID) (model.Repair, error) {
	var result model.Repair
	err := r.col.FindOneAndDelete(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return model.Repair{}, err
	}
	return result, nil
}

func (r *RepairRepository) RepairDetail(id bson.ObjectID) (dto.RepairDetail, error) {
	var result dto.RepairDetail

	pipeline := mongo.Pipeline{
		{
			{
				Key: "$match",
				Value: bson.M{
					"_id": id,
				},
			},
		},
		{
			{
				Key: "$lookup",
				Value: bson.M{
					"from":         "customers",
					"localField":   "customerId",
					"foreignField": "_id",
					"as":           "customer",
				},
			},
		},
		{
			{
				Key: "$unwind",
				Value: "$customer",
			},
		},
	}

	cursor, err := r.col.Aggregate(context.Background(), pipeline)
	if err != nil {
		return result, err
	}

	if cursor.Next(context.Background()) {
		err := cursor.Decode(&result)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}