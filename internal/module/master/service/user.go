package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error) {
	return s.repo.GetUsers(ctx, req)
}
