package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetReadyProducts(ctx context.Context, req *entity.GetReadyProductsReq) (*entity.GetReadyProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ReadyProduct
	}

	var (
		resp  = new(entity.GetReadyProductsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
				id,
				kode,
				nama,
				satuan,
				jumlah
			FROM products
			WHERE deleted_at IS NULL AND kategori = 'Barang Jadi'`
	)

	resp.Items = make([]entity.ReadyProduct, 0)

	if req.Q != "" {
		query += ` AND (
			nama ILIKE '%' || ? || '%'
			OR kode ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetReadyProducts - failed to query ready_products")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.ReadyProduct)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetReadyProduct(ctx context.Context, req *entity.GetReadyProductReq) (*entity.GetReadyProductResp, error) {
	var data = new(entity.ReadyProduct)

	query := `
		SELECT id, kode, nama,satuan, jumlah
		FROM products
		WHERE id = ? AND deleted_at IS NULL AND kategori = Barang Jadi`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetReadyProduct - ready_product not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Produk jadi tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetReadyProduct - failed to get ready_product")
		return nil, err
	}

	return &entity.GetReadyProductResp{ReadyProduct: *data}, nil
}

func (r *masterRepo) CreateReadyProduct(ctx context.Context, req *entity.CreateReadyProductReq) (*entity.CreateReadyProductResp, error) {
	query := `
		INSERT INTO products (id, kode, satuan, nama, jumlah, kategori)
		VALUES (?, ?, ?, ?, ?, ?)`

	Id := ulid.Make().String()

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.Kode, req.Satuan, req.Nama, req.Jumlah, "Barang Jadi"); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateReadyProduct - failed to create ready_product")
		return nil, err
	}

	return &entity.CreateReadyProductResp{Id: Id}, nil
}

func (r *masterRepo) UpdateReadyProduct(ctx context.Context, req *entity.UpdateReadyProductReq) error {
	query := `
		UPDATE products
		SET kode = ?, nama = ?, jumlah = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Kode, req.Nama, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateReadyProduct - failed to update ready_product")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteReadyProduct(ctx context.Context, req *entity.DeleteReadyProductReq) error {
	query := `
		UPDATE ready_products
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteReadyProduct - failed to soft delete ready_product")
		return err
	}
	return nil
}
