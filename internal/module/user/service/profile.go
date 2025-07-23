package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

func (s *userService) GetProfile(ctx context.Context, req *entity.AuthListenReq) (*entity.AuthListenResp, error) {
	return s.repo.GetProfile(ctx, req)
}
func (s *userService) UpdateProfile(ctx context.Context, req *entity.UpdateProfileReq) error {
	return s.repo.UpdateProfile(ctx, req)
}
