package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

type UserRepository interface {
	Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error)
}

type UserService interface {
	Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error)
}
