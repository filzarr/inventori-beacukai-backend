package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetProductions(ctx context.Context, req *entity.GetProductionsReq) (*entity.GetProductionsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Production
	}

	var (
		resp  = new(entity.GetProductionsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
            SELECT 
              COUNT(*) OVER() AS total_data,
              pc.id, p.kode as kode_barang, pc.jumlah, p.satuan, p.nama AS nama_barang
            FROM productions pc
			JOIN products p ON pc.kode_barang = p.kode
            WHERE pc.deleted_at IS NULL
        `
	)

	resp.Items = make([]entity.Production, 0)

	if req.Q != "" {
		query += ` AND kode_barang ILIKE '%' || ? || '%'`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetProductions - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Production)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetProduction(ctx context.Context, req *entity.GetProductionReq) (*entity.GetProductionResp, error) {
	var (
		resp = new(entity.GetProductionResp)
		data = new(entity.Production)
	)

	query := `
        SELECT id, kode_barang, jumlah
        FROM productions
        WHERE id = ? AND deleted_at IS NULL
    `
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetProduction - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data produksi tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetProduction - failed to get")
		return nil, err
	}

	resp.Production = *data
	return resp, nil
}

func (r *masterRepo) CreateProduction(ctx context.Context, req *entity.CreateProductionReq) (*entity.CreateProductionResp, error) {
	query := `
        INSERT INTO productions (id, kode_barang, jumlah)
        VALUES (?, ?, ?)
    `
	id := ulid.Make().String()
	resp := new(entity.CreateProductionResp)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), id, req.KodeBarang, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateProduction - failed to insert")
		return nil, err
	}
	resp.Id = id
	return resp, nil
}

func (r *masterRepo) UpdateProduction(ctx context.Context, req *entity.UpdateProductionReq) error {
	query := `
        UPDATE productions
        SET kode_barang = ?, jumlah = ?, updated_at = NOW()
        WHERE id = ? AND deleted_at IS NULL
    `
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.KodeBarang, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateProduction - failed to update")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteProduction(ctx context.Context, req *entity.DeleteProductionReq) error {
	query := `
        UPDATE productions
        SET deleted_at = NOW()
        WHERE id = ? AND deleted_at IS NULL
    `
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteProduction - failed to delete")
		return err
	}
	return nil
}
