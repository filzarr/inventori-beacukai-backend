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
