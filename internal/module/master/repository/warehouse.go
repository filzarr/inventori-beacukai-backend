package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetWarehouses(ctx context.Context, req *entity.GetWarehousesReq) (*entity.GetWarehousesResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Warehouse
	}

	var (
		resp  = new(entity.GetWarehousesResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, id, kode, nama, kategori , keterangan
			FROM warehouses
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.Warehouse, 0)

	if req.Q != "" {
		query += ` AND (
			kode ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetWarehouse - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Warehouse)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetWarehouse(ctx context.Context, req *entity.GetWarehouseReq) (*entity.GetWarehouseResp, error) {
	var (
		resp = new(entity.GetWarehouseResp)
		data = new(entity.Warehouse)
	)

	query := `
		SELECT id, kode, nama, kategori ,keterangan
		FROM warehouses
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetBcDocument - not found")
			return nil, errmg.NewCustomErrors(404).SetMessage("Dokumen BC tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetBcDocument - failed to get")
		return nil, err
	}

	resp.Warehouse = *data
	return resp, nil
}

func (r *masterRepo) CreateWarehouse(ctx context.Context, req *entity.CreateWarehouseReq) (*entity.CreateWarehouseResp, error) {
	query := `INSERT INTO warehouses (id, kode, nama, kategori, keterangan) VALUES (?, ?, ?, ?, ?)`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateWarehouseResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.Kode, req.Nama, req.Kategori, req.Keterangan); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::createWarehouse - failed to create")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateWarehouse(ctx context.Context, req *entity.UpdateWarehouseReq) error {
	query := `UPDATE warehouses SET kode = ?, nama = ?, keterangan = ?, kategori = ? WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Kode, req.Nama, req.Keterangan, req.Kategori, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateWarehouse - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteWarehouse(ctx context.Context, req *entity.DeleteWarehouseReq) error {
	query := `UPDATE warehouses SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteWarehouse - failed to delete")
		return err
	}

	return nil
}
