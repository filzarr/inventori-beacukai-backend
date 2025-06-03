package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

type MasterRepository interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error)
	GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error)
	CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error)
	UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error
	DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error
}

type MasterService interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error)
	GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error)
	CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error)
	UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error
	DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error
}
