package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/dto"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RepairHandler struct {
	service *service.RepairService
}

func NewRepairHandler(service *service.RepairService) *RepairHandler {
	return &RepairHandler{service: service}
}

func (h *RepairHandler) Create(c fiber.Ctx) error {
	var req dto.CreateRepairDTO

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid body",
			"error":   err.Error(),
		})
	}

	result, err := h.service.Create(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(result)
}

func (h *RepairHandler) GetAll(c fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}

func (h *RepairHandler) GetById(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	Repair, err := h.service.GetById(objID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Repair not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(Repair)
}

func (h *RepairHandler) UpdateById(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	var Repair dto.UpdateRepairDTO
	if err := c.Bind().Body(&Repair); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request Body",
		})
	}
	result, err := h.service.UpdateById(objID, Repair)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Repair not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *RepairHandler) DeleteById(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	result, err := h.service.DeleteById(objID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Repair not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *RepairHandler) GetRepairDetail(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	Repair, err := h.service.RepairDetail(objID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Repair not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(Repair)
}