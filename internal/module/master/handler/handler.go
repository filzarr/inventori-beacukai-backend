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
	router.Get("/inventory/bahanbaku", h.getInventoriesBahanBaku)
	router.Get("/inventory/:id", h.getInventory)
	router.Post("/inventory", h.createInventory)
	router.Put("/inventory/:id", h.updateInventory)
	router.Delete("/inventory/:id", h.deleteInventory)
	// Gudang Routes
	router.Get("/gudang", h.getGudangs)
	router.Get("/gudang/:id", h.getGudang)
	router.Post("/gudang", h.createGudang)
	router.Put("/gudang/:id", h.updateGudang)
	router.Delete("/gudang/:id", h.deleteGudang)
	// Mutasi bahan routes
	router.Get("/mutasi-bahan", h.getMutasiBahans)
	router.Get("/mutasi-bahan/:id", h.getMutasiBahan)
	router.Post("/mutasi-bahan", h.createMutasiBahan)
	router.Put("/mutasi-bahan/:id", h.updateMutasiBahan)
	router.Put("/mutasi-bahan/:id/saldo", h.updateSaldoMutasi)
	router.Delete("/mutasi-bahan/:id", h.deleteMutasiBahan)
	router.Post("/mutasi-bahan/update", h.updateStatusMutasiBahan)
	router.Get("/laporan/mutasi-bahan", h.getLaporanMutasiBahan)
}
