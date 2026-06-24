package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/handler"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/repository"
	"github.com/poonyawat/vehicle-ac-backend/modules/customers/service"
)

func CusotmerSetup(router fiber.Router) {
	repo := repository.NewCustomerRepository()
	svc := service.NewCustomerService(repo)
	h := handler.NewCustomerHandler(svc)

	r := router.Group("/customers")

	r.Post("/", h.Create)
	r.Get("/", h.GetAll)
	r.Get("/:id", h.GetById)
	r.Put("/:id", h.UpdateById)
	r.Delete("/:id", h.DeleteById)
}
