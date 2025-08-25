package repository

import (
	"context"
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/log/entity"
	"inventori-beacukai-backend/internal/module/log/ports"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.LogRepository = &logRepo{}

type logRepo struct {
	db *sqlx.DB
}

func NewLogRepository() *logRepo {
	return &logRepo{
		db: adapter.Adapters.Postgres,
	}
}

func (r *logRepo) GetLogs(ctx context.Context, req *entity.GetLogReq) (*entity.GetLogResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Log
	}

	var (
		resp  = new(entity.GetLogResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, id, table_name, operation, user_id, old_data, new_data, changed_at
			FROM audit_logs ORDER BY changed_at DESC`
	)

	resp.Items = make([]entity.Log, 0)

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLogs - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Log)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}
