package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/log/entity"
)

type LogRepository interface {
	GetLogs(ctx context.Context, req *entity.GetLogReq) (*entity.GetLogResp, error)
}

type LogService interface {
	GetLogs(ctx context.Context, req *entity.GetLogReq) (*entity.GetLogResp, error)
}
