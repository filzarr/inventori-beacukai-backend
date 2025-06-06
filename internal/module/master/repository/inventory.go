package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"strings"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetInventories(ctx context.Context, req *entity.GetInventoriesReq) (*entity.GetInventoriesResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Inventory
	}

	var (
		resp    = new(entity.GetInventoriesResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 3)
		filters = []string{"deleted_at IS NULL"}
	)
	resp.Items = make([]entity.Inventory, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			id,
			inventories_status,
			kode_barang,
			nama_barang,
			kategori,
			jumlah
		FROM
			inventories  
	`
	if req.Kategori != "" {
		filters = append(filters, "kategori = ?")
		args = append(args, req.Kategori)
	}
	if req.Q != "" {
		filters = append(filters, "nama_barang ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetInventory - failed to query Inventory")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Inventory)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetInventory(ctx context.Context, req *entity.GetInventoryReq) (*entity.GetInventoryResp, error) {
	var (
		resp = new(entity.GetInventoryResp)
		data = new(entity.Inventory)
	)

	query := `
		SELECT 
			id,
			inventories_status,
			kode_barang,
			nama_barang,
			kategori, 
			jumlah
		FROM
			inventories 
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetInventory - inventory not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Inventory tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetInventory - failed to get Inventory")
		return nil, err
	}

	resp.Inventory = *data

	return resp, nil
}

func (r *masterRepo) CreateInventory(ctx context.Context, req *entity.CreateInventoryReq) (*entity.CreateInventoryResp, error) {
	query := `
		INSERT INTO inventories (
			id,
			inventories_status,
			kode_barang,
			nama_barang,
			kategori,
			jumlah
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateInventoryResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		Id, req.InventoriesStatus, req.KodeBarang, req.NamaBarang, req.Kategori, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateInventory - failed to create inventory")
		return nil, err
	}

	resp.Id = Id

	return resp, nil
}

func (r *masterRepo) UpdateInventory(ctx context.Context, req *entity.UpdateInventoryReq) error {
	query := `
		UPDATE inventories
		SET
			inventories_status = ?,
			kode_barang = ?,
			nama_barang = ?,
			kategori = ?, 
			jumlah = ?
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		req.InventoriesStatus, req.KodeBarang, req.NamaBarang, req.Kategori, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateInventory - failed to update inventory")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteInventory(ctx context.Context, req *entity.DeleteInventoryReq) error {
	query := `
		UPDATE inventories
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteInventory - failed to delete inventory")
		return err
	}

	return nil
}

func (r *masterRepo) GetInventoriesBahanBaku(ctx context.Context, req *entity.GetInventoriesBahanBakuReq) (*entity.GetInventoriesBahanBakuResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Inventory
	}

	var (
		resp    = new(entity.GetInventoriesBahanBakuResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 3)
		filters = []string{"deleted_at IS NULL AND (kategori = 'Bahan Baku' OR kategori = 'Bahan Penolong')"}
	)
	resp.Items = make([]entity.Inventory, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			id,
			inventories_status,
			kode_barang,
			nama_barang,
			kategori,
			jumlah
		FROM
			inventories  
	`
	if req.Q != "" {
		filters = append(filters, "nama_barang ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetInventory - failed to query Inventory")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Inventory)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}
func joinFilters(filters []string) string {
	return strings.Join(filters, " AND ")
}
