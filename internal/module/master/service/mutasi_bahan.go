package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetMutasiBahans(ctx context.Context, req *entity.GetMutasiBahansReq) (*entity.GetMutasiBahansResp, error) {
	return s.repo.GetMutasiBahans(ctx, req)
}
func (s *masterService) GetMutasiBahan(ctx context.Context, req *entity.GetMutasiBahanReq) (*entity.GetMutasiBahanResp, error) {
	return s.repo.GetMutasiBahan(ctx, req)
}
func (s *masterService) CreateMutasiBahan(ctx context.Context, req *entity.CreateMutasiBahanReq) (*entity.CreateMutasiBahanResp, error) {
	return s.repo.CreateMutasiBahan(ctx, req)
}
func (s *masterService) UpdateMutasiBahan(ctx context.Context, req *entity.UpdateMutasiBahanReq) error {
	return s.repo.UpdateMutasiBahan(ctx, req)
}

func (s *masterService) DeleteMutasiBahan(ctx context.Context, req *entity.DeleteMutasiBahanReq) error {
	return s.repo.DeleteMutasiBahan(ctx, req)
}
func (s *masterService) UpdateStatusMutasiBahan(ctx context.Context, req *entity.UpdateStatusMutasiBahanReq) error {
	return s.repo.UpdateStatusMutasiBahan(ctx, req)
}
func (s *masterService) UpdateSaldoMutasi(ctx context.Context, req *entity.UpdateSaldoMutasiReq) error {
	return s.repo.UpdateSaldoMutasi(ctx, req)
}
func (s *masterService) GetLaporanMutasiBahan(ctx context.Context, req *entity.GetLaporanMutasiBahanReq) (*entity.GetLaporanMutasiBahanResp, error) {
	return s.repo.GetLaporanMutasiBahan(ctx, req)
}
