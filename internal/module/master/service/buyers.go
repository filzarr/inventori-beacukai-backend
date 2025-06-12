package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetBuyers(ctx context.Context, req *entity.GetBuyersReq) (*entity.GetBuyersResp, error) {
	return s.repo.GetBuyers(ctx, req)
}

func (s *masterService) GetBuyer(ctx context.Context, req *entity.GetBuyerReq) (*entity.GetBuyerResp, error) {
	return s.repo.GetBuyer(ctx, req)
}

func (s *masterService) CreateBuyer(ctx context.Context, req *entity.CreateBuyerReq) (*entity.CreateBuyerResp, error) {
	return s.repo.CreateBuyer(ctx, req)
}

func (s *masterService) UpdateBuyer(ctx context.Context, req *entity.UpdateBuyerReq) error {
	return s.repo.UpdateBuyer(ctx, req)
}

func (s *masterService) DeleteBuyer(ctx context.Context, req *entity.DeleteBuyerReq) error {
	return s.repo.DeleteBuyer(ctx, req)
}
