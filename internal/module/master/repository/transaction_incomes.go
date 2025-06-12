package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetTransactionIncomes(ctx context.Context, req *entity.GetTransactionIncomesReq) (*entity.GetTransactionIncomesResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.TransactionIncome
	}

	var (
		resp  = new(entity.GetTransactionIncomesResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data,
				   t.id, b.kategori AS kode_document_bc ,b.no_document AS no_document, b.tanggal AS tgl_document_bc,  t.no_kontrak
			FROM transaction_incomes t
			JOIN contracts c ON t.no_kontrak = c.no_kontrak
			JOIN bc_documents b ON t.no_document = b.no_document
			WHERE t.deleted_at IS NULL`
	)

	resp.Items = make([]entity.TransactionIncome, 0)

	if req.Q != "" {
		query += ` AND (
			no_kontrak ILIKE '%' || ? || '%' OR
			kode_barang ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetTransactionIncomes - failed to query transaction incomes")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.TransactionIncome)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetTransactionIncome(ctx context.Context, req *entity.GetTransactionIncomeReq) (*entity.GetTransactionIncomeResp, error) {
	var data = new(entity.TransactionIncome)

	query := `
		SELECT id, no_kontrak, kode_barang, jumlah
		FROM transaction_incomes
		WHERE id = ? AND deleted_at IS NULL`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTransactionIncome - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Transaksi pemasukan tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTransactionIncome - failed to get transaction income")
		return nil, err
	}

	return &entity.GetTransactionIncomeResp{TransactionIncome: *data}, nil
}

func (r *masterRepo) CreateTransactionIncome(ctx context.Context, req *entity.CreateTransactionIncomeReq) (*entity.CreateTransactionIncomeResp, error) {
	query := `
		INSERT INTO transaction_incomes (id, no_kontrak, no_document)
		VALUES (?, ?, ?)`

	id := ulid.Make().String()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("repo::CreateTransactionIncome - failed to begin transaction")
		return nil, err
	}
	defer tx.Rollback()

	// Insert ke transaction_incomes
	if _, err := tx.ExecContext(ctx, tx.Rebind(query), id, req.NoKontrak, req.NoDocumentBc); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateTransactionIncome - failed to insert")
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("repo::CreateTransactionIncome - failed to commit transaction")
		return nil, err
	}

	return &entity.CreateTransactionIncomeResp{Id: id}, nil
}

func (r *masterRepo) UpdateTransactionIncome(ctx context.Context, req *entity.UpdateTransactionIncomeReq) error {
	query := `
		UPDATE transaction_incomes
		SET no_kontrak = ?, kode_barang = ?, jumlah = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.Jumlah, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateTransactionIncome - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteTransactionIncome(ctx context.Context, req *entity.DeleteTransactionIncomeReq) error {
	query := `
		UPDATE transaction_incomes
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteTransactionIncome - failed to soft delete")
		return err
	}

	return nil
}
