package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

type MasterRepository interface {
	// Users
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)

	// Supliers
	GetSupliers(ctx context.Context, req *entity.GetSupliersReq) (*entity.GetSupliersResp, error)
	GetSuplier(ctx context.Context, req *entity.GetSuplierReq) (*entity.GetSuplierResp, error)
	CreateSuplier(ctx context.Context, req *entity.CreateSuplierReq) (*entity.CreateSuplierResp, error)
	UpdateSuplier(ctx context.Context, req *entity.UpdateSuplierReq) error
	DeleteSuplier(ctx context.Context, req *entity.DeleteSuplierReq) error

	// Products
	GetProducts(ctx context.Context, req *entity.GetProductsReq) (*entity.GetProductsResp, error)
	GetProduct(ctx context.Context, req *entity.GetProductReq) (*entity.GetProductResp, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductReq) (*entity.CreateProductResp, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductReq) error
	DeleteProduct(ctx context.Context, req *entity.DeleteProductReq) error

	// Currencies
	GetCurrencies(ctx context.Context, req *entity.GetCurrenciesReq) (*entity.GetCurrenciesResp, error)
	GetCurrency(ctx context.Context, req *entity.GetCurrencyReq) (*entity.GetCurrencyResp, error)
	CreateCurrency(ctx context.Context, req *entity.CreateCurrencyReq) (*entity.CreateCurrencyResp, error)
	UpdateCurrency(ctx context.Context, req *entity.UpdateCurrencyReq) error
	DeleteCurrency(ctx context.Context, req *entity.DeleteCurrencyReq) error

	// BCDocuments
	GetBcDocuments(ctx context.Context, req *entity.GetBcDocumentsReq) (*entity.GetBcDocumentsResp, error)
	GetBcDocument(ctx context.Context, req *entity.GetBcDocumentReq) (*entity.GetBcDocumentResp, error)
	CreateBcDocument(ctx context.Context, req *entity.CreateBcDocumentReq) (*entity.CreateBcDocumentResp, error)
	UpdateBcDocument(ctx context.Context, req *entity.UpdateBcDocumentReq) error
	DeleteBcDocument(ctx context.Context, req *entity.DeleteBcDocumentReq) error

	// Contracts
	GetContracts(ctx context.Context, req *entity.GetContractsReq) (*entity.GetContractsResp, error)
	GetContract(ctx context.Context, req *entity.GetContractReq) (*entity.GetContractResp, error)
	CreateContract(ctx context.Context, req *entity.CreateContractReq) (*entity.CreateContractResp, error)
	UpdateContract(ctx context.Context, req *entity.UpdateContractReq) error
	DeleteContract(ctx context.Context, req *entity.DeleteContractReq) error
	UpdateContractDocument(ctx context.Context, req *entity.UpdateContractDocumentReq) error
	GetTransactions(ctx context.Context, req *entity.GetTransactionsReq) (*entity.GetTransactionsResp, error)
	GetContractNotRequired(ctx context.Context, req *entity.GetContractNotRequiredReq) (*entity.GetContractNotRequiredResp, error)

	// Warehouses
	GetWarehouses(ctx context.Context, req *entity.GetWarehousesReq) (*entity.GetWarehousesResp, error)
	GetWarehouse(ctx context.Context, req *entity.GetWarehouseReq) (*entity.GetWarehouseResp, error)
	CreateWarehouse(ctx context.Context, req *entity.CreateWarehouseReq) (*entity.CreateWarehouseResp, error)
	UpdateWarehouse(ctx context.Context, req *entity.UpdateWarehouseReq) error
	DeleteWarehouse(ctx context.Context, req *entity.DeleteWarehouseReq) error
	// WarehousesStocks
	GetWarehousesStocks(ctx context.Context, req *entity.GetWarehousesStocksReq) (*entity.GetWarehousesStocksResp, error)
	GetWarehousesStock(ctx context.Context, req *entity.GetWarehousesStockReq) (*entity.GetWarehousesStockResp, error)
	CreateWarehousesStock(ctx context.Context, req *entity.CreateWarehousesStockReq) (*entity.CreateWarehousesStockResp, error)
	UpdateWarehousesStock(ctx context.Context, req *entity.UpdateWarehousesStockReq) error
	UpdateStockWarehouses(ctx context.Context, req *entity.UpdateStockWarehousesReq) error
	DeleteWarehousesStock(ctx context.Context, req *entity.DeleteWarehousesStockReq) error
	// ContractsBc
	GetContractsBc(ctx context.Context, req *entity.GetContractsBcReq) (*entity.GetContractsBcResp, error)
	GetContractBc(ctx context.Context, req *entity.GetContractBcReq) (*entity.GetContractBcResp, error)
	CreateContractBc(ctx context.Context, req *entity.CreateContractBcReq) (*entity.CreateContractBcResp, error)
	UpdateContractBc(ctx context.Context, req *entity.UpdateContractBcReq) error
	DeleteContractBc(ctx context.Context, req *entity.DeleteContractBcReq) error

	// SaldoAwal
	GetSaldoAwals(ctx context.Context, req *entity.GetSaldoAwalsReq) (*entity.GetSaldoAwalsResp, error)
	GetSaldoAwal(ctx context.Context, req *entity.GetSaldoAwalReq) (*entity.GetSaldoAwalResp, error)
	CreateSaldoAwal(ctx context.Context, req *entity.CreateSaldoAwalReq) (*entity.CreateSaldoAwalResp, error)
	UpdateSaldoAwal(ctx context.Context, req *entity.UpdateSaldoAwalReq) error
	DeleteSaldoAwal(ctx context.Context, req *entity.DeleteSaldoAwalReq) error

	// ContractProducts
	GetContractProducts(ctx context.Context, req *entity.GetContractProductsReq) (*entity.GetContractProductsResp, error)
	GetContractProduct(ctx context.Context, req *entity.GetContractProductReq) (*entity.GetContractProductResp, error)
	CreateContractProduct(ctx context.Context, req *entity.CreateContractProductReq) (*entity.CreateContractProductResp, error)
	UpdateContractProduct(ctx context.Context, req *entity.UpdateContractProductReq) error
	DeleteContractProduct(ctx context.Context, req *entity.DeleteContractProductReq) error

	// IncomeInventories
	GetIncomeInventories(ctx context.Context, req *entity.GetIncomeInventoriesReq) (*entity.GetIncomeInventoriesResp, error)
	GetIncomeInventory(ctx context.Context, req *entity.GetIncomeInventoryReq) (*entity.GetIncomeInventoryResp, error)
	CreateIncomeInventory(ctx context.Context, req *entity.CreateIncomeInventoryReq) (*entity.CreateIncomeInventoryResp, error)
	UpdateIncomeInventory(ctx context.Context, req *entity.UpdateIncomeInventoryReq) error
	DeleteIncomeInventory(ctx context.Context, req *entity.DeleteIncomeInventoryReq) error

	// IncomeInventoryProducts
	GetIncomeInventoryProducts(ctx context.Context, req *entity.GetIncomeInventoryProductsReq) (*entity.GetIncomeInventoryProductsResp, error)
	GetIncomeInventoryProduct(ctx context.Context, req *entity.GetIncomeInventoryProductReq) (*entity.GetIncomeInventoryProductResp, error)
	CreateIncomeInventoryProduct(ctx context.Context, req *entity.CreateIncomeInventoryProductReq) (*entity.CreateIncomeInventoryProductResp, error)
	UpdateIncomeInventoryProduct(ctx context.Context, req *entity.UpdateIncomeInventoryProductReq) error
	DeleteIncomeInventoryProduct(ctx context.Context, req *entity.DeleteIncomeInventoryProductReq) error
	GetIncomeInventoryProductsByContract(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractReq) (*entity.GetIncomeInventoryProductsByContractResp, error)
	GetIncomeInventoryProductsByContractAndKode(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractAndKodeReq) (*entity.GetIncomeInventoryProductsByContractAndKodeResp, error)

	// OutcomeInventoriesProducts
	GetOutcomesInventoriesProducts(ctx context.Context, req *entity.GetOutcomesInventoriesProductsReq) (*entity.GetOutcomesInventoriesProductsResp, error)
	GetOutcomesInventoriesProduct(ctx context.Context, req *entity.GetOutcomesInventoriesProductReq) (*entity.GetOutcomesInventoriesProductResp, error)
	CreateOutcomesInventoriesProduct(ctx context.Context, req *entity.CreateOutcomesInventoriesProductReq) (*entity.CreateOutcomesInventoriesProductResp, error)
	UpdateOutcomesInventoriesProduct(ctx context.Context, req *entity.UpdateOutcomesInventoriesProductReq) error
	DeleteOutcomesInventoriesProduct(ctx context.Context, req *entity.DeleteOutcomesInventoriesProductReq) error

	// TransactionIncomes
	GetTransactionIncomes(ctx context.Context, req *entity.GetTransactionIncomesReq) (*entity.GetTransactionIncomesResp, error)
	GetTransactionIncome(ctx context.Context, req *entity.GetTransactionIncomeReq) (*entity.GetTransactionIncomeResp, error)
	CreateTransactionIncome(ctx context.Context, req *entity.CreateTransactionIncomeReq) (*entity.CreateTransactionIncomeResp, error)
	UpdateTransactionIncome(ctx context.Context, req *entity.UpdateTransactionIncomeReq) error
	DeleteTransactionIncome(ctx context.Context, req *entity.DeleteTransactionIncomeReq) error

	// ReadyProducts
	GetReadyProducts(ctx context.Context, req *entity.GetReadyProductsReq) (*entity.GetReadyProductsResp, error)
	GetReadyProduct(ctx context.Context, req *entity.GetReadyProductReq) (*entity.GetReadyProductResp, error)
	CreateReadyProduct(ctx context.Context, req *entity.CreateReadyProductReq) (*entity.CreateReadyProductResp, error)
	UpdateReadyProduct(ctx context.Context, req *entity.UpdateReadyProductReq) error
	DeleteReadyProduct(ctx context.Context, req *entity.DeleteReadyProductReq) error

	// Buyers
	GetBuyers(ctx context.Context, req *entity.GetBuyersReq) (*entity.GetBuyersResp, error)
	GetBuyer(ctx context.Context, req *entity.GetBuyerReq) (*entity.GetBuyerResp, error)
	CreateBuyer(ctx context.Context, req *entity.CreateBuyerReq) (*entity.CreateBuyerResp, error)
	UpdateBuyer(ctx context.Context, req *entity.UpdateBuyerReq) error
	DeleteBuyer(ctx context.Context, req *entity.DeleteBuyerReq) error

	// TransfersProducts
	GetTransfersProducts(ctx context.Context, req *entity.GetTransfersProductsReq) (*entity.GetTransfersProductsResp, error)
	GetTransferProduct(ctx context.Context, req *entity.GetTransferProductReq) (*entity.GetTransferProductResp, error)
	CreateTransferProduct(ctx context.Context, req *entity.CreateTransferProductReq) (*entity.CreateTransferProductResp, error)
	UpdateTransferProduct(ctx context.Context, req *entity.UpdateTransferProductReq) error
	DeleteTransferProduct(ctx context.Context, req *entity.DeleteTransferProductReq) error

	// Laporan
	GetLaporanMutasi(ctx context.Context, req *entity.GetLaporanMutasiReq) (*entity.GetLaporanMutasiResp, error)
	GetLaporanMutasiPemasukan(ctx context.Context, req *entity.GetLaporanMutasiPemasukanReq) (*entity.GetLaporanMutasiPemasukanResp, error)
	GetLaporanMutasiWip(ctx context.Context, req *entity.GetLaporanMutasiWipReq) (*entity.GetLaporanMutasiWipResp, error)
	GetLaporanMutasiJenisDokumen(ctx context.Context, req *entity.GetLaporanMutasiJenisDokumenReq) (*entity.GetLaporanMutasiJenisDokumenResp, error)

	// Productions
	GetProductions(ctx context.Context, req *entity.GetProductionsReq) (*entity.GetProductionsResp, error)
	GetProduction(ctx context.Context, req *entity.GetProductionReq) (*entity.GetProductionResp, error)
	CreateProduction(ctx context.Context, req *entity.CreateProductionReq) (*entity.CreateProductionResp, error)
	UpdateProduction(ctx context.Context, req *entity.UpdateProductionReq) error
	DeleteProduction(ctx context.Context, req *entity.DeleteProductionReq) error

	// ProductsMovement
	GetProductsMovement(ctx context.Context, req *entity.GetProductsMovementReq) (*entity.GetProductsMovementResp, error)
	GetProductsMovementByID(ctx context.Context, req *entity.GetProductsMovementReqID) (*entity.GetProductsMovementRespID, error)
	CreateProductsMovement(ctx context.Context, req *entity.CreateProductsMovementReq) (*entity.CreateProductsMovementResp, error)
	UpdateProductsMovement(ctx context.Context, req *entity.UpdateProductsMovementReq) error
	UpdateStatusProductsMovement(ctx context.Context, req *entity.UpdateStatusProductsMoveMentReq) error
	DeleteProductsMovement(ctx context.Context, req *entity.DeleteProductsMovementReq) error
}

type MasterService interface {
	// Users
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)

	// Supliers
	GetSupliers(ctx context.Context, req *entity.GetSupliersReq) (*entity.GetSupliersResp, error)
	GetSuplier(ctx context.Context, req *entity.GetSuplierReq) (*entity.GetSuplierResp, error)
	CreateSuplier(ctx context.Context, req *entity.CreateSuplierReq) (*entity.CreateSuplierResp, error)
	UpdateSuplier(ctx context.Context, req *entity.UpdateSuplierReq) error
	DeleteSuplier(ctx context.Context, req *entity.DeleteSuplierReq) error

	// Products
	GetProducts(ctx context.Context, req *entity.GetProductsReq) (*entity.GetProductsResp, error)
	GetProduct(ctx context.Context, req *entity.GetProductReq) (*entity.GetProductResp, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductReq) (*entity.CreateProductResp, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductReq) error
	DeleteProduct(ctx context.Context, req *entity.DeleteProductReq) error

	// Currencies
	GetCurrencies(ctx context.Context, req *entity.GetCurrenciesReq) (*entity.GetCurrenciesResp, error)
	GetCurrency(ctx context.Context, req *entity.GetCurrencyReq) (*entity.GetCurrencyResp, error)
	CreateCurrency(ctx context.Context, req *entity.CreateCurrencyReq) (*entity.CreateCurrencyResp, error)
	UpdateCurrency(ctx context.Context, req *entity.UpdateCurrencyReq) error
	DeleteCurrency(ctx context.Context, req *entity.DeleteCurrencyReq) error

	// BCDocuments
	GetBcDocuments(ctx context.Context, req *entity.GetBcDocumentsReq) (*entity.GetBcDocumentsResp, error)
	GetBcDocument(ctx context.Context, req *entity.GetBcDocumentReq) (*entity.GetBcDocumentResp, error)
	CreateBcDocument(ctx context.Context, req *entity.CreateBcDocumentReq) (*entity.CreateBcDocumentResp, error)
	UpdateBcDocument(ctx context.Context, req *entity.UpdateBcDocumentReq) error
	DeleteBcDocument(ctx context.Context, req *entity.DeleteBcDocumentReq) error

	// Contracts
	GetContracts(ctx context.Context, req *entity.GetContractsReq) (*entity.GetContractsResp, error)
	GetContract(ctx context.Context, req *entity.GetContractReq) (*entity.GetContractResp, error)
	CreateContract(ctx context.Context, req *entity.CreateContractReq) (*entity.CreateContractResp, error)
	UpdateContract(ctx context.Context, req *entity.UpdateContractReq) error
	DeleteContract(ctx context.Context, req *entity.DeleteContractReq) error
	UpdateContractDocument(ctx context.Context, req *entity.UpdateContractDocumentReq) error
	GetTransactions(ctx context.Context, req *entity.GetTransactionsReq) (*entity.GetTransactionsResp, error)
	GetContractNotRequired(ctx context.Context, req *entity.GetContractNotRequiredReq) (*entity.GetContractNotRequiredResp, error)

	// Warehouses
	GetWarehouses(ctx context.Context, req *entity.GetWarehousesReq) (*entity.GetWarehousesResp, error)
	GetWarehouse(ctx context.Context, req *entity.GetWarehouseReq) (*entity.GetWarehouseResp, error)
	CreateWarehouse(ctx context.Context, req *entity.CreateWarehouseReq) (*entity.CreateWarehouseResp, error)
	UpdateWarehouse(ctx context.Context, req *entity.UpdateWarehouseReq) error
	DeleteWarehouse(ctx context.Context, req *entity.DeleteWarehouseReq) error

	// WarehousesStocks
	GetWarehousesStocks(ctx context.Context, req *entity.GetWarehousesStocksReq) (*entity.GetWarehousesStocksResp, error)
	GetWarehousesStock(ctx context.Context, req *entity.GetWarehousesStockReq) (*entity.GetWarehousesStockResp, error)
	CreateWarehousesStock(ctx context.Context, req *entity.CreateWarehousesStockReq) (*entity.CreateWarehousesStockResp, error)
	UpdateWarehousesStock(ctx context.Context, req *entity.UpdateWarehousesStockReq) error
	UpdateStockWarehouses(ctx context.Context, req *entity.UpdateStockWarehousesReq) error
	DeleteWarehousesStock(ctx context.Context, req *entity.DeleteWarehousesStockReq) error

	// ContractsBc
	GetContractsBc(ctx context.Context, req *entity.GetContractsBcReq) (*entity.GetContractsBcResp, error)
	GetContractBc(ctx context.Context, req *entity.GetContractBcReq) (*entity.GetContractBcResp, error)
	CreateContractBc(ctx context.Context, req *entity.CreateContractBcReq) (*entity.CreateContractBcResp, error)
	UpdateContractBc(ctx context.Context, req *entity.UpdateContractBcReq) error
	DeleteContractBc(ctx context.Context, req *entity.DeleteContractBcReq) error

	// SaldoAwal
	GetSaldoAwals(ctx context.Context, req *entity.GetSaldoAwalsReq) (*entity.GetSaldoAwalsResp, error)
	GetSaldoAwal(ctx context.Context, req *entity.GetSaldoAwalReq) (*entity.GetSaldoAwalResp, error)
	CreateSaldoAwal(ctx context.Context, req *entity.CreateSaldoAwalReq) (*entity.CreateSaldoAwalResp, error)
	UpdateSaldoAwal(ctx context.Context, req *entity.UpdateSaldoAwalReq) error
	DeleteSaldoAwal(ctx context.Context, req *entity.DeleteSaldoAwalReq) error

	// ContractProducts
	GetContractProducts(ctx context.Context, req *entity.GetContractProductsReq) (*entity.GetContractProductsResp, error)
	GetContractProduct(ctx context.Context, req *entity.GetContractProductReq) (*entity.GetContractProductResp, error)
	CreateContractProduct(ctx context.Context, req *entity.CreateContractProductReq) (*entity.CreateContractProductResp, error)
	UpdateContractProduct(ctx context.Context, req *entity.UpdateContractProductReq) error
	DeleteContractProduct(ctx context.Context, req *entity.DeleteContractProductReq) error

	// IncomeInventories
	GetIncomeInventories(ctx context.Context, req *entity.GetIncomeInventoriesReq) (*entity.GetIncomeInventoriesResp, error)
	GetIncomeInventory(ctx context.Context, req *entity.GetIncomeInventoryReq) (*entity.GetIncomeInventoryResp, error)
	CreateIncomeInventory(ctx context.Context, req *entity.CreateIncomeInventoryReq) (*entity.CreateIncomeInventoryResp, error)
	UpdateIncomeInventory(ctx context.Context, req *entity.UpdateIncomeInventoryReq) error
	DeleteIncomeInventory(ctx context.Context, req *entity.DeleteIncomeInventoryReq) error

	// OutcomeInventoriesProducts
	GetOutcomesInventoriesProducts(ctx context.Context, req *entity.GetOutcomesInventoriesProductsReq) (*entity.GetOutcomesInventoriesProductsResp, error)
	GetOutcomesInventoriesProduct(ctx context.Context, req *entity.GetOutcomesInventoriesProductReq) (*entity.GetOutcomesInventoriesProductResp, error)
	CreateOutcomesInventoriesProduct(ctx context.Context, req *entity.CreateOutcomesInventoriesProductReq) (*entity.CreateOutcomesInventoriesProductResp, error)
	UpdateOutcomesInventoriesProduct(ctx context.Context, req *entity.UpdateOutcomesInventoriesProductReq) error
	DeleteOutcomesInventoriesProduct(ctx context.Context, req *entity.DeleteOutcomesInventoriesProductReq) error

	// IncomeInventoryProducts
	GetIncomeInventoryProducts(ctx context.Context, req *entity.GetIncomeInventoryProductsReq) (*entity.GetIncomeInventoryProductsResp, error)
	GetIncomeInventoryProduct(ctx context.Context, req *entity.GetIncomeInventoryProductReq) (*entity.GetIncomeInventoryProductResp, error)
	CreateIncomeInventoryProduct(ctx context.Context, req *entity.CreateIncomeInventoryProductReq) (*entity.CreateIncomeInventoryProductResp, error)
	UpdateIncomeInventoryProduct(ctx context.Context, req *entity.UpdateIncomeInventoryProductReq) error
	DeleteIncomeInventoryProduct(ctx context.Context, req *entity.DeleteIncomeInventoryProductReq) error
	GetIncomeInventoryProductsByContract(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractReq) (*entity.GetIncomeInventoryProductsByContractResp, error)
	GetIncomeInventoryProductsByContractAndKode(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractAndKodeReq) (*entity.GetIncomeInventoryProductsByContractAndKodeResp, error)

	// TransactionIncomes
	GetTransactionIncomes(ctx context.Context, req *entity.GetTransactionIncomesReq) (*entity.GetTransactionIncomesResp, error)
	GetTransactionIncome(ctx context.Context, req *entity.GetTransactionIncomeReq) (*entity.GetTransactionIncomeResp, error)
	CreateTransactionIncome(ctx context.Context, req *entity.CreateTransactionIncomeReq) (*entity.CreateTransactionIncomeResp, error)
	UpdateTransactionIncome(ctx context.Context, req *entity.UpdateTransactionIncomeReq) error
	DeleteTransactionIncome(ctx context.Context, req *entity.DeleteTransactionIncomeReq) error

	// ReadyProducts
	GetReadyProducts(ctx context.Context, req *entity.GetReadyProductsReq) (*entity.GetReadyProductsResp, error)
	GetReadyProduct(ctx context.Context, req *entity.GetReadyProductReq) (*entity.GetReadyProductResp, error)
	CreateReadyProduct(ctx context.Context, req *entity.CreateReadyProductReq) (*entity.CreateReadyProductResp, error)
	UpdateReadyProduct(ctx context.Context, req *entity.UpdateReadyProductReq) error
	DeleteReadyProduct(ctx context.Context, req *entity.DeleteReadyProductReq) error

	// Buyers
	GetBuyers(ctx context.Context, req *entity.GetBuyersReq) (*entity.GetBuyersResp, error)
	GetBuyer(ctx context.Context, req *entity.GetBuyerReq) (*entity.GetBuyerResp, error)
	CreateBuyer(ctx context.Context, req *entity.CreateBuyerReq) (*entity.CreateBuyerResp, error)
	UpdateBuyer(ctx context.Context, req *entity.UpdateBuyerReq) error
	DeleteBuyer(ctx context.Context, req *entity.DeleteBuyerReq) error

	// TransfersProducts
	GetTransfersProducts(ctx context.Context, req *entity.GetTransfersProductsReq) (*entity.GetTransfersProductsResp, error)
	GetTransferProduct(ctx context.Context, req *entity.GetTransferProductReq) (*entity.GetTransferProductResp, error)
	CreateTransferProduct(ctx context.Context, req *entity.CreateTransferProductReq) (*entity.CreateTransferProductResp, error)
	UpdateTransferProduct(ctx context.Context, req *entity.UpdateTransferProductReq) error
	DeleteTransferProduct(ctx context.Context, req *entity.DeleteTransferProductReq) error

	// Laporan
	GetLaporanMutasi(ctx context.Context, req *entity.GetLaporanMutasiReq) (*entity.GetLaporanMutasiResp, error)
	GetLaporanMutasiPemasukan(ctx context.Context, req *entity.GetLaporanMutasiPemasukanReq) (*entity.GetLaporanMutasiPemasukanResp, error)
	GetLaporanMutasiWip(ctx context.Context, req *entity.GetLaporanMutasiWipReq) (*entity.GetLaporanMutasiWipResp, error)
	GetLaporanMutasiJenisDokumen(ctx context.Context, req *entity.GetLaporanMutasiJenisDokumenReq) (*entity.GetLaporanMutasiJenisDokumenResp, error)

	// Productions
	GetProductions(ctx context.Context, req *entity.GetProductionsReq) (*entity.GetProductionsResp, error)
	GetProduction(ctx context.Context, req *entity.GetProductionReq) (*entity.GetProductionResp, error)
	CreateProduction(ctx context.Context, req *entity.CreateProductionReq) (*entity.CreateProductionResp, error)
	UpdateProduction(ctx context.Context, req *entity.UpdateProductionReq) error
	DeleteProduction(ctx context.Context, req *entity.DeleteProductionReq) error

	// ProductsMovement
	GetProductsMovement(ctx context.Context, req *entity.GetProductsMovementReq) (*entity.GetProductsMovementResp, error)
	GetProductsMovementByID(ctx context.Context, req *entity.GetProductsMovementReqID) (*entity.GetProductsMovementRespID, error)
	CreateProductsMovement(ctx context.Context, req *entity.CreateProductsMovementReq) (*entity.CreateProductsMovementResp, error)
	UpdateProductsMovement(ctx context.Context, req *entity.UpdateProductsMovementReq) error
	UpdateStatusProductsMovement(ctx context.Context, req *entity.UpdateStatusProductsMoveMentReq) error
	DeleteProductsMovement(ctx context.Context, req *entity.DeleteProductsMovementReq) error
}
