package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

type MasterRepository interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
}

type MasterService interface {
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
}
