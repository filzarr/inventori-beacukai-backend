package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error) {
	return s.repo.Login(ctx, req)
}
