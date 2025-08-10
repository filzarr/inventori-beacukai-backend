package handler

import (
	m "inventori-beacukai-backend/internal/middleware"
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
	router.Post("/contracts/document", h.updateContractDocument)
	router.Put("/contracts/:id", h.updateContract)
	router.Delete("/contracts/:id", h.deleteContract)
	router.Get("/contracts-transactions", h.getTransactions)
	router.Get("/contracts-not-required", h.getContractNotRequired)

	// warehouses
	router.Get("/warehouses", m.AuthBearer, h.getWarehouses)
	router.Get("/warehouses/:id", m.AuthBearer, h.getWarehouse)
	router.Post("/warehouses", m.AuthBearer, h.createWarehouse)
	router.Put("/warehouses/:id", m.AuthBearer, h.updateWarehouse)
	router.Delete("/warehouses/:id", m.AuthBearer, h.deleteWarehouse)
	// warehouses-stocks
	router.Get("/warehouses-stocks", m.AuthBearer, h.getWarehousesStocks)
	router.Get("/warehouses-stocks/:id", m.AuthBearer, h.getWarehousesStock)
	router.Post("/warehouses-stocks", m.AuthBearer, h.createWarehousesStock)
	router.Put("/warehouses-stocks/:id", m.AuthBearer, h.updateWarehousesStock)
	router.Delete("/warehouses-stocks/:id", m.AuthBearer, h.deleteWarehousesStock)
	// contracts-bc
	router.Get("/contracts-bc", m.AuthBearer, h.getContractsBc)
	router.Get("/contracts-bc/:id", m.AuthBearer, h.getContractBc)
	router.Post("/contracts-bc", m.AuthBearer, h.createContractBc)
	router.Put("/contracts-bc/:id", m.AuthBearer, h.updateContractBc)
	router.Delete("/contracts-bc/:id", m.AuthBearer, h.deleteContractBc)

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
	router.Get("/income-inventories-by-contract", h.getIncomeInventoriesProductsByContract)
	router.Get("/income-inventories-by-contract-and-kode", h.getIncomeInventoriesProductsByContractAndKode)

	// income_inventories_products
	router.Get("/income-inventories-products", h.getIncomeInventoriesProducts)
	router.Get("/income-inventories-products/:id", h.getIncomeInventoriesProduct)
	router.Post("/income-inventories-products", h.createIncomeInventoriesProduct)
	router.Put("/income-inventories-products/:id", h.updateIncomeInventoriesProduct)
	router.Delete("/income-inventories-products/:id", h.deleteIncomeInventoriesProduct)

	// income_inventories_products
	router.Get("/outcomes-inventories-products", h.getOutcomesInventoriesProducts)
	router.Get("/outcomes-inventories-products/:id", h.getOutcomesInventoriesProduct)
	router.Post("/outcomes-inventories-products", h.createOutcomesInventoriesProduct)
	router.Put("/outcomes-inventories-products/:id", h.updateOutcomesInventoriesProduct)
	router.Delete("/outcomes-inventories-products/:id", h.deleteOutcomesInventoriesProduct)

	// transaction_incomes
	router.Get("/transaction-incomes", h.getTransactionIncomes)
	router.Get("/transaction-incomes/:id", h.getTransactionIncome)
	router.Post("/transaction-incomes", h.createTransactionIncome)
	router.Put("/transaction-incomes/:id", h.updateTransactionIncome)
	router.Delete("/transaction-incomes/:id", h.deleteTransactionIncome)

	// ready_products
	router.Get("/ready-products", h.getReadyProducts)
	router.Get("/ready-products/:id", h.getReadyProduct)
	router.Post("/ready-products", h.createReadyProduct)
	router.Put("/ready-products/:id", h.updateReadyProduct)
	router.Delete("/ready-products/:id", h.deleteReadyProduct)

	// Buyers
	router.Get("/buyers", h.getBuyers)
	router.Get("/buyers/:id", h.getBuyer)
	router.Post("/buyers", h.createBuyer)
	router.Put("/buyers/:id", h.updateBuyer)
	router.Delete("/buyers/:id", h.deleteBuyer)

	// transfers_products
	router.Get("/transfers-products", h.getTransfersProducts)
	router.Get("/transfers-products/:id", h.getTransferProduct)
	router.Post("/transfers-products", h.createTransferProduct)
	router.Put("/transfers-products/:id", h.updateTransferProduct)
	router.Delete("/transfers-products/:id", h.deleteTransferProduct)

	// laporan
	router.Get("/laporan-mutasi", h.getLaporanMutasi)
	router.Get("/laporan-mutasi/pemasukan", h.getLaporanMutasiPemasukan)

	// productions
	router.Get("/productions", h.getProductions)
	router.Get("/productions/:id", h.getProduction)
	router.Post("/productions", h.createProduction)
	router.Put("/productions/:id", h.updateProduction)
	router.Delete("/productions/:id", h.deleteProduction)

	// products_movement
	router.Get("/products_movement", h.getProductsMovement)
	router.Get("/products_movement/:id", h.getProductsMovementByID)
	router.Post("/products_movement", h.createProductsMovement)
	router.Post("/products-movement-status", h.updateStatusProductsMovement)
	router.Put("/products_movement/:id", h.updateProductsMovement)
	router.Delete("/products_movement/:id", h.deleteProductsMovement)
}
