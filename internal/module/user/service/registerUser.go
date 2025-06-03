package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) RegisterUser(ctx context.Context, req *entity.RegisterReq) (*entity.RegisterResp, error) {
	return s.repo.RegisterUser(ctx, req)
}
