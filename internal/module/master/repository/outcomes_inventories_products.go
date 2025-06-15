package repository

import (
	"context"
	"database/sql"

	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"inventori-beacukai-backend/internal/module/master/entity"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetOutcomesInventoriesProducts(ctx context.Context, req *entity.GetOutcomesInventoriesProductsReq) (*entity.GetOutcomesInventoriesProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.OutcomesInventoriesProduct
	}

	var (
		resp = new(entity.GetOutcomesInventoriesProductsResp)
		data = make([]dao, 0)
		args = make([]any, 0)
	)
	resp.Items = make([]entity.OutcomesInventoriesProduct, 0)

	query := `
		SELECT DISTINCT ON (iip.id)
		COUNT(*) OVER() AS total_data,
		iip.id,
		iip.kode_barang,
		p.nama AS nama_barang,
		cp.jumlah AS jumlah_kontrak,
		iip.no_kontrak AS no_kontrak,
		iip.jumlah AS jumlah_masuk
		FROM income_inventories_products iip
		JOIN products p ON iip.kode_barang = p.kode
		JOIN contract_products cp ON iip.no_kontrak = cp.no_kontrak
		WHERE iip.deleted_at IS NULL 
	`

	if req.Q != "" {
		query += ` AND (
			iip.kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetOutcomesInventoriesProducts - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.OutcomesInventoriesProduct)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetOutcomesInventoriesProduct(ctx context.Context, req *entity.GetOutcomesInventoriesProductReq) (*entity.GetOutcomesInventoriesProductResp, error) {
	var (
		resp = new(entity.GetOutcomesInventoriesProductResp)
		data = new(entity.OutcomesInventoriesProduct)
	)

	query := `
		SELECT
			id,
			id_inventories,
			kode_barang,
			jumlah,
			created_at,
			updated_at,
			deleted_at
		FROM
			outcomes_inventories_products
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetOutcomesInventoriesProduct - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetOutcomesInventoriesProduct - failed to get")
		return nil, err
	}

	resp.OutcomesInventoriesProduct = *data
	return resp, nil
}

func (r *masterRepo) CreateOutcomesInventoriesProduct(ctx context.Context, req *entity.CreateOutcomesInventoriesProductReq) (*entity.CreateOutcomesInventoriesProductResp, error) {
	query := `
		INSERT INTO outcomes_inventories_products (
			id,
			no_kontrak,
			kode_barang,
			stok_awal,
			jumlah
		) VALUES (?, ?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateOutcomesInventoriesProductResp)
	)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::CreateOutcomesInventoriesProduct - failed to start transaction")
		return nil, err
	}
	defer tx.Rollback()

	// Insert ke income_inventories_products
	if _, err := tx.ExecContext(ctx, tx.Rebind(query), Id, req.NoKontrak, req.KodeBarang, req.SaldoAwal, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateOutcomesInventoriesProduct - failed to insert")
		return nil, err
	}

	// Update stok di tabel products
	updateStockQuery := `
		UPDATE products
		SET jumlah = jumlah - ?
		WHERE kode = ?
	`

	if _, err := tx.ExecContext(ctx, tx.Rebind(updateStockQuery), req.Jumlah, req.KodeBarang); err != nil {
		log.Error().Err(err).Msg("repo::CreateOutcomesInventoriesProduct - failed to update product stock")
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("repo::CreateOutcomesInventoriesProduct - failed to commit transaction")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateOutcomesInventoriesProduct(ctx context.Context, req *entity.UpdateOutcomesInventoriesProductReq) error {
	query := `
		UPDATE outcomes_inventories_products
		SET
			no_kontrak = ?,
			kode_barang = ?,
			jumlah = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.KodeBarang, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateOutcomesInventoriesProduct - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteOutcomesInventoriesProduct(ctx context.Context, req *entity.DeleteOutcomesInventoriesProductReq) error {
	query := `
		UPDATE outcomes_inventories_products
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteOutcomesInventoriesProduct - failed to delete")
		return err
	}

	return nil
}
