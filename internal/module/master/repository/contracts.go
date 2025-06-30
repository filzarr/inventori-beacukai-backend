package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetContracts(ctx context.Context, req *entity.GetContractsReq) (*entity.GetContractsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Contract
	}

	var (
		resp  = new(entity.GetContractsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
			       c.id, c.kategori, c.no_kontrak, c.no_document, s.name AS nama_pemasok, s.alamat AS alamat_pemasok, c.tanggal
			FROM contracts c
			JOIN supliers s ON c.supliers_id = s.id
			WHERE c.deleted_at IS NULL`
	)
	if req.Document {
		query += ` AND c.no_document IS NULL`

	}
	if req.Q != "" {
		query += ` AND c.no_kontrak ILIKE '%' || ? || '%'`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetContracts - failed to query contracts")
		return nil, err
	}

	resp.Items = make([]entity.Contract, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Contract)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetContract(ctx context.Context, req *entity.GetContractReq) (*entity.GetContractResp, error) {
	var resp = new(entity.GetContractResp)
	var data = new(entity.Contract)

	query := `SELECT COUNT(*) OVER() AS total_data,
			       c.id, c.kategori, c.no_kontrak, s.name, s.alamat, c.tanggal
			FROM contracts c
			JOIN supliers s ON c.supliers_id = s.id
			WHERE c.deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errmsg.NewCustomErrors(404).SetMessage("Kontrak tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetContract - failed to get contract")
		return nil, err
	}

	resp.Contract = *data
	return resp, nil
}

func (r *masterRepo) CreateContract(ctx context.Context, req *entity.CreateContractReq) (*entity.CreateContractResp, error) {
	query := `INSERT INTO contracts (id, no_kontrak, kategori_kontrak, supliers_id, tanggal) VALUES (?, ?, ?, ?, ?)`
	var resp = new(entity.CreateContractResp)
	id := ulid.Make().String()

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), id, req.NoKontrak, req.Kategori, req.SupliersId, req.Tanggal); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateContract - failed to create contract")
		return nil, err
	}

	resp.Id = id
	return resp, nil
}

func (r *masterRepo) UpdateContract(ctx context.Context, req *entity.UpdateContractReq) error {
	query := `UPDATE contracts SET no_kontrak = ?, supliers_id = ?, kategori = ?, tanggal = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.SupliersId, req.Kategori, req.Tanggal, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateContract - failed to update contract")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteContract(ctx context.Context, req *entity.DeleteContractReq) error {
	query := `UPDATE contracts SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteContract - failed to delete contract")
		return err
	}

	return nil
}

func (r *masterRepo) UpdateContractDocument(ctx context.Context, req *entity.UpdateContractDocumentReq) error {
	query := `UPDATE contracts SET no_document = ?, updated_at = NOW() WHERE no_kontrak = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoDocumentBc, req.NoKontrak); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateDocumentContract - failed to update contract")
		return err
	}

	return nil
}

func (r *masterRepo) GetTransactions(ctx context.Context, req *entity.GetTransactionsReq) (*entity.GetTransactionsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.Transaction
	}

	var (
		resp  = new(entity.GetTransactionsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
			       c.id, bc.kategori AS kode_document, bc.no_document AS no_document, bc.tanggal AS tanggal_document, c.no_kontrak
			FROM contracts c
			JOIN bc_documents bc ON c.no_document = bc.no_document
			WHERE c.deleted_at IS NULL AND c.no_document IS NOT NULL`
	)
	if req.Q != "" {
		query += ` AND c.no_kontrak ILIKE '%' || ? || '%'`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetTransactions - failed to query contracts")
		return nil, err
	}

	resp.Items = make([]entity.Transaction, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.Transaction)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}
