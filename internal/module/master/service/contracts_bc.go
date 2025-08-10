package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetContractsBc(ctx context.Context, req *entity.GetContractsBcReq) (*entity.GetContractsBcResp, error) {
	return s.repo.GetContractsBc(ctx, req)
}

func (s *masterService) GetContractBc(ctx context.Context, req *entity.GetContractBcReq) (*entity.GetContractBcResp, error) {
	return s.repo.GetContractBc(ctx, req)
}

func (s *masterService) CreateContractBc(ctx context.Context, req *entity.CreateContractBcReq) (*entity.CreateContractBcResp, error) {
	return s.repo.CreateContractBc(ctx, req)
}

func (s *masterService) UpdateContractBc(ctx context.Context, req *entity.UpdateContractBcReq) error {
	return s.repo.UpdateContractBc(ctx, req)
}

func (s *masterService) DeleteContractBc(ctx context.Context, req *entity.DeleteContractBcReq) error {
	return s.repo.DeleteContractBc(ctx, req)
}