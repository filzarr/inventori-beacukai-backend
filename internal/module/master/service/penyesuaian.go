package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetPenyesuaian(ctx context.Context, req *entity.GetPenyesuaianReq) (*entity.GetPenyesuaianResp, error) {
	return s.repo.GetPenyesuaian(ctx, req)
}

func (s *masterService) CreatePenyesuaian(ctx context.Context, req *entity.CreatePenyesuaianReq) (*entity.CreatePenyesuaianResp, error) {
	return s.repo.CreatePenyesuaian(ctx, req)
}
