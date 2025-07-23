package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"

	"github.com/rs/zerolog/log"
)

func (r *userRepo) GetRole(ctx context.Context, req *entity.GetRolesReq) (*entity.GetRolesResp, error) {
	type dao struct {
		entity.Role
	}
	var (
		resp = new(entity.GetRolesResp)
		data = make([]dao, 0)
	)
	resp.Items = make([]entity.Role, 0)
	query := `SELECT id, name from roles WHERE deleted_at IS NULL`

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query)); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetRoles - failed to query")
		return nil, err
	}
	for _, d := range data {
		resp.Items = append(resp.Items, d.Role)
	}

	return resp, nil
}
