package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetProductions(ctx context.Context, req *entity.GetProductionsReq) (*entity.GetProductionsResp, error) {
	return s.repo.GetProductions(ctx, req)
}

func (s *masterService) GetProduction(ctx context.Context, req *entity.GetProductionReq) (*entity.GetProductionResp, error) {
	return s.repo.GetProduction(ctx, req)
}

func (s *masterService) CreateProduction(ctx context.Context, req *entity.CreateProductionReq) (*entity.CreateProductionResp, error) {
	return s.repo.CreateProduction(ctx, req)
}

func (s *masterService) UpdateProduction(ctx context.Context, req *entity.UpdateProductionReq) error {
	return s.repo.UpdateProduction(ctx, req)
}

func (s *masterService) DeleteProduction(ctx context.Context, req *entity.DeleteProductionReq) error {
	return s.repo.DeleteProduction(ctx, req)
}
