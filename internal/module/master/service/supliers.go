package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetSupliers(ctx context.Context, req *entity.GetSupliersReq) (*entity.GetSupliersResp, error) {
	return s.repo.GetSupliers(ctx, req)
}

func (s *masterService) GetSuplier(ctx context.Context, req *entity.GetSuplierReq) (*entity.GetSuplierResp, error) {
	return s.repo.GetSuplier(ctx, req)
}

func (s *masterService) CreateSuplier(ctx context.Context, req *entity.CreateSuplierReq) (*entity.CreateSuplierResp, error) {
	return s.repo.CreateSuplier(ctx, req)
}

func (s *masterService) UpdateSuplier(ctx context.Context, req *entity.UpdateSuplierReq) error {
	return s.repo.UpdateSuplier(ctx, req)
}

func (s *masterService) DeleteSuplier(ctx context.Context, req *entity.DeleteSuplierReq) error {
	return s.repo.DeleteSuplier(ctx, req)
}
