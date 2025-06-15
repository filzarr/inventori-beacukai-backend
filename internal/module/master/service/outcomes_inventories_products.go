package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetOutcomesInventoriesProducts(ctx context.Context, req *entity.GetOutcomesInventoriesProductsReq) (*entity.GetOutcomesInventoriesProductsResp, error) {
	return s.repo.GetOutcomesInventoriesProducts(ctx, req)
}

func (s *masterService) GetOutcomesInventoriesProduct(ctx context.Context, req *entity.GetOutcomesInventoriesProductReq) (*entity.GetOutcomesInventoriesProductResp, error) {
	return s.repo.GetOutcomesInventoriesProduct(ctx, req)
}

func (s *masterService) CreateOutcomesInventoriesProduct(ctx context.Context, req *entity.CreateOutcomesInventoriesProductReq) (*entity.CreateOutcomesInventoriesProductResp, error) {
	return s.repo.CreateOutcomesInventoriesProduct(ctx, req)
}

func (s *masterService) UpdateOutcomesInventoriesProduct(ctx context.Context, req *entity.UpdateOutcomesInventoriesProductReq) error {
	return s.repo.UpdateOutcomesInventoriesProduct(ctx, req)
}

func (s *masterService) DeleteOutcomesInventoriesProduct(ctx context.Context, req *entity.DeleteOutcomesInventoriesProductReq) error {
	return s.repo.DeleteOutcomesInventoriesProduct(ctx, req)
}
