package repository

import (
	"context"
	"database/sql"

	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetSaldoAwals(ctx context.Context, req *entity.GetSaldoAwalsReq) (*entity.GetSaldoAwalsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.SaldoAwal
	}

	var (
		resp  = new(entity.GetSaldoAwalsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, s.id, s.kode_barang, p.nama AS nama_barang, s.saldo_awal
			FROM saldo_awals s
			JOIN products p ON s.kode_barang = p.kode
			WHERE s.deleted_at IS NULL`
	)

	resp.Items = make([]entity.SaldoAwal, 0)

	if req.Q != "" {
		query += ` AND (
			kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` ORDER BY s.created_at DESC LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetSaldoAwals - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.SaldoAwal)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetSaldoAwal(ctx context.Context, req *entity.GetSaldoAwalReq) (*entity.GetSaldoAwalResp, error) {
	var (
		resp = new(entity.GetSaldoAwalResp)
		data = new(entity.SaldoAwal)
	)

	query := `SELECT id, kode_barang, saldo_awal, created_at, updated_at FROM saldo_awals WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetSaldoAwal - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Saldo awal tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetSaldoAwal - failed to get")
		return nil, err
	}

	resp.SaldoAwal = *data
	return resp, nil
}

func (r *masterRepo) CreateSaldoAwal(ctx context.Context, req *entity.CreateSaldoAwalReq) (*entity.CreateSaldoAwalResp, error) {
	query := `INSERT INTO saldo_awals (id, kode_barang, saldo_awal) VALUES (?, ?, ?)`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateSaldoAwalResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.KodeBarang, req.SaldoAwal); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateSaldoAwal - failed to create")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateSaldoAwal(ctx context.Context, req *entity.UpdateSaldoAwalReq) error {
	query := `
		UPDATE saldo_awals
		SET kode_barang = ?, saldo_awal = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.KodeBarang, req.SaldoAwal, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateSaldoAwal - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteSaldoAwal(ctx context.Context, req *entity.DeleteSaldoAwalReq) error {
	query := `UPDATE saldo_awals SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteSaldoAwal - failed to delete")
		return err
	}

	return nil
}
