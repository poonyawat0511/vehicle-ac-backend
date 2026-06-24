package service

import (
	"time"

	"github.com/poonyawat/vehicle-ac-backend/modules/customers/dto"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/model"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(req dto.CreateCustomerDTO) (*model.Customer, error) {
	customer := model.Customer{
		ID:           bson.NewObjectID(),
		Name:         req.Name,
		Phone:        req.Phone,
		VehiclePlate: req.VehiclePlate,
		VehicleBrand: req.VehicleBrand,
		VehicleModel: req.VehicleModel,
		CreatedAt:    time.Now(),
	}

	err := s.repo.Create(customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *CustomerService) GetAll() ([]model.Customer, error) {
	return s.repo.FindAll()
}

func (s *CustomerService) GetById(id bson.ObjectID) (model.Customer, error) {
	customer, err := s.repo.FindById(id)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (s *CustomerService) UpdateById(id bson.ObjectID, req dto.UpdateCustomerDTO) (model.Customer, error) {
	customer := model.Customer{
		Name:         req.Name,
		Phone:        req.Phone,
		VehiclePlate: req.VehiclePlate,
		VehicleBrand: req.VehicleBrand,
		VehicleModel: req.VehicleModel,
		UpdatedAt:    time.Now(),
	}
	result, err := s.repo.UpdateById(id, customer)
	if err != nil {
		return model.Customer{}, err
	}
	return result, nil
}

func (s *CustomerService) DeleteById(id bson.ObjectID) (model.Customer, error) {
	result, err := s.repo.DeleteById(id)
	if err != nil {
		return model.Customer{}, err
	}

	return result, nil
}
