package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error) {
	return s.repo.GetInventories(ctx, req)
}
func (s *masterService) GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error) {
	return s.repo.GetInventory(ctx, req)
}
func (s *masterService) CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error) {
	return s.repo.CreateInventory(ctx, req)
}
func (s *masterService) UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error {
	return s.repo.UpdateInventory(ctx, req)
}

func (s *masterService) DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error {
	return s.repo.DeleteInventory(ctx, req)
}
func (s *masterService) GetInventoriesBahanBaku(ctx context.Context, req *entity.GetInventoriesBahanBakuReq) (*entity.GetInventoriesBahanBakuResp, error) {
	return s.repo.GetInventoriesBahanBaku(ctx, req)
}
