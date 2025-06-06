package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetGudangs(ctx context.Context, req *entity.GetGudangsReq) (*entity.GetGudangsResp, error) {
	return s.repo.GetGudangs(ctx, req)
}
func (s *masterService) GetGudang(ctx context.Context, req *entity.GetGudangReq) (*entity.GetGudangResp, error) {
	return s.repo.GetGudang(ctx, req)
}
func (s *masterService) CreateGudang(ctx context.Context, req *entity.CreateGudangReq) (*entity.CreateGudangResp, error) {
	return s.repo.CreateGudang(ctx, req)
}
func (s *masterService) UpdateGudang(ctx context.Context, req *entity.UpdateGudangReq) error {
	return s.repo.UpdateGudang(ctx, req)
}

func (s *masterService) DeleteGudang(ctx context.Context, req *entity.DeleteGudangReq) error {
	return s.repo.DeleteGudang(ctx, req)
}
