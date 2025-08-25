package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetContractsBc(ctx context.Context, req *entity.GetContractsBcReq) (*entity.GetContractsBcResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ContractsBc
	}

	var (
		resp  = new(entity.GetContractsBcResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, bc.id, bc.no_kontrak, bc.kode_document_bc, bc.tanggal_document_bc, bc.nomor_document_bc
			FROM contracts_bc bc
			LEFT JOIN contracts c ON bc.no_kontrak = c.no_kontrak
			WHERE bc.deleted_at IS NULL`
	)

	resp.Items = make([]entity.ContractsBc, 0)

	if req.Q != "" {
		query += ` AND (
			bc.no_kontrak ILIKE '%' || ? || '%' OR
			bc.kode_document_bc ILIKE '%' || ? || '%' OR
			bc.kode_mata_uang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q, req.Q)
	}

	if req.NoKontrak != "" {
		query += ` AND bc.no_kontrak = ?`
		args = append(args, req.NoKontrak)
	}

	if req.KategoriKontrak != "" {
		query += ` AND c.kategori_kontrak = ?`
		args = append(args, req.KategoriKontrak)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetContractsBc - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.ContractsBc)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetContractBc(ctx context.Context, req *entity.GetContractBcReq) (*entity.GetContractBcResp, error) {
	var (
		resp = new(entity.GetContractBcResp)
		data = new(entity.ContractsBc)
	)

	query := `
		SELECT id, no_kontrak, kode_document_bc, nomor_document_bc, tanggal_document_bc
		FROM contracts_bc
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetContractBc - not found")
			return nil, errmg.NewCustomErrors(404).SetMessage("Contract BC tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetContractBc - failed to get")
		return nil, err
	}

	resp.ContractsBc = *data
	return resp, nil
}

func (r *masterRepo) CreateContractBc(ctx context.Context, req *entity.CreateContractBcReq) (*entity.CreateContractBcResp, error) {
	query := `INSERT INTO contracts_bc (id, no_kontrak, kode_document_bc, nomor_document_bc, tanggal_document_bc) VALUES (?, ?, ?, ?, ?)`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateContractBcResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.NoKontrak, req.KodeDocumentBc, req.NomorDocumentBc, req.TanggalDocumentBc); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateContractBc - failed to create")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateContractBc(ctx context.Context, req *entity.UpdateContractBcReq) error {
	query := `UPDATE contracts_bc SET no_kontrak = ?, kode_document_bc = ?, tanggal_document_bc = ?, nomor_document_bc = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.KodeDocumentBc, req.TanggalDocumentBc, req.NomorDocumentBc, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateContractBc - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteContractBc(ctx context.Context, req *entity.DeleteContractBcReq) error {
	query := `UPDATE contracts_bc SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteContractBc - failed to delete")
		return err
	}

	return nil
}
