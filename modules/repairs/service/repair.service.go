package service

import (
	"time"

	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/dto"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/model"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RepairService struct {
	repo *repository.RepairRepository
}

func NewRepairService(repo *repository.RepairRepository) *RepairService {
	return &RepairService{repo: repo}
}

func (s *RepairService) Create(req dto.CreateRepairDTO) (*model.Repair, error) {
	Repair := model.Repair{
		ID:         bson.NewObjectID(),
		CustomerID: req.CustomerID,
		Problem:    req.Problem,
		Solution:   req.Solution,
		Price:      req.Price,
		CreatedAt:  time.Now(),
	}

	err := s.repo.Create(Repair)
	if err != nil {
		return nil, err
	}

	return &Repair, nil
}

func (s *RepairService) GetAll() ([]model.Repair, error) {
	return s.repo.FindAll()
}

func (s *RepairService) GetById(id bson.ObjectID) (model.Repair, error) {
	Repair, err := s.repo.FindById(id)
	if err != nil {
		return model.Repair{}, err
	}
	return Repair, nil
}

func (s *RepairService) UpdateById(id bson.ObjectID, req dto.UpdateRepairDTO) (model.Repair, error) {
	Repair := model.Repair{
		CustomerID: req.CustomerID,
		Problem:    req.Problem,
		Solution:   req.Solution,
		Price:      req.Price,
		UpdatedAt:  time.Now(),
	}
	result, err := s.repo.UpdateById(id, Repair)
	if err != nil {
		return model.Repair{}, err
	}
	return result, nil
}

func (s *RepairService) DeleteById(id bson.ObjectID) (model.Repair, error) {
	result, err := s.repo.DeleteById(id)
	if err != nil {
		return model.Repair{}, err
	}

	return result, nil
}
