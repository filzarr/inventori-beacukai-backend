package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetTransfersProducts(ctx context.Context, req *entity.GetTransfersProductsReq) (*entity.GetTransfersProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.TransferProduct
	}

	var (
		resp  = new(entity.GetTransfersProductsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
				id,
				kode_barang,
				jumlah
			FROM transfers_products
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.TransferProduct, 0)

	if req.Q != "" {
		query += ` AND (
			kode_barang ILIKE '%' || ? || '%'
			OR jumlah ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetTransfersProducts - failed to query transfers_products")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.TransferProduct)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetTransferProduct(ctx context.Context, req *entity.GetTransferProductReq) (*entity.GetTransferProductResp, error) {
	var data = new(entity.TransferProduct)

	query := `
		SELECT id, kode_barang, jumlah
		FROM transfers_products
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTransferProduct - transfer product not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Transfer Product tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTransferProduct - failed to get transfer product")
		return nil, err
	}

	return &entity.GetTransferProductResp{TransferProduct: *data}, nil
}

func (r *masterRepo) CreateTransferProduct(ctx context.Context, req *entity.CreateTransferProductReq) (*entity.CreateTransferProductResp, error) {
	queryInsert := `
		INSERT INTO transfers_products (id, kode_barang, jumlah)
		VALUES (?, ?, ?)`

	queryUpdateStock := `
		UPDATE products
		SET jumlah = jumlah - ?
		WHERE kode = ?`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateTransferProductResp)
	)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::CreateTransferProduct - failed to start transaction")
		return nil, err
	}
	defer func() {
		_ = tx.Rollback() // Rollback akan otomatis dibatalkan jika commit sukses
	}()

	// Insert ke tabel transfers_products
	if _, err := tx.ExecContext(ctx, tx.Rebind(queryInsert), Id, req.KodeBarang, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateTransferProduct - failed to insert transfer record")
		return nil, err
	}

	// Update stok di tabel products
	if _, err := tx.ExecContext(ctx, tx.Rebind(queryUpdateStock), req.Jumlah, req.KodeBarang); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateTransferProduct - failed to update product stock")
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("repo::CreateTransferProduct - failed to commit transaction")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateTransferProduct(ctx context.Context, req *entity.UpdateTransferProductReq) error {
	query := `
		UPDATE transfers_products
		SET kode_barang = ?, jumlah = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.KodeBarang, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateTransferProduct - failed to update transfer product")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteTransferProduct(ctx context.Context, req *entity.DeleteTransferProductReq) error {
	query := `
		UPDATE transfers_products
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteTransferProduct - failed to soft delete transfer product")
		return err
	}
	return nil
}
