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
		SELECT
			COUNT(*) OVER() AS total_data,
			cp.id,
			cp.no_kontrak AS no_kontrak,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.saldo_awal AS saldo_awal,
			cp.jumlah AS jumlah_kontrak,
			cp.nilai_barang_fog,
			cp.nilai_barang_rp,
			COALESCE(SUM(iip.jumlah), 0) AS jumlah_realisasi
		FROM contract_products cp
		JOIN products p ON cp.kode_barang = p.kode
		JOIN contracts c ON c.no_kontrak = cp.no_kontrak
		LEFT JOIN income_inventories_products iip 
			ON iip.kode_barang = cp.kode_barang 
			AND iip.no_kontrak = cp.no_kontrak
		WHERE cp.deleted_at IS NULL AND c.deleted_at IS NULL
	`

	if req.Q != "" {
		query += ` AND (
			p.kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}
	if req.Kategori != "" {
		query += ` AND c.kategori_kontrak = ?`
		args = append(args, req.Kategori)
	}
	if req.Full {
		query += ` AND COALESCE(iip.jumlah, 0) < cp.jumlah`
	}

	query += `
		GROUP BY
			cp.id,
			cp.no_kontrak,
			p.kode,
			p.nama,
			p.saldo_awal,
			cp.jumlah
		ORDER BY p.kode`

	query += ` LIMIT ? OFFSET ?;`
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
	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateIncomeInventoryProductResp)
	)
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to start transaction")
		return nil, err
	}
	existingProductQuery := `
		SELECT COUNT(*) FROM products WHERE kode = ? AND deleted_at IS NULL
	`
	var count int
	if err := tx.GetContext(ctx, &count, tx.Rebind(existingProductQuery), req.KodeBarang); err != nil {
		log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to check existing product")
		tx.Rollback()
		return nil, err
	}
	if count == 0 {
		log.Warn().Msg("repo::CreateIncomeInventoryProduct - product not found")
		tx.Rollback()
		return nil, errmsg.NewCustomErrors(404).SetMessage("Produk tidak ditemukan")
	}
	query := `
		INSERT INTO income_inventories_products (
			id,
			no_kontrak,
			kode_barang,
			nomor_document_bc,
			warehouse_location,
			driver,
			license_plate,
			bruto_weight,
			netto_weight,
			empty_weight,
			starting_time,
			ending_time,
			stok_awal,
			jumlah,
			tanggal
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	// Handle NULL warehouse_location
	var warehouseLocation interface{}
	if req.WarehouseLocation != nil {
		warehouseLocation = *req.WarehouseLocation
	} else {
		warehouseLocation = nil
	}

	_, err = tx.ExecContext(ctx, tx.Rebind(query), Id,
		req.NoKontrak,
		req.KodeBarang,
		req.NomorDocumentBc,
		warehouseLocation,
		req.Driver,
		req.LicensePlate,
		req.BrutoWeight,
		req.NettoWeight,
		req.EmptyWeight,
		req.StartingTime,
		req.EndingTime,
		req.SaldoAwal,
		req.Jumlah,
		req.Tanggal)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateIncomeInventoryProduct - failed to insert")
		tx.Rollback()
		return nil, err
	}

	// Only update warehouse stocks if warehouse_location is provided
	if req.WarehouseLocation != nil {
		var wsCount int
		checkWSQuery := `
			SELECT COUNT(*) 
			FROM warehouses_stocks 
			WHERE warehouse_kode = ? 
			  AND kode_barang = ? 
			  AND deleted_at IS NULL
		`
		if err := tx.GetContext(ctx, &wsCount, tx.Rebind(checkWSQuery), *req.WarehouseLocation, req.KodeBarang); err != nil {
			log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to check warehouses_stocks")
			tx.Rollback()
			return nil, err
		}

		if wsCount == 0 {
			wsID := ulid.Make().String()
			insertWSQuery := `
				INSERT INTO warehouses_stocks (
					id,
					warehouse_kode,
					kode_barang,
					jumlah
				) VALUES (?, ?, ?, ?)
			`
			if _, err := tx.ExecContext(ctx, tx.Rebind(insertWSQuery),
				wsID,
				*req.WarehouseLocation,
				req.KodeBarang,
				req.Jumlah,
			); err != nil {
				log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to insert warehouses_stocks")
				tx.Rollback()
				return nil, err
			}
		} else {
			updateWSQuery := `
				UPDATE warehouses_stocks
				SET jumlah = jumlah + ?, updated_at = CURRENT_TIMESTAMP
				WHERE warehouse_kode = ? 
				  AND kode_barang = ?
			`
			if _, err := tx.ExecContext(ctx, tx.Rebind(updateWSQuery),
				req.Jumlah,
				*req.WarehouseLocation,
				req.KodeBarang,
			); err != nil {
				log.Error().Err(err).Msg("repo::CreateIncomeInventoryProduct - failed to update warehouses_stocks")
				tx.Rollback()
				return nil, err
			}
		}
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

func (r *masterRepo) GetIncomeInventoryProductsByContract(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractReq) (*entity.GetIncomeInventoryProductsByContractResp, error) {
	var (
		resp = new(entity.GetIncomeInventoryProductsByContractResp)
		data = make([]entity.IncomeInventoryProductByContract, 0)
	)

	query := `
		SELECT
			cp.id,
			cp.no_kontrak AS no_kontrak,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.saldo_awal AS saldo_awal,
			cp.jumlah AS jumlah_kontrak,
			cp.nilai_barang_fog,
			cp.nilai_barang_rp,
			COALESCE(SUM(iip.jumlah), 0) AS jumlah_realisasi
		FROM contract_products cp
		JOIN products p ON cp.kode_barang = p.kode
		LEFT JOIN income_inventories_products iip 
			ON iip.kode_barang = cp.kode_barang 
			AND iip.no_kontrak = cp.no_kontrak
		WHERE cp.no_kontrak = ?
		AND cp.deleted_at IS NULL
		GROUP BY
			cp.id,
			cp.no_kontrak,
			p.kode,
			p.nama,
			p.saldo_awal,
			cp.jumlah
		ORDER BY p.kode;
	`

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), req.NoKontrak); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventoryProductByContract - failed to query")
		return nil, err
	}

	resp.Items = data
	return resp, nil
}

func (r *masterRepo) GetIncomeInventoryProductsByContractAndKode(ctx context.Context, req *entity.GetIncomeInventoryProductsByContractAndKodeReq) (*entity.GetIncomeInventoryProductsByContractAndKodeResp, error) {
	var (
		resp = new(entity.GetIncomeInventoryProductsByContractAndKodeResp)
		data = make([]entity.IncomeInventoryProductsByContractAndKode, 0)
	)

	query := `
		SELECT
			iip.id,
			iip.no_kontrak AS no_kontrak,
			iip.kode_barang AS kode_barang,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			w.nama AS lokasi_penyimpanan,
			iip.driver AS driver,
			iip.jumlah AS jumlah_masuk,
			iip.license_plate AS license_plate,
			iip.bruto_weight AS bruto_weight,
			iip.netto_weight AS netto_weight,
			iip.starting_time AS jam_masuk,
			iip.ending_time AS jam_keluar,
			iip.tanggal AS tanggal
		FROM income_inventories_products iip
			JOIN products p ON iip.kode_barang = p.kode
			JOIN warehouses w ON iip.warehouse_location = w.kode
		WHERE iip.no_kontrak = ?
			AND iip.kode_barang = ?
			AND iip.deleted_at IS NULL
	`
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), req.NoKontrak, req.KodeBarang); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventoryProductByContractAndKode - failed to query")
		return nil, err
	}

	resp.Items = data
	return resp, nil
}
