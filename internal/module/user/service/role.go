package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) GetRole(ctx context.Context, req *entity.GetRolesReq) (*entity.GetRolesResp, error) {
	return s.repo.GetRole(ctx, req)
}
