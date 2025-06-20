package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
	"inventori-beacukai-backend/pkg/types"

	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetLaporanMutasi(ctx context.Context, req *entity.GetLaporanMutasiReq) (*entity.GetLaporanMutasiResp, error) {
	var (
		data  = make([]entity.LaporanMutasi, 0)
		args  = make([]any, 0)
		query = `
		SELECT   
			p.id,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.saldo_awal AS saldo_awal,
			SUM(iip.jumlah) AS pemasukan,
			p.jumlah AS stok_opname
		FROM
			products p
		JOIN income_inventories_products iip ON p.kode = iip.kode_barang
		GROUP BY 
			p.id ,p.kode
		`
	)

	if req.Q != "" {
		query += ` AND (
			p.kode ILIKE '%' || ? || '%' OR
			p.nama ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q, req.Q, req.Q)
	}

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasi - failed to query")
		return nil, err
	}

	return &entity.GetLaporanMutasiResp{
		Items: data,
		Meta:  types.Meta{TotalData: len(data)}, // Optional: Hitung manual kalau tidak pakai COUNT()
	}, nil
}

func (r *masterRepo) GetLaporanMutasiPemasukan(ctx context.Context, req *entity.GetLaporanMutasiPemasukanReq) (*entity.GetLaporanMutasiPemasukanResp, error) {
	var (
		data  = make([]entity.LaporanMutasiPemasukan, 0)
		args  = make([]any, 0)
		query = `
		SELECT   
			cp.id,
			bc.kategori AS kode_document,
			bc.no_document AS no_document,
			bc.tanggal,
			cp.no_kontrak,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			p.satuan AS satuan,
			iip.jumlah AS jumlah
			
		FROM
			contract_products cp
		JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		JOIN bc_documents bc ON c.no_document = bc.no_document
		JOIN products p ON cp.kode_barang = p.kode
		JOIN income_inventories_products iip ON c.no_kontrak = iip.no_kontrak
		`
	)

	if req.KodeBarang != "" {
		query += `WHERE
			cp.kode_barang = ?
		`
		args = append(args, req.KodeBarang)
	}
	if req.Q != "" {
		query += ` AND (
			p.kode ILIKE '%' || ? || '%' OR
			p.nama ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q, req.Q, req.Q)
	}

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasi - failed to query")
		return nil, err
	}

	return &entity.GetLaporanMutasiPemasukanResp{
		Items: data,
		Meta:  types.Meta{TotalData: len(data)}, // Optional: Hitung manual kalau tidak pakai COUNT()
	}, nil
}
