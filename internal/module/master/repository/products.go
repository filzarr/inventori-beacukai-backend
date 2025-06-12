package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetProducts(ctx context.Context, req *entity.GetProductsReq) (*entity.GetProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Product
	}

	var (
		resp  = new(entity.GetProductsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
				id,
				kode,
				nama,
				satuan,
				kategori,
				saldo_awal,
				jumlah
			FROM products
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.Product, 0)

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
		log.Error().Err(err).Any("req", req).Msg("repo::GetProducts - failed to query products")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Product)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetProduct(ctx context.Context, req *entity.GetProductReq) (*entity.GetProductResp, error) {
	var data = new(entity.Product)

	query := `
		SELECT id, kode, nama, kategori, jumlah
		FROM products
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetProduct - product not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Produk tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetProduct - failed to get product")
		return nil, err
	}

	return &entity.GetProductResp{Product: *data}, nil
}

func (r *masterRepo) CreateProduct(ctx context.Context, req *entity.CreateProductReq) (*entity.CreateProductResp, error) {
	query := `
		INSERT INTO products (id, kode, nama, satuan, kategori, saldo_awal, jumlah)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	Id := ulid.Make().String()

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.Kode, req.Nama, req.Satuan, req.Kategori, req.SaldoAwal, req.SaldoAwal); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateProduct - failed to create product")
		return nil, err
	}

	return &entity.CreateProductResp{Id: Id}, nil
}

func (r *masterRepo) UpdateProduct(ctx context.Context, req *entity.UpdateProductReq) error {
	query := `
		UPDATE products
		SET kode = ?, nama = ?, kategori = ?, saldo_awal= ? , jumlah = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Kode, req.Nama, req.Kategori, req.SaldoAwal, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateProduct - failed to update product")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteProduct(ctx context.Context, req *entity.DeleteProductReq) error {
	query := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteProduct - failed to soft delete product")
		return err
	}
	return nil
}
