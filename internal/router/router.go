package router

import (
	userHandler "inventori-beacukai-backend/internal/module/user/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userHandler.NewUserHandler().Register(app.Group("/users"))
}
