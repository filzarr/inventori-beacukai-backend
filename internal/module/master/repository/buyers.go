package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetBuyers(ctx context.Context, req *entity.GetBuyersReq) (*entity.GetBuyersResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Buyer
	}

	var (
		resp  = new(entity.GetBuyersResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 2)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
				id,
				name,
				alamat,
				npwp
			FROM buyers
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.Buyer, 0)

	if req.Q != "" {
		query += ` AND (
			name ILIKE '%' || ? || '%'
			OR alamat ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetBuyers - failed to query buyers")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Buyer)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetBuyer(ctx context.Context, req *entity.GetBuyerReq) (*entity.GetBuyerResp, error) {
	var data = new(entity.Buyer)

	query := `
		SELECT id, name, alamat, npwp
		FROM buyers
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetBuyer - buyer not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Buyer tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetBuyer - failed to get buyer")
		return nil, err
	}

	return &entity.GetBuyerResp{Buyer: *data}, nil
}

func (r *masterRepo) CreateBuyer(ctx context.Context, req *entity.CreateBuyerReq) (*entity.CreateBuyerResp, error) {
	query := `
		INSERT INTO buyers (id, name, alamat, npwp)
		VALUES (?, ?, ?, ?)`

	Id := ulid.Make().String()

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.Name, req.Alamat, req.Npwp); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateBuyer - failed to create buyer")
		return nil, err
	}

	return &entity.CreateBuyerResp{Id: Id}, nil
}

func (r *masterRepo) UpdateBuyer(ctx context.Context, req *entity.UpdateBuyerReq) error {
	query := `
		UPDATE buyers
		SET name = ?, alamat = ?, npwp = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Name, req.Alamat, req.Npwp, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateBuyer - failed to update buyer")
		return err
	}
	return nil
}

func (r *masterRepo) DeleteBuyer(ctx context.Context, req *entity.DeleteBuyerReq) error {
	query := `
		UPDATE buyers
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteBuyer - failed to soft delete buyer")
		return err
	}
	return nil
}
