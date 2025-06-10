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
	// supliers
	router.Get("/supliers", h.getSupliers)
	router.Get("/supliers/:id", h.getSuplier)
	router.Post("/supliers", h.createSuplier)
	router.Put("/supliers/:id", h.updateSuplier)
	router.Delete("/supliers/:id", h.deleteSuplier)

	// products
	router.Get("/products", h.getProducts)
	router.Get("/products/:id", h.getProduct)
	router.Post("/products", h.createProduct)
	router.Put("/products/:id", h.updateProduct)
	router.Delete("/products/:id", h.deleteProduct)

	// currencies
	router.Get("/currencies", h.getCurrencies)
	router.Get("/currencies/:id", h.getCurrency)
	router.Post("/currencies", h.createCurrency)
	router.Put("/currencies/:id", h.updateCurrency)
	router.Delete("/currencies/:id", h.deleteCurrency)

	// Saldo Awal
	router.Get("/saldo-awals", h.getSaldoAwals)
	router.Get("/saldo-awals/:id", h.getSaldoAwal)
	router.Post("/saldo-awals", h.createSaldoAwal)
	router.Put("/saldo-awals/:id", h.updateSaldoAwal)
	router.Delete("/saldo-awals/:id", h.deleteSaldoAwal)

	// bc_documents
	router.Get("/bc-documents", h.getBcDocuments)
	router.Get("/bc-documents/:id", h.getBcDocument)
	router.Post("/bc-documents", h.createBcDocument)
	router.Put("/bc-documents/:id", h.updateBcDocument)
	router.Delete("/bc-documents/:id", h.deleteBcDocument)

	// contracts
	router.Get("/contracts", h.getContracts)
	router.Get("/contracts/:id", h.getContract)
	router.Post("/contracts", h.createContract)
	router.Put("/contracts/:id", h.updateContract)
	router.Delete("/contracts/:id", h.deleteContract)

	// contract_products
	router.Get("/contract-products", h.getContractProducts)
	router.Get("/contract-products/:id", h.getContractProduct)
	router.Post("/contract-products", h.createContractProduct)
	router.Put("/contract-products/:id", h.updateContractProduct)
	router.Delete("/contract-products/:id", h.deleteContractProduct)

	// income_inventories
	router.Get("/income-inventories", h.getIncomeInventories)
	router.Get("/income-inventories/:id", h.getIncomeInventory)
	router.Post("/income-inventories", h.createIncomeInventory)
	router.Put("/income-inventories/:id", h.updateIncomeInventory)
	router.Delete("/income-inventories/:id", h.deleteIncomeInventory)

	// income_inventories_products
	router.Get("/income-inventories-products", h.getIncomeInventoriesProducts)
	router.Get("/income-inventories-products/:id", h.getIncomeInventoriesProduct)
	router.Post("/income-inventories-products", h.createIncomeInventoriesProduct)
	router.Put("/income-inventories-products/:id", h.updateIncomeInventoriesProduct)
	router.Delete("/income-inventories-products/:id", h.deleteIncomeInventoriesProduct)

	// transaction_incomes
	router.Get("/transaction-incomes", h.getTransactionIncomes)
	router.Get("/transaction-incomes/:id", h.getTransactionIncome)
	router.Post("/transaction-incomes", h.createTransactionIncome)
	router.Put("/transaction-incomes/:id", h.updateTransactionIncome)
	router.Delete("/transaction-incomes/:id", h.deleteTransactionIncome)

	// laporan
	router.Get("/laporan-mutasi", h.getLaporanMutasi)
}
