package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetProductsMovement(ctx context.Context, req *entity.GetProductsMovementReq) (*entity.GetProductsMovementResp, error) {
	return s.repo.GetProductsMovement(ctx, req)
}

func (s *masterService) GetProductsMovementByID(ctx context.Context, req *entity.GetProductsMovementReqID) (*entity.GetProductsMovementRespID, error) {
	return s.repo.GetProductsMovementByID(ctx, req)
}

func (s *masterService) CreateProductsMovement(ctx context.Context, req *entity.CreateProductsMovementReq) (*entity.CreateProductsMovementResp, error) {
	return s.repo.CreateProductsMovement(ctx, req)
}

func (s *masterService) UpdateProductsMovement(ctx context.Context, req *entity.UpdateProductsMovementReq) error {
	return s.repo.UpdateProductsMovement(ctx, req)
}

func (s *masterService) DeleteProductsMovement(ctx context.Context, req *entity.DeleteProductsMovementReq) error {
	return s.repo.DeleteProductsMovement(ctx, req)
}

func (s *masterService) UpdateStatusProductsMovement(ctx context.Context, req *entity.UpdateStatusProductsMoveMentReq) error {
	return s.repo.UpdateStatusProductsMovement(ctx, req)
}
