package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetLaporanMutasi(ctx context.Context, req *entity.GetLaporanMutasiReq) (*entity.GetLaporanMutasiResp, error) {
	return s.repo.GetLaporanMutasi(ctx, req)
}
