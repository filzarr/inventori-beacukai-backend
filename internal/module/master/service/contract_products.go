package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetContractProducts(ctx context.Context, req *entity.GetContractProductsReq) (*entity.GetContractProductsResp, error) {
	return s.repo.GetContractProducts(ctx, req)
}

func (s *masterService) GetContractProduct(ctx context.Context, req *entity.GetContractProductReq) (*entity.GetContractProductResp, error) {
	return s.repo.GetContractProduct(ctx, req)
}

func (s *masterService) CreateContractProduct(ctx context.Context, req *entity.CreateContractProductReq) (*entity.CreateContractProductResp, error) {
	return s.repo.CreateContractProduct(ctx, req)
}

func (s *masterService) UpdateContractProduct(ctx context.Context, req *entity.UpdateContractProductReq) error {
	return s.repo.UpdateContractProduct(ctx, req)
}

func (s *masterService) DeleteContractProduct(ctx context.Context, req *entity.DeleteContractProductReq) error {
	return s.repo.DeleteContractProduct(ctx, req)
}
