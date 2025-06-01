package service

import "inventori-beacukai-backend/internal/module/master/ports"

var _ ports.MasterService = &masterService{}

type masterService struct {
	repo ports.MasterRepository
}

func NewMasterService(repo ports.MasterService) *masterService {
	return &masterService{
		repo: repo,
	}
}
