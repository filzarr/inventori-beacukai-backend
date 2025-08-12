package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetProductsMovement(ctx context.Context, req *entity.GetProductsMovementReq) (*entity.GetProductsMovementResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ProductsMovement
	}

	resp := new(entity.GetProductsMovementResp)
	data := []dao{}
	args := []any{}
	query := `
		SELECT COUNT(*) OVER() AS total_data, pm.id, pm.kode_barang, pm.jumlah, pm.status_perpindahan AS status, p.nama AS nama_barang, p.satuan, w.nama AS gudang_pemohon
		FROM products_movement pm
		JOIN products p ON pm.kode_barang = p.kode
		JOIN warehouses wf ON pm.warehouses_from = wf.kode
		JOIN warehouses w ON pm.warehouses_to = w.kode
		WHERE pm.deleted_at IS NULL
    `

	if req.Q != "" {
		query += ` AND pm.kode_barang ILIKE '%' || ? || '%'`
		args = append(args, req.Q)
	}
	if req.Status != "" {
		query += ` AND pm.status_perpindahan = ?`
		args = append(args, req.Status)
	}
	if req.GudangPemohon != "" {
		query += ` AND wf.kategori = ?`
		args = append(args, req.GudangPemohon)
	}
	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetProductsMovement - query failed")
		return nil, err
	}

	resp.Items = make([]entity.ProductsMovement, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.ProductsMovement)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetProductsMovementByID(ctx context.Context, req *entity.GetProductsMovementReqID) (*entity.GetProductsMovementRespID, error) {
	resp := new(entity.GetProductsMovementRespID)
	data := new(entity.ProductsMovement)

	query := `
        SELECT id, kode_barang, jumlah
        FROM products_movement
        WHERE id = ? AND deleted_at IS NULL
    `
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetProductsMovementByID - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Record tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetProductsMovementByID - get failed")
		return nil, err
	}
	resp.ProductsMovement = *data
	return resp, nil
}

func (r *masterRepo) CreateProductsMovement(ctx context.Context, req *entity.CreateProductsMovementReq) (*entity.CreateProductsMovementResp, error) {

	newID := ulid.Make().String()
	query := `
		INSERT INTO products_movement (id, kode_barang,no_kontrak, jumlah, warehouses_from, warehouses_to)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	query = r.db.Rebind(query)

	_, err := r.db.ExecContext(ctx, query, newID, req.KodeBarang, req.NoKontrak, req.Jumlah, req.WarehouseFrom, req.WarehouseTo)
	if err != nil {
		log.Error().
			Err(err).
			Any("req", req).
			Msg("repo::CreateProductsMovement - insert failed")
		return nil, err
	}

	// Return response
	return &entity.CreateProductsMovementResp{
		Id: newID,
	}, nil
}

func (r *masterRepo) UpdateProductsMovement(ctx context.Context, req *entity.UpdateProductsMovementReq) error {
	query := `
        UPDATE products_movement
        SET kode_barang = ?, jumlah = ?, updated_at = NOW()
        WHERE id = ? AND deleted_at IS NULL
    `
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.KodeBarang, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateProductsMovement - update failed")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteProductsMovement(ctx context.Context, req *entity.DeleteProductsMovementReq) error {
	query := `UPDATE products_movement SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteProductsMovement - delete failed")
		return err
	}
	return nil
}

func (r *masterRepo) UpdateStatusProductsMovement(ctx context.Context, req *entity.UpdateStatusProductsMoveMentReq) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to begin tx")
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 1. Ambil data pergerakan produk
	var kodeBarang string
	var warehouse_to string
	var warehouse_from string
	var jumlah int
	selectMovementQuery := `
		SELECT kode_barang, jumlah, warehouses_from, warehouses_to 
		FROM products_movement 
		WHERE id = ? AND deleted_at IS NULL
	`
	movementData := struct {
		KodeBarang    *string `db:"kode_barang"`
		Jumlah        *int    `db:"jumlah"`
		WarehouseTo   *string `db:"warehouses_to"`
		WarehouseFrom *string `db:"warehouses_from"`
	}{}
	if err = tx.GetContext(ctx, &movementData, tx.Rebind(selectMovementQuery), req.Id); err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to get movement data")
		return err
	}

	if movementData.KodeBarang == nil || movementData.Jumlah == nil {
		return fmt.Errorf("invalid movement data: kode_barang or jumlah is nil")
	}
	kodeBarang = *movementData.KodeBarang
	jumlah = *movementData.Jumlah
	warehouse_to = *movementData.WarehouseTo
	warehouse_from = *movementData.WarehouseFrom

	log.Debug().Str("kode_barang", kodeBarang).Int("jumlah", jumlah).Msg("movement data")

	var productExists bool
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE kode = ? AND deleted_at IS NULL)`
	if err = tx.GetContext(ctx, &productExists, tx.Rebind(query), kodeBarang); err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to check product existence")
		return err
	}
	if !productExists {
		return fmt.Errorf("produk dengan kode_barang '%s' tidak ditemukan di tabel products", kodeBarang)
	}

	updateStatusQuery := `
		UPDATE products_movement
		SET status_perpindahan = 'Diterima', updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND deleted_at IS NULL
	`
	if _, err = tx.ExecContext(ctx, tx.Rebind(updateStatusQuery), req.Id); err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to update status")
		return err
	}

	updateProductStokQuery := `
		UPDATE warehouses_stocks 
		SET jumlah = jumlah - ?, updated_at = CURRENT_TIMESTAMP
		WHERE kode_barang = ? AND warehouse_kode = ? AND deleted_at IS NULL
	`
	if _, err = tx.ExecContext(ctx, tx.Rebind(updateProductStokQuery), jumlah, kodeBarang, warehouse_from); err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to update product stock")
		return err
	}

	var exists bool
	checkProductionQuery := `
		SELECT EXISTS(
			SELECT 1 FROM warehouses_stocks WHERE kode_barang = ? AND warehouse_kode = ? AND deleted_at IS NULL
		)
	`
	if err = tx.GetContext(ctx, &exists, tx.Rebind(checkProductionQuery), kodeBarang, warehouse_to); err != nil {
		log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to check production existence")
		return err
	}

	if exists {
		updateProductionQuery := `
			UPDATE warehouses_stocks
			SET jumlah = jumlah + ?, updated_at = CURRENT_TIMESTAMP
			WHERE kode_barang = ? AND warehouse_kode = ? AND  deleted_at IS NULL
		`
		if _, err = tx.ExecContext(ctx, tx.Rebind(updateProductionQuery), jumlah, kodeBarang, warehouse_to); err != nil {
			log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to update production")
			return err
		}
	} else {
		insertProductionQuery := `
			INSERT INTO warehouses_stocks (id, kode_barang, jumlah, warehouse_kode, created_at, updated_at)
			VALUES (?, ?, ?, ? ,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		`
		newID := ulid.Make().String()
		if _, err = tx.ExecContext(ctx, tx.Rebind(insertProductionQuery), newID, kodeBarang, jumlah, warehouse_to); err != nil {
			log.Error().Err(err).Msg("repo::UpdateStatusProductsMovement - failed to insert production")
			return err
		}
	}

	return nil
}
