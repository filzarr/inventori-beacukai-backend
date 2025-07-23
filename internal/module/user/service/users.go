package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error) {
	return s.repo.GetUsers(ctx, req)
}

func (s *userService) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	return s.repo.DeleteUser(ctx, req)
}
