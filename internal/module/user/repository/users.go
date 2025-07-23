package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"

	"github.com/rs/zerolog/log"
)

func (r *userRepo) GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Users
	}

	var (
		resp  = new(entity.GetUsersResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, u.id, u.name, u.email, r.name AS role  
			FROM users u
			JOIN roles r ON u.role_id = r.id 
			WHERE u.deleted_at IS NULL AND u.superadmin IS false`
	)

	resp.Items = make([]entity.Users, 0)

	if req.Q != "" {
		query += ` AND (
			u.name ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetUsers - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Users)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteUser - failed to delete")
		return err
	}

	return nil
}
