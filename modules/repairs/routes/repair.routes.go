package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/handler"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/repository"
	"github.com/poonyawat/vehicle-ac-backend/modules/repairs/service"
)

func RepairSetup(router fiber.Router) {
	repo := repository.NewRepairRepository()
	svc := service.NewRepairService(repo)
	h := handler.NewRepairHandler(svc)

	r := router.Group("/repairs")

	r.Post("/", h.Create)
	r.Get("/", h.GetAll)
	r.Get("/:id", h.GetById)
	r.Put("/:id", h.UpdateById)
	r.Delete("/:id", h.DeleteById)
}
