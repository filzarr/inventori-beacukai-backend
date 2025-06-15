package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetTransfersProducts(ctx context.Context, req *entity.GetTransfersProductsReq) (*entity.GetTransfersProductsResp, error) {
	return s.repo.GetTransfersProducts(ctx, req)
}

func (s *masterService) GetTransferProduct(ctx context.Context, req *entity.GetTransferProductReq) (*entity.GetTransferProductResp, error) {
	return s.repo.GetTransferProduct(ctx, req)
}

func (s *masterService) CreateTransferProduct(ctx context.Context, req *entity.CreateTransferProductReq) (*entity.CreateTransferProductResp, error) {
	return s.repo.CreateTransferProduct(ctx, req)
}

func (s *masterService) UpdateTransferProduct(ctx context.Context, req *entity.UpdateTransferProductReq) error {
	return s.repo.UpdateTransferProduct(ctx, req)
}

func (s *masterService) DeleteTransferProduct(ctx context.Context, req *entity.DeleteTransferProductReq) error {
	return s.repo.DeleteTransferProduct(ctx, req)
}
