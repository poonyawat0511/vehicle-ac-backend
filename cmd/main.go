package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/poonyawat/vehicle-ac-backend/config"
	customerRoutes "github.com/poonyawat/vehicle-ac-backend/modules/customers/routes"
	repairRoutes "github.com/poonyawat/vehicle-ac-backend/modules/repairs/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectMongo()
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} [${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Vehicle AC API running",
		})
	})

	api := app.Group("/api")
	customerRoutes.CusotmerSetup(api)
	repairRoutes.RepairSetup(api)

	log.Fatal(app.Listen(":8080"))
}
