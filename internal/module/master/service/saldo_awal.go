package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetSaldoAwals(ctx context.Context, req *entity.GetSaldoAwalsReq) (*entity.GetSaldoAwalsResp, error) {
	return s.repo.GetSaldoAwals(ctx, req)
}

func (s *masterService) GetSaldoAwal(ctx context.Context, req *entity.GetSaldoAwalReq) (*entity.GetSaldoAwalResp, error) {
	return s.repo.GetSaldoAwal(ctx, req)
}

func (s *masterService) CreateSaldoAwal(ctx context.Context, req *entity.CreateSaldoAwalReq) (*entity.CreateSaldoAwalResp, error) {
	return s.repo.CreateSaldoAwal(ctx, req)
}

func (s *masterService) UpdateSaldoAwal(ctx context.Context, req *entity.UpdateSaldoAwalReq) error {
	return s.repo.UpdateSaldoAwal(ctx, req)
}

func (s *masterService) DeleteSaldoAwal(ctx context.Context, req *entity.DeleteSaldoAwalReq) error {
	return s.repo.DeleteSaldoAwal(ctx, req)
}
