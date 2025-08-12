package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetWarehousesStocks(ctx context.Context, req *entity.GetWarehousesStocksReq) (*entity.GetWarehousesStocksResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.WarehousesStock
	}

	var (
		resp  = new(entity.GetWarehousesStocksResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, ws.id, ws.warehouse_kode, ws.kode_barang, ws.jumlah, p.nama AS nama_barang, p.satuan AS satuan, p.kategori AS kategori
			FROM warehouses_stocks ws
			LEFT JOIN products p ON ws.kode_barang = p.kode
			WHERE ws.deleted_at IS NULL`
	)

	resp.Items = make([]entity.WarehousesStock, 0)
	if req.WarehouseKode != "" {
		query += ` AND ws.warehouse_kode = ?`
		args = append(args, req.WarehouseKode)
	}
	if req.Q != "" {
		query += ` AND (
			ws.warehouse_kode ILIKE '%' || ? || '%' OR
			ws.kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetWarehousesStocks - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.WarehousesStock)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetWarehousesStock(ctx context.Context, req *entity.GetWarehousesStockReq) (*entity.GetWarehousesStockResp, error) {
	var (
		resp = new(entity.GetWarehousesStockResp)
		data = new(entity.WarehousesStock)
	)

	query := `
		SELECT id, warehouse_kode, kode_barang, jumlah
		FROM warehouses_stocks
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetWarehousesStock - not found")
			return nil, errmg.NewCustomErrors(404).SetMessage("Stok gudang tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetWarehousesStock - failed to get")
		return nil, err
	}

	resp.WarehousesStock = *data
	return resp, nil
}

func (r *masterRepo) CreateWarehousesStock(ctx context.Context, req *entity.CreateWarehousesStockReq) (*entity.CreateWarehousesStockResp, error) {
	query := `INSERT INTO warehouses_stocks (id, warehouse_kode, kode_barang, jumlah) VALUES (?, ?, ?, ?)`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateWarehousesStockResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.WarehouseKode, req.KodeBarang, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateWarehousesStock - failed to create")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateWarehousesStock(ctx context.Context, req *entity.UpdateWarehousesStockReq) error {
	query := `UPDATE warehouses_stocks SET warehouse_kode = ?, kode_barang = ?, jumlah = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.WarehouseKode, req.KodeBarang, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateWarehousesStock - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteWarehousesStock(ctx context.Context, req *entity.DeleteWarehousesStockReq) error {
	query := `UPDATE warehouses_stocks SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteWarehousesStock - failed to delete")
		return err
	}

	return nil
}

func (r *masterRepo) UpdateStockWarehouses(ctx context.Context, req *entity.UpdateStockWarehousesReq) error {
	query := `UPDATE warehouses_stocks SET 
	 jumlah = jumlah + ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateStockWarehouse - failed to update")
		return err
	}

	return nil
}
