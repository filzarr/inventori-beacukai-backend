package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/log/entity"
	"inventori-beacukai-backend/internal/module/log/ports"
)

var _ ports.LogService = &logService{}

type logService struct {
	repo ports.LogService
}

func NewLogService(repo ports.LogRepository) *logService {
	return &logService{
		repo: repo,
	}
}

func (s *logService) GetLogs(ctx context.Context, req *entity.GetLogReq) (*entity.GetLogResp, error) {
	return s.repo.GetLogs(ctx, req)
}
