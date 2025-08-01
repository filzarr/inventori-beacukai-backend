package repository

import (
	"context"
	"database/sql"

	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"inventori-beacukai-backend/internal/module/master/entity"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetIncomeInventoryProducts(ctx context.Context, req *entity.GetIncomeInventoryProductsReq) (*entity.GetIncomeInventoryProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.IncomeInventoryProduct
	}

	var (
		resp = new(entity.GetIncomeInventoryProductsResp)
		data = make([]dao, 0)
		args = make([]any, 0)
	)
	resp.Items = make([]entity.IncomeInventoryProduct, 0)

	query := `
		WITH total_masuk AS (
			SELECT
				no_kontrak,
				kode_barang,
				SUM(jumlah) AS jumlah_masuk
			FROM income_inventories_products
			WHERE deleted_at IS NULL
			GROUP BY no_kontrak, kode_barang
		)
		SELECT
			COUNT(*) OVER() AS total_data,
			cp.no_kontrak,
			cp.kode_barang,
			p.nama AS nama_barang,
			COALESCE(tm.jumlah_masuk, 0) AS jumlah_masuk,
			cp.jumlah AS jumlah_kontrak,
			MIN(iip.id) AS id  -- ambil salah satu id untuk representasi
		FROM contract_products cp
		JOIN products p ON cp.kode_barang = p.kode
		LEFT JOIN total_masuk tm ON tm.no_kontrak = cp.no_kontrak AND tm.kode_barang = cp.kode_barang
		LEFT JOIN income_inventories_products iip ON iip.no_kontrak = cp.no_kontrak AND iip.kode_barang = cp.kode_barang AND iip.deleted_at IS NULL
		WHERE cp.deleted_at IS NULL
	`

	if req.Q != "" {
		query += ` AND (
			iip.kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}
	if req.Full {
		query += ` AND COALESCE(tm.jumlah_masuk, 0) < cp.jumlah`
	}

	query += `
		GROUP BY cp.no_kontrak, cp.kode_barang, p.nama, cp.jumlah, tm.jumlah_masuk
		ORDER BY cp.no_kontrak, cp.kode_barang`

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventoryProducts - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.IncomeInventoryProduct)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetIncomeInventoryProduct(ctx context.Context, req *entity.GetIncomeInventoryProductReq) (*entity.GetIncomeInventoryProductResp, error) {
	var (
		resp = new(entity.GetIncomeInventoryProductResp)
		data = new(entity.IncomeInventoryProduct)
	)

	query := `
		SELECT
			id,
			id_inventories,
			kode_barang,
			jumlah,
			lokasi,
			created_at,
			updated_at,
			deleted_at
		FROM
			income_inventories_products
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetIncomeInventoryProduct - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventoryProduct - failed to get")
		return nil, err
	}

	resp.IncomeInventoryProduct = *data
	return resp, nil
}

func (r *masterRepo) CreateIncomeInventoryProduct(ctx context.Context, req *entity.CreateIncomeInventoryProductReq) (*entity.CreateIncomeInventoryProductResp, error) {
	query := `
		INSERT INTO income_inventories_products (
			id,
			no_kontrak,
			kode_barang,
			lokasi,
			stok_awal,
			jumlah
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateIncomeInventoryProductResp)
	)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to start transaction")
		return nil, err
	}
	defer tx.Rollback()

	// Insert ke income_inventories_products
	if _, err := tx.ExecContext(ctx, tx.Rebind(query), Id, req.NoKontrak, req.KodeBarang, req.Lokasi, req.SaldoAwal, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateIncomeInventoryProduct - failed to insert")
		return nil, err
	}

	// Update stok di tabel products
	updateStockQuery := `
		UPDATE products
		SET jumlah = jumlah + ?
		WHERE kode = ?
	`

	if _, err := tx.ExecContext(ctx, tx.Rebind(updateStockQuery), req.Jumlah, req.KodeBarang); err != nil {
		log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to update product stock")
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to commit transaction")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateIncomeInventoryProduct(ctx context.Context, req *entity.UpdateIncomeInventoryProductReq) error {
	query := `
		UPDATE income_inventories_products
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
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateIncomeInventoryProduct - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteIncomeInventoryProduct(ctx context.Context, req *entity.DeleteIncomeInventoryProductReq) error {
	query := `
		UPDATE income_inventories_products
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteIncomeInventoryProduct - failed to delete")
		return err
	}

	return nil
}
