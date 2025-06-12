package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetReadyProducts(ctx context.Context, req *entity.GetReadyProductsReq) (*entity.GetReadyProductsResp, error) {
	return s.repo.GetReadyProducts(ctx, req)
}

func (s *masterService) GetReadyProduct(ctx context.Context, req *entity.GetReadyProductReq) (*entity.GetReadyProductResp, error) {
	return s.repo.GetReadyProduct(ctx, req)
}

func (s *masterService) CreateReadyProduct(ctx context.Context, req *entity.CreateReadyProductReq) (*entity.CreateReadyProductResp, error) {
	return s.repo.CreateReadyProduct(ctx, req)
}

func (s *masterService) UpdateReadyProduct(ctx context.Context, req *entity.UpdateReadyProductReq) error {
	return s.repo.UpdateReadyProduct(ctx, req)
}

func (s *masterService) DeleteReadyProduct(ctx context.Context, req *entity.DeleteReadyProductReq) error {
	return s.repo.DeleteReadyProduct(ctx, req)
}
