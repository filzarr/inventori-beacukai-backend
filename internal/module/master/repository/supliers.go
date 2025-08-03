package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetSupliers(ctx context.Context, req *entity.GetSupliersReq) (*entity.GetSupliersResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Suplier
	}

	var (
		resp = new(entity.GetSupliersResp)
		data = make([]dao, 0)
		args = make([]any, 0, 3)
	)
	resp.Items = make([]entity.Suplier, 0)

	query := `
		SELECT
			COUNT(*) OVER() AS total_data,
			id,
			name,
			alamat,
			npwp
		FROM supliers
		WHERE deleted_at IS NULL
	`

	if req.Q != "" {
		query += ` AND (
			name ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetSupliers - failed to query supliers")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Suplier)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetSuplier(ctx context.Context, req *entity.GetSuplierReq) (*entity.GetSuplierResp, error) {
	var (
		resp = new(entity.GetSuplierResp)
		data = new(entity.Suplier)
	)

	query := `
		SELECT id, name, alamat, npwp
		FROM supliers
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetSuplier - suplier not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Pemasok tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetSuplier - failed to get suplier")
		return nil, err
	}

	resp.Suplier = *data
	return resp, nil
}

func (r *masterRepo) CreateSuplier(ctx context.Context, req *entity.CreateSuplierReq) (*entity.CreateSuplierResp, error) {
	query := `
		INSERT INTO supliers (
			id,
			name,
			alamat,
			npwp
		) VALUES (?, ?, ?, ?)`

	var (
		id   = ulid.Make().String()
		resp = new(entity.CreateSuplierResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), id, req.Name, req.Alamat, req.Npwp); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateSuplier - failed to insert suplier")
		return nil, err
	}

	resp.Id = id
	return resp, nil
}

func (r *masterRepo) UpdateSuplier(ctx context.Context, req *entity.UpdateSuplierReq) error {
	query := `
		UPDATE supliers
		SET
			name = ?,
			alamat = ?,
			npwp = ?,
			updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Name, req.Alamat, req.Npwp, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateSuplier - failed to update suplier")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteSuplier(ctx context.Context, req *entity.DeleteSuplierReq) error {
	query := `
		UPDATE supliers
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteSuplier - failed to soft delete suplier")
		return err
	}

	return nil
}
