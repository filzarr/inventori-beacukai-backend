package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetIncomeInventories(ctx context.Context, req *entity.GetIncomeInventoriesReq) (*entity.GetIncomeInventoriesResp, error) {
	return s.repo.GetIncomeInventories(ctx, req)
}

func (s *masterService) GetIncomeInventory(ctx context.Context, req *entity.GetIncomeInventoryReq) (*entity.GetIncomeInventoryResp, error) {
	return s.repo.GetIncomeInventory(ctx, req)
}

func (s *masterService) CreateIncomeInventory(ctx context.Context, req *entity.CreateIncomeInventoryReq) (*entity.CreateIncomeInventoryResp, error) {
	return s.repo.CreateIncomeInventory(ctx, req)
}

func (s *masterService) UpdateIncomeInventory(ctx context.Context, req *entity.UpdateIncomeInventoryReq) error {
	return s.repo.UpdateIncomeInventory(ctx, req)
}

func (s *masterService) DeleteIncomeInventory(ctx context.Context, req *entity.DeleteIncomeInventoryReq) error {
	return s.repo.DeleteIncomeInventory(ctx, req)
}
