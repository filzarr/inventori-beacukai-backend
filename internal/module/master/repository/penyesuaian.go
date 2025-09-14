package repository

import (
	"context"
	"fmt"
	"inventori-beacukai-backend/internal/module/master/entity"

	"github.com/oklog/ulid/v2"
)

func (r *masterRepo) GetPenyesuaian(ctx context.Context, req *entity.GetPenyesuaianReq) (*entity.GetPenyesuaianResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Penyesuaian
	}

	var (
		resp  = new(entity.GetPenyesuaianResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
			       p.id, p.kode_barang, b.nama AS nama_barang, w.nama AS gudang, p.jumlah
			FROM penyesuaian p
			JOIN products b ON p.kode_barang = b.kode
			LEFT JOIN warehouses w ON p.warehouse_kode = w.kode
			WHERE p.deleted_at IS NULL`
	)

	if req.Q != "" {
		query += ` AND (p.kode_barang ILIKE '%' || ? || '%' OR b.nama ILIKE '%' || ? || '%')`
		args = append(args, req.Q, req.Q)
	}
	if req.KodeBarang != "" {
		query += ` AND p.kode_barang = ?`
		args = append(args, req.KodeBarang)
	}
	query += ` ORDER BY p.created_at DESC LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		return nil, err
	}
	resp.Items = make([]entity.Penyesuaian, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Penyesuaian)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) CreatePenyesuaian(ctx context.Context, req *entity.CreatePenyesuaianReq) (*entity.CreatePenyesuaianResp, error) {
	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreatePenyesuaianResp)
	)
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	var currentStock int
	queryCheck := `SELECT jumlah FROM warehouses_stocks WHERE kode_barang = $1 AND warehouse_kode = $2 FOR UPDATE`
	err = tx.GetContext(ctx, &currentStock, queryCheck, req.KodeBarang, req.Warehouse)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if currentStock < req.Jumlah {
		tx.Rollback()
		return nil, fmt.Errorf("stok tidak mencukupi: stok saat ini %d, diminta %d", currentStock, req.Jumlah)
	}
	newStock := currentStock - req.Jumlah

	// Catat jumlah penyesuaian (jumlah yang dipindahkan/diambil), bukan stok baru
	query := `INSERT INTO penyesuaian (id, kode_barang, warehouse_kode, jumlah) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, query, Id, req.KodeBarang, req.Warehouse, req.Jumlah)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// Perbarui stok gudang menjadi stok baru (bukan jumlah yang diminta)
	queryUpdate := `UPDATE warehouses_stocks SET jumlah = $1 WHERE kode_barang = $2 AND warehouse_kode = $3`
	_, err = tx.ExecContext(ctx, queryUpdate, newStock, req.KodeBarang, req.Warehouse)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	resp.Id = Id
	return resp, nil
}
