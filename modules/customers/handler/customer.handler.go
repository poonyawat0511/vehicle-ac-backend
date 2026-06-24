package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/dto"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(service *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) Create(c fiber.Ctx) error {
	var req dto.CreateCustomerDTO

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid body",
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

func (h *CustomerHandler) GetAll(c fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(data)
}

func (h *CustomerHandler) GetById(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	customer, err := h.service.GetById(objID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(customer)
}

func (h *CustomerHandler) UpdateById(c fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Format",
		})
	}
	var customer dto.UpdateCustomerDTO
	if err := c.Bind().Body(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request Body",
		})
	}
	result, err := h.service.UpdateById(objID, customer)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *CustomerHandler) DeleteById(c fiber.Ctx) error {
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
			"error": "Customer not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
