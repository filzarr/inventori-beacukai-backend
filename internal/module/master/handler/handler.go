package handler

import (
	"inventori-beacukai-backend/internal/module/master/ports"
	"inventori-beacukai-backend/internal/module/master/repository"
	"inventori-beacukai-backend/internal/module/master/service"

	"github.com/gofiber/fiber/v2"
)

type MasterHandler struct {
	service ports.MasterService
}

func NewMasterHandler() *MasterHandler {
	var (
		repo    = repository.NewMasterRepository()
		svc     = service.NewMasterService(repo)
		handler = new(MasterHandler)
	)
	handler.service = svc
	return handler
}

func (h *MasterHandler) Register(router fiber.Router) {
	// User Router
	router.Get("/users", h.getUsers)
	// Inventory Router
	router.Get("/inventory", h.getInventories)
	router.Get("/inventory/:id", h.getInventory)
	router.Post("/inventory", h.createInventory)
	router.Put("/inventory/:id", h.updateInventory)
	router.Delete("/inventory/:id", h.deleteInventory)
}
