package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetProducts(ctx context.Context, req *entity.GetProductsReq) (*entity.GetProductsResp, error) {
	return s.repo.GetProducts(ctx, req)
}

func (s *masterService) GetProduct(ctx context.Context, req *entity.GetProductReq) (*entity.GetProductResp, error) {
	return s.repo.GetProduct(ctx, req)
}

func (s *masterService) CreateProduct(ctx context.Context, req *entity.CreateProductReq) (*entity.CreateProductResp, error) {
	return s.repo.CreateProduct(ctx, req)
}

func (s *masterService) UpdateProduct(ctx context.Context, req *entity.UpdateProductReq) error {
	return s.repo.UpdateProduct(ctx, req)
}

func (s *masterService) DeleteProduct(ctx context.Context, req *entity.DeleteProductReq) error {
	return s.repo.DeleteProduct(ctx, req)
}
