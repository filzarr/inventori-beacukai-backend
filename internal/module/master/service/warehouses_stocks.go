package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetWarehousesStocks(ctx context.Context, req *entity.GetWarehousesStocksReq) (*entity.GetWarehousesStocksResp, error) {
	return s.repo.GetWarehousesStocks(ctx, req)
}

func (s *masterService) GetWarehousesStock(ctx context.Context, req *entity.GetWarehousesStockReq) (*entity.GetWarehousesStockResp, error) {
	return s.repo.GetWarehousesStock(ctx, req)
}

func (s *masterService) CreateWarehousesStock(ctx context.Context, req *entity.CreateWarehousesStockReq) (*entity.CreateWarehousesStockResp, error) {
	return s.repo.CreateWarehousesStock(ctx, req)
}

func (s *masterService) UpdateWarehousesStock(ctx context.Context, req *entity.UpdateWarehousesStockReq) error {
	return s.repo.UpdateWarehousesStock(ctx, req)
}

func (s *masterService) DeleteWarehousesStock(ctx context.Context, req *entity.DeleteWarehousesStockReq) error {
	return s.repo.DeleteWarehousesStock(ctx, req)
}

func (s *masterService) UpdateStockWarehouses(ctx context.Context, req *entity.UpdateStockWarehousesReq) error {
	return s.repo.UpdateStockWarehouses(ctx, req)
}
