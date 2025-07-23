package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) ChangePassword(ctx context.Context, req *entity.ChangePasswordReq) error {
	return s.repo.ChangePassword(ctx, req)
}
