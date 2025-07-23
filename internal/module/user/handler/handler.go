package handler

import (
	m "inventori-beacukai-backend/internal/middleware"
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
	router.Post("/change-password", m.AuthBearer, h.changePassword)
	router.Get("/listen", m.AuthBearer, h.getProfile)
	router.Post("/register", h.registerUser)
	router.Get("/get", m.AuthBearer, h.getUsers)
	router.Get("/roles", m.AuthBearer, h.getRole)
	router.Post("/update-profile", m.AuthBearer, h.updateProfile)
	router.Delete("/delete/:id", m.AuthBearer, h.deleteUser)
}
