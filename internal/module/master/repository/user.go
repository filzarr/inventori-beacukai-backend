package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"

	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.User
	}

	var (
		resp = new(entity.GetUsersResp)
		data = make([]dao, 0)
		args = make([]any, 0, 3)
	)
	resp.Items = make([]entity.User, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			u.id,
			u.name,
			u.email,
			r.name as role
		FROM
			users u
		JOIN
			roles r
			ON u.role_id = r.id
		WHERE
			u.deleted_at IS NULL
	`

	if req.Q != "" {
		query += ` AND (
			u.name ILIKE '%' || ? || '%' OR
			u.email ILIKE '%' || ? || '%' 
		)
		`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetUsers - failed to query Users")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.User)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}
