package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
	"inventori-beacukai-backend/pkg/types"
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetLaporanMutasi(ctx context.Context, req *entity.GetLaporanMutasiReq) (*entity.GetLaporanMutasiResp, error) {
	var (
		data = make([]entity.LaporanMutasi, 0)
		args = make([]any, 0)
	)

	query := `
		SELECT
			p.id,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.saldo_awal,
			COALESCE(SUM(iip.jumlah), 0) AS pemasukan,
			p.jumlah AS stok_opname
		FROM products p
		LEFT JOIN income_inventories_products iip ON p.kode = iip.kode_barang
	`

	// Menyusun kondisi WHERE
	conditions := make([]string, 0)
	if req.Kategori != "" {
		conditions = append(conditions, "p.kategori = ?")
		args = append(args, req.Kategori)
	}
	if req.Q != "" {
		conditions = append(conditions, "(p.kode ILIKE ? OR p.nama ILIKE ?)")
		q := "%" + req.Q + "%"
		args = append(args, q, q)
	}

	// Gabungkan kondisi WHERE jika ada
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Group by semua kolom non-agregat
	query += `
		GROUP BY p.id, p.kode, p.nama, p.saldo_awal, p.jumlah
	`

	// Eksekusi query
	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasi - failed to query")
		return nil, err
	}

	return &entity.GetLaporanMutasiResp{
		Items: data,
		Meta:  types.Meta{TotalData: len(data)},
	}, nil
}

func (r *masterRepo) GetLaporanMutasiPemasukan(ctx context.Context, req *entity.GetLaporanMutasiPemasukanReq) (*entity.GetLaporanMutasiPemasukanResp, error) {
	var (
		data       = make([]entity.LaporanMutasiPemasukan, 0)
		args       = make([]any, 0)
		conditions = make([]string, 0)
	)

	query := `
		SELECT   
			cp.id,
			bc.kategori AS kategori,
			bc.kode_document AS kode_document, 
			c.tanggal_document_bc AS tanggal,
			cp.no_kontrak,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.satuan AS satuan,
			iip.jumlah AS jumlah
		FROM contract_products cp
		JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		JOIN bc_documents bc ON c.kode_document_bc = bc.kode_document
		JOIN products p ON cp.kode_barang = p.kode
		JOIN income_inventories_products iip 
			ON c.no_kontrak = iip.no_kontrak AND cp.kode_barang = iip.kode_barang
	`
	// log.Info().Msg(req.KodeBarang)
	if req.KodeBarang != "" {
		conditions = append(conditions, "p.kode = ?")
		args = append(args, req.KodeBarang)
	}
	if req.Q != "" {
		conditions = append(conditions, "(p.kode ILIKE ? OR p.nama ILIKE ?)")
		q := "%" + req.Q + "%"
		args = append(args, q, q)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasiPemasukan - failed to query")
		return nil, err
	}

	return &entity.GetLaporanMutasiPemasukanResp{
		Items: data,
		Meta:  types.Meta{TotalData: len(data)},
	}, nil
}

func (r *masterRepo) GetLaporanMutasiWip(ctx context.Context, req *entity.GetLaporanMutasiWipReq) (*entity.GetLaporanMutasiWipResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.LaporanMutasiWip
	}

	var (
		resp = new(entity.GetLaporanMutasiWipResp)
		data = make([]dao, 0)
		args = make([]any, 0, 3)
	)

	query := `SELECT COUNT(*) OVER() AS total_data,
				pm.id, p.nama AS nama_barang, p.satuan AS satuan, pm.jumlah
				FROM products_movement pm
				LEFT JOIN products p ON pm.kode_barang = p.kode
				WHERE pm.status_perpindahan = 'Diterima'
				AND pm.deleted_at IS NULL
	`
	if req.Q != "" {
		query += ` AND p.nama ILIKE '%' || ? || '%'
			OR p.kode  ILIKE '%' || ? || '%'
		`
		args = append(args, req.Q, req.Q)
	}

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanWIP - failed to query Laporan WIP")
		return nil, err
	}
	resp.Items = make([]entity.LaporanMutasiWip, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.LaporanMutasiWip)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetLaporanMutasiJenisDokumen(ctx context.Context, req *entity.GetLaporanMutasiJenisDokumenReq) (*entity.GetLaporanMutasiJenisDokumenResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.LaporanMutasiJenisDokumen
	}

	var (
		resp = new(entity.GetLaporanMutasiJenisDokumenResp)
		data = make([]dao, 0)
		args = make([]any, 0, 3)
	)

	query := `
			SELECT 
			COUNT(*) OVER() AS total_data,
			cb.id, 
			cb.no_kontrak, 
			cb.kode_document_bc AS kode_document, 
			cb.nomor_document_bc AS nomor_document, 
			cb.tanggal_document_bc AS tanggal_document, 
			s.name AS pemasok, 
			COALESCE(
				json_agg(
					json_build_object(
						'kode_barang', p.kode,
						'nama_barang', p.nama,
						'satuan', p.satuan,
						'jumlah', iip.jumlah,
						'harga_satuan', cp.harga_satuan
					)
				) FILTER (WHERE p.kode IS NOT NULL), 
				'[]'
			) AS barang
			FROM contracts_bc cb
			LEFT JOIN contracts c ON cb.no_kontrak = c.no_kontrak
			LEFT JOIN supliers s ON s.id = c.supliers_id  
			LEFT JOIN income_inventories_products iip ON iip.nomor_document_bc = cb.nomor_document_bc
			LEFT JOIN products p ON p.kode = iip.kode_barang 
			LEFT JOIN contract_products cp 
				ON cp.no_kontrak = c.no_kontrak 
				AND cp.kode_barang = p.kode
				WHERE cb.deleted_at IS NULL

	`
	if req.Q != "" {
		query += ` AND cb.no_kontrak ILIKE '%' || ? || '%'
		`
		args = append(args, req.Q)
	}
	query += ` GROUP BY cb.id, cb.no_kontrak, cb.kode_document_bc, cb.tanggal_document_bc, s.name`
	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasiDokumen - failed to query Laporan Mutasi Dokumen")
		return nil, err
	}
	resp.Items = make([]entity.LaporanMutasiJenisDokumen, 0, len(data))
	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.LaporanMutasiJenisDokumen)
	}
	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}
