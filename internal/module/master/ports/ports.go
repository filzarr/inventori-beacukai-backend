package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

type MasterRepository interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	// Inventories
	GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error)
	GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error)
	CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error)
	UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error
	DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error
	GetInventoriesBahanBaku(ctx context.Context, req *entity.GetInventoriesBahanBakuReq) (*entity.GetInventoriesBahanBakuResp, error)

	// Gudang
	GetGudangs(ctx context.Context, req *entity.GetGudangsReq) (*entity.GetGudangsResp, error)
	GetGudang(ctx context.Context, req *entity.GetGudangReq) (*entity.GetGudangResp, error)
	CreateGudang(ctx context.Context, req *entity.CreateGudangReq) (*entity.CreateGudangResp, error)
	UpdateGudang(ctx context.Context, req *entity.UpdateGudangReq) error
	DeleteGudang(ctx context.Context, req *entity.DeleteGudangReq) error
	// mutasi bahan baku
	GetMutasiBahans(ctx context.Context, req *entity.GetMutasiBahansReq) (*entity.GetMutasiBahansResp, error)
	GetMutasiBahan(ctx context.Context, req *entity.GetMutasiBahanReq) (*entity.GetMutasiBahanResp, error)
	CreateMutasiBahan(ctx context.Context, req *entity.CreateMutasiBahanReq) (*entity.CreateMutasiBahanResp, error)
	UpdateMutasiBahan(ctx context.Context, req *entity.UpdateMutasiBahanReq) error
	DeleteMutasiBahan(ctx context.Context, req *entity.DeleteMutasiBahanReq) error
	UpdateStatusMutasiBahan(ctx context.Context, req *entity.UpdateStatusMutasiBahanReq) error
	UpdateSaldoMutasi(ctx context.Context, req *entity.UpdateSaldoMutasiReq) error
	GetLaporanMutasiBahan(ctx context.Context, req *entity.GetLaporanMutasiBahanReq) (*entity.GetLaporanMutasiBahanResp, error)
}

type MasterService interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error)
	GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error)
	CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error)
	UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error
	DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error
	GetInventoriesBahanBaku(ctx context.Context, req *entity.GetInventoriesBahanBakuReq) (*entity.GetInventoriesBahanBakuResp, error)

	// Gudang
	GetGudangs(ctx context.Context, req *entity.GetGudangsReq) (*entity.GetGudangsResp, error)
	GetGudang(ctx context.Context, req *entity.GetGudangReq) (*entity.GetGudangResp, error)
	CreateGudang(ctx context.Context, req *entity.CreateGudangReq) (*entity.CreateGudangResp, error)
	UpdateGudang(ctx context.Context, req *entity.UpdateGudangReq) error
	DeleteGudang(ctx context.Context, req *entity.DeleteGudangReq) error
	// mutasi bahan baku
	GetMutasiBahans(ctx context.Context, req *entity.GetMutasiBahansReq) (*entity.GetMutasiBahansResp, error)
	GetMutasiBahan(ctx context.Context, req *entity.GetMutasiBahanReq) (*entity.GetMutasiBahanResp, error)
	CreateMutasiBahan(ctx context.Context, req *entity.CreateMutasiBahanReq) (*entity.CreateMutasiBahanResp, error)
	UpdateMutasiBahan(ctx context.Context, req *entity.UpdateMutasiBahanReq) error
	DeleteMutasiBahan(ctx context.Context, req *entity.DeleteMutasiBahanReq) error
	UpdateStatusMutasiBahan(ctx context.Context, req *entity.UpdateStatusMutasiBahanReq) error
	UpdateSaldoMutasi(ctx context.Context, req *entity.UpdateSaldoMutasiReq) error
	GetLaporanMutasiBahan(ctx context.Context, req *entity.GetLaporanMutasiBahanReq) (*entity.GetLaporanMutasiBahanResp, error)
}
