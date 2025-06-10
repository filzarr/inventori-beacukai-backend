package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetContracts(ctx context.Context, req *entity.GetContractsReq) (*entity.GetContractsResp, error) {
	return s.repo.GetContracts(ctx, req)
}

func (s *masterService) GetContract(ctx context.Context, req *entity.GetContractReq) (*entity.GetContractResp, error) {
	return s.repo.GetContract(ctx, req)
}

func (s *masterService) CreateContract(ctx context.Context, req *entity.CreateContractReq) (*entity.CreateContractResp, error) {
	return s.repo.CreateContract(ctx, req)
}

func (s *masterService) UpdateContract(ctx context.Context, req *entity.UpdateContractReq) error {
	return s.repo.UpdateContract(ctx, req)
}

func (s *masterService) DeleteContract(ctx context.Context, req *entity.DeleteContractReq) error {
	return s.repo.DeleteContract(ctx, req)
}
