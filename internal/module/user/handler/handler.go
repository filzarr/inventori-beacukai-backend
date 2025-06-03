package handler

import (
	"inventori-beacukai-backend/internal/module/user/ports"
	"inventori-beacukai-backend/internal/module/user/repository"
	"inventori-beacukai-backend/internal/module/user/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler() *UserHandler {
	var (
		repo    = repository.NewUserRepository()
		svc     = service.NewUserService(repo)
		handler = new(UserHandler)
	)
	handler.service = svc

	return handler
}

func (h *UserHandler) Register(router fiber.Router) {
	router.Post("/login", h.login)
	router.Post("/register", h.registerUser)
}
