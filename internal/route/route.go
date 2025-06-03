package route

import (
	masterHandler "inventori-beacukai-backend/internal/module/master/handler"
	userHandler "inventori-beacukai-backend/internal/module/user/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userHandler.NewUserHandler().Register(app.Group("/users"))
	masterHandler.NewMasterHandler().Register(app.Group("/api/v1"))
}
