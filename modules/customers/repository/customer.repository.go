package repository

import (
	"context"

	"github.com/poonyawat/vehicle-ac-backend/config"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CustomerRepository struct {
	col *mongo.Collection
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		col: config.DB.Collection("customers"),
	}
}

func (r *CustomerRepository) Create(customer model.Customer) error {
	_, err := r.col.InsertOne(context.Background(), customer)
	return err
}

func (r *CustomerRepository) FindAll() ([]model.Customer, error) {
	cursor, err := r.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var result []model.Customer
	err = cursor.All(context.Background(), &result)
	return result, err
}

func (r *CustomerRepository) FindById(id bson.ObjectID) (model.Customer, error) {
	var result model.Customer
	err := r.col.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return model.Customer{}, err
	}

	return result, nil
}

func (r *CustomerRepository) UpdateById(id bson.ObjectID, updateData model.Customer) (model.Customer, error) {
	var result model.Customer
	update := bson.M{
		"$set": updateData,
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.col.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update, opts).Decode(&result)
	if err != nil {
		return model.Customer{}, err
	}
	return result, nil
}

func (r *CustomerRepository) DeleteById(id bson.ObjectID) (model.Customer, error) {
	var result model.Customer
	err := r.col.FindOneAndDelete(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return model.Customer{}, err
	}
	return result, nil
}
