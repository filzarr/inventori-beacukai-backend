package repository

import (
	"context"
	"database/sql"

	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetBcDocuments(ctx context.Context, req *entity.GetBcDocumentsReq) (*entity.GetBcDocumentsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.BcDocument
	}

	var (
		resp  = new(entity.GetBcDocumentsResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, id, kategori, kode_document 
			FROM bc_documents
			WHERE deleted_at IS NULL`
	)

	resp.Items = make([]entity.BcDocument, 0)

	if req.Q != "" {
		query += ` AND (
			kode_document ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetBcDocuments - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.BcDocument)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetBcDocument(ctx context.Context, req *entity.GetBcDocumentReq) (*entity.GetBcDocumentResp, error) {
	var (
		resp = new(entity.GetBcDocumentResp)
		data = new(entity.BcDocument)
	)

	query := `SELECT id, kategori, kode_document, tanggal FROM bc_documents WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetBcDocument - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Dokumen BC tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetBcDocument - failed to get")
		return nil, err
	}

	resp.BcDocument = *data
	return resp, nil
}

func (r *masterRepo) CreateBcDocument(ctx context.Context, req *entity.CreateBcDocumentReq) (*entity.CreateBcDocumentResp, error) {
	query := `INSERT INTO bc_documents (id, kategori, kode_document) VALUES (?, ?, ?)`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateBcDocumentResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.Kategori, req.KodeDocument); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateBcDocument - failed to create")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateBcDocument(ctx context.Context, req *entity.UpdateBcDocumentReq) error {
	query := `
		UPDATE bc_documents
		SET kategori = ?, kode_document = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Kategori, req.KodeDocument, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateBcDocument - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteBcDocument(ctx context.Context, req *entity.DeleteBcDocumentReq) error {
	query := `UPDATE bc_documents SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteBcDocument - failed to delete")
		return err
	}

	return nil
}
