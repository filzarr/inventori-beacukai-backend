package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetIncomeInventoryProducts(ctx context.Context, req *entity.GetIncomeInventoryProductsReq) (*entity.GetIncomeInventoryProductsResp, error) {
	return s.repo.GetIncomeInventoryProducts(ctx, req)
}

func (s *masterService) GetIncomeInventoryProduct(ctx context.Context, req *entity.GetIncomeInventoryProductReq) (*entity.GetIncomeInventoryProductResp, error) {
	return s.repo.GetIncomeInventoryProduct(ctx, req)
}

func (s *masterService) CreateIncomeInventoryProduct(ctx context.Context, req *entity.CreateIncomeInventoryProductReq) (*entity.CreateIncomeInventoryProductResp, error) {
	return s.repo.CreateIncomeInventoryProduct(ctx, req)
}

func (s *masterService) UpdateIncomeInventoryProduct(ctx context.Context, req *entity.UpdateIncomeInventoryProductReq) error {
	return s.repo.UpdateIncomeInventoryProduct(ctx, req)
}

func (s *masterService) DeleteIncomeInventoryProduct(ctx context.Context, req *entity.DeleteIncomeInventoryProductReq) error {
	return s.repo.DeleteIncomeInventoryProduct(ctx, req)
}

func (s *masterService) GetIncomeInventoryProductsByContract(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractReq) (*entity.GetIncomeInventoryProductsByContractResp, error) {
	return s.repo.GetIncomeInventoryProductsByContract(ctx, req)
}

func (s *masterService) GetIncomeInventoryProductsByContractAndKode(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractAndKodeReq) (*entity.GetIncomeInventoryProductsByContractAndKodeResp, error) {
	return s.repo.GetIncomeInventoryProductsByContractAndKode(ctx, req)
}
