package repository

import (
	"context"
	"database/sql"

	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetCurrencies(ctx context.Context, req *entity.GetCurrenciesReq) (*entity.GetCurrenciesResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Currency
	}

	var (
		resp  = new(entity.GetCurrenciesResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
			id,
			kode,
			mata_uang
			FROM currencies
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.Currency, 0)

	if req.Q != "" {
		query += ` AND (kode ILIKE '%' || ? || '%' OR mata_uang ILIKE '%' || ? || '%')`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetCurrencies - failed to query currencies")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Currency)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetCurrency(ctx context.Context, req *entity.GetCurrencyReq) (*entity.GetCurrencyResp, error) {
	var data = new(entity.Currency)

	query := `
		SELECT id, kode, mata_uang
		FROM currencies
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetCurrency - currency not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data mata uang tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetCurrency - failed to get currency")
		return nil, err
	}

	return &entity.GetCurrencyResp{Currency: *data}, nil
}

func (r *masterRepo) CreateCurrency(ctx context.Context, req *entity.CreateCurrencyReq) (*entity.CreateCurrencyResp, error) {
	query := `
		INSERT INTO currencies (id, kode, mata_uang)
		VALUES (?, ?, ?)`

	id := ulid.Make().String()

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), id, req.Kode, req.MataUang); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateCurrency - failed to insert")
		return nil, err
	}

	return &entity.CreateCurrencyResp{Id: id}, nil
}

func (r *masterRepo) UpdateCurrency(ctx context.Context, req *entity.UpdateCurrencyReq) error {
	query := `
		UPDATE currencies
		SET kode = ?, mata_uang = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Kode, req.MataUang, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateCurrency - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteCurrency(ctx context.Context, req *entity.DeleteCurrencyReq) error {
	query := `
		UPDATE currencies SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteCurrency - failed to delete")
		return err
	}

	return nil
}
