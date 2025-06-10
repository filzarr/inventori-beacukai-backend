package repository

import (
	"context"
	"database/sql"

	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetIncomeInventories(ctx context.Context, req *entity.GetIncomeInventoriesReq) (*entity.GetIncomeInventoriesResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.IncomeInventory
	}

	var (
		resp = new(entity.GetIncomeInventoriesResp)
		data = make([]dao, 0)
		args = make([]any, 0, 3)
	)
	resp.Items = make([]entity.IncomeInventory, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			id,
			no_kontrak,
			kategori_barang,
			created_at,
			updated_at
		FROM
			income_inventories
		WHERE
			deleted_at IS NULL
	`

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventories - failed to query income_inventories")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.IncomeInventory)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetIncomeInventory(ctx context.Context, req *entity.GetIncomeInventoryReq) (*entity.GetIncomeInventoryResp, error) {
	var (
		resp = new(entity.GetIncomeInventoryResp)
		data = new(entity.IncomeInventory)
	)

	query := `
		SELECT
			id,
			no_kontrak,
			kategori_barang,
			created_at,
			updated_at
		FROM
			income_inventories
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetIncomeInventory - income_inventory not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetIncomeInventory - failed to get income_inventory")
		return nil, err
	}

	resp.IncomeInventory = *data

	return resp, nil
}

func (r *masterRepo) CreateIncomeInventory(ctx context.Context, req *entity.CreateIncomeInventoryReq) (*entity.CreateIncomeInventoryResp, error) {
	query := `
		INSERT INTO income_inventories (
			id,
			no_kontrak,
			kategori_barang
		) VALUES (?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateIncomeInventoryResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.NoKontrak, req.KategoriBarang); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateIncomeInventory - failed to create income_inventory")
		return nil, err
	}

	resp.Id = Id

	return resp, nil
}

func (r *masterRepo) UpdateIncomeInventory(ctx context.Context, req *entity.UpdateIncomeInventoryReq) error {
	query := `
		UPDATE income_inventories
		SET
			no_kontrak = ?,
			kategori_barang = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.KategoriBarang, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateIncomeInventory - failed to update income_inventory")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteIncomeInventory(ctx context.Context, req *entity.DeleteIncomeInventoryReq) error {
	query := `
		UPDATE income_inventories
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteIncomeInventory - failed to delete income_inventory")
		return err
	}

	return nil
}
