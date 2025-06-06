package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetMutasiBahans(ctx context.Context, req *entity.GetMutasiBahansReq) (*entity.GetMutasiBahansResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.MutasiBahan
	}

	var (
		resp    = new(entity.GetMutasiBahansResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 3)
		filters = []string{"m.deleted_at IS NULL"}
	)
	resp.Items = make([]entity.MutasiBahan, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			m.id,
			i.nama_barang,
			i.id as id_barang, 
			g.nama as nama_gudang,
			m.jumlah,
			m.status
		FROM
			mutasi_bahan_baku_penolong m
		JOIN
			inventories i
			ON m.inventories_id = i.id
		JOIN 
			gudang g
			ON m.gudang_id = g.id
	`
	if req.Gudang != "" {
		filters = append(filters, "m.gudang_id = ?")
		args = append(args, req.Gudang)
	}
	if req.Q != "" {
		filters = append(filters, "i.nama_barang ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetMutasiBahans - failed to query MutasiBahan")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.MutasiBahan)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *masterRepo) GetMutasiBahan(ctx context.Context, req *entity.GetMutasiBahanReq) (*entity.GetMutasiBahanResp, error) {
	var (
		resp = new(entity.GetMutasiBahanResp)
		data = new(entity.MutasiBahan)
	)

	query := `
		SELECT 
			m.id,
			i.nama_barang,
			i.id as id_barang,
			i.kode_barang as kode_barang,
			g.id as gudang_id,
			g.nama as nama_gudang,
			m.jumlah,
			m.status,
			m.saldo_awal,
			m.pemasukan,
			m.pengeluaran,
			m.penyesuaian
		FROM
			mutasi_bahan_baku_penolong m
		JOIN
			inventories i
			ON m.inventories_id = i.id
		JOIN 
			gudang g
			ON m.gudang_id = g.id
		WHERE
			m.id = ? 
			AND m.deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetMutasiBahan - Mutasi Bahan not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Mutasi Bahan tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetMutasiBahan - failed to get Mutasi Bahan")
		return nil, err
	}

	resp.MutasiBahan = *data

	return resp, nil
}

func (r *masterRepo) CreateMutasiBahan(ctx context.Context, req *entity.CreateMutasiBahanReq) (*entity.CreateMutasiBahanResp, error) {
	id := ulid.Make().String()
	resp := new(entity.CreateMutasiBahanResp)

	// INSERT mutasi_bahan_baku_penolong
	insertQuery := `
		INSERT INTO mutasi_bahan_baku_penolong (
			id,
			inventories_id,
			gudang_id,
			jumlah
		) VALUES (?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(insertQuery),
		id, req.Inventories, req.Gudang, req.Jumlah); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateMutasiBahan - failed to insert mutasi bahan")
		return nil, err
	}

	resp.Id = id
	return resp, nil
}

func (r *masterRepo) UpdateMutasiBahan(ctx context.Context, req *entity.UpdateMutasiBahanReq) error {
	query := `
		UPDATE mutasi_bahan_baku_penolong
		SET
			inventories_id = ?,
			gudang_id = ?,
			jumlah = ?, 
			status = ?,
			saldo_awal = ?,
			pemasukan = ?,
			pengeluaran = ?,
			penyesuaian = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		req.Inventories, req.Gudang, req.Jumlah, req.Status, req.SaldoAwal, req.Pemasukan, req.Pengeluaran, req.Penyesuaian, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateMutasiBahan - failed to update mutasibahan")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteMutasiBahan(ctx context.Context, req *entity.DeleteMutasiBahanReq) error {
	query := `
		UPDATE mutasi_bahan_baku_penolong
		SET
			deleted_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteMutasiBahan - failed to delete mutasi bahan")
		return err
	}

	return nil
}

func (r *masterRepo) UpdateStatusMutasiBahan(ctx context.Context, req *entity.UpdateStatusMutasiBahanReq) error {
	query := `
		UPDATE mutasi_bahan_baku_penolong
		SET 
			status = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		req.Status, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateStatusMutasiBahan - failed to update mutasibahan")
		return err
	}

	// UPDATE inventories
	updateQuery := `
		UPDATE inventories
		SET
			jumlah = jumlah - ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(updateQuery),
		req.Jumlah, req.Inventory); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateStatusMutasiBahan - failed to update inventory stok")
		return err
	}

	return nil
}
func (r *masterRepo) UpdateSaldoMutasi(ctx context.Context, req *entity.UpdateSaldoMutasiReq) error {
	query := `
		UPDATE mutasi_bahan_baku_penolong
		SET 
			saldo_awal = ?,
			pemasukan = ?,
			pengeluaran = ?,
			penyesuaian = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		req.SaldoAwal, req.Pemasukan, req.Pengeluaran, req.Penyesuaian, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateSaldoMutasi - failed to update mutasibahan")
		return err
	}

	return nil
}

func (r *masterRepo) GetLaporanMutasiBahan(ctx context.Context, req *entity.GetLaporanMutasiBahanReq) (*entity.GetLaporanMutasiBahanResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.LaporanMutasiBahan
	}

	var (
		resp    = new(entity.GetLaporanMutasiBahanResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 3)
		filters = []string{"m.deleted_at IS NULL AND m.status = 'Diterima'"}
	)
	resp.Items = make([]entity.LaporanMutasiBahan, 0)

	query := `
		SELECT
			COUNT (*) OVER() AS total_data,
			m.id,
			i.kode_barang as kode_barang,
			i.nama_barang as nama_barang,
			m.saldo_awal as saldo_awal,
			m.pemasukan as pemasukan,
			m.penyesuaian as penyesuaian,
			m.jumlah as jumlah,
			m.created_at as created_at
		FROM
			mutasi_bahan_baku_penolong m
		JOIN
			inventories i
			ON m.inventories_id = i.id
		JOIN 
			gudang g
			ON m.gudang_id = g.id
	`
	if req.Gudang != "" {
		filters = append(filters, "m.gudang_id = ?")
		args = append(args, req.Gudang)
	}
	if req.Q != "" {
		filters = append(filters, "i.nama_barang ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)
	log.Debug().Any("args", args).Msg("Parsed Query")
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasiBahan - failed to query MutasiBahan")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.LaporanMutasiBahan)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}
