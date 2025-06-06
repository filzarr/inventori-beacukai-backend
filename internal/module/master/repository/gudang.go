package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetGudangs(ctx context.Context, req *entity.GetGudangsReq) (*entity.GetGudangsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Gudang
	}

	var (
		resp    = new(entity.GetGudangsResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 3)
		filters = []string{"g.deleted_at IS NULL"}
	)
	resp.Items = make([]entity.Gudang, 0)

	query := `
		SELECT 
			g.id,
			g.nama,
			g.lokasi,
			u.name as user
		FROM
			gudang g
		JOIN
			users u
			on  g.user_id = u.id 
	`

	if req.Q != "" {
		filters = append(filters, "g.nama ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetGudangs - failed to query Gudang")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Gudang)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetGudang(ctx context.Context, req *entity.GetGudangReq) (*entity.GetGudangResp, error) {
	var (
		resp = new(entity.GetGudangResp)
		data = new(entity.Gudang)
	)

	query := `
		SELECT 
			g.id,
			g.nama,
			g.lokasi,
			u.name as user, 
		FROM
			gudang g
		JOIN
			users u
			on  g.user_id = u.id
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetGudang - gudang not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Gudang tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetGudang - failed to get Gudang")
		return nil, err
	}

	resp.Gudang = *data

	return resp, nil
}

func (r *masterRepo) CreateGudang(ctx context.Context, req *entity.CreateGudangReq) (*entity.CreateGudangResp, error) {
	query := `
		INSERT INTO gudang (
			id,
			nama,
			lokasi,
			user_id
		) VALUES (?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateGudangResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		Id, req.Nama, req.Lokasi, req.User); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateGudang - failed to create gudang")
		return nil, err
	}

	resp.Id = Id

	return resp, nil
}

func (r *masterRepo) UpdateGudang(ctx context.Context, req *entity.UpdateGudangReq) error {
	query := `
		UPDATE gudang
		SET
			nama = ?,
			lokasi = ?,
			user_id = ?
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		req.Nama, req.Lokasi, req.User, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateGudang - failed to update gudang")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteGudang(ctx context.Context, req *entity.DeleteGudangReq) error {
	query := `
		UPDATE gudang
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteGudang - failed to delete gudang")
		return err
	}

	return nil
}
