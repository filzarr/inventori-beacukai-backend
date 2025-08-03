package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetWarehouses(ctx context.Context, req *entity.GetWarehousesReq) (*entity.GetWarehousesResp, error) {
	return s.repo.GetWarehouses(ctx, req)
}

func (s *masterService) GetWarehouse(ctx context.Context, req *entity.GetWarehouseReq) (*entity.GetWarehouseResp, error) {
	return s.repo.GetWarehouse(ctx, req)
}

func (s *masterService) CreateWarehouse(ctx context.Context, req *entity.CreateWarehouseReq) (*entity.CreateWarehouseResp, error) {
	return s.repo.CreateWarehouse(ctx, req)
}

func (s *masterService) UpdateWarehouse(ctx context.Context, req *entity.UpdateWarehouseReq) error {
	return s.repo.UpdateWarehouse(ctx, req)
}

func (s *masterService) DeleteWarehouse(ctx context.Context, req *entity.DeleteWarehouseReq) error {
	return s.repo.DeleteWarehouse(ctx, req)
}
