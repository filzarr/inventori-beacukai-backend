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
			b.kategori AS kode_dokumen,
			t.no_document AS no_dokumen,
			TO_CHAR(b.tanggal, 'YYYY-MM-DD') AS tanggal_dokumen,
			t.no_kontrak,
			p.kode AS kode_barang,
			p.nama AS nama_barang,
			cp.satuan AS satuan,
			COALESCE(iip.stok_awal, 0) AS saldo_awal,
			COALESCE(iip.jumlah, 0) AS pemasukan,
			t.jumlah AS stok_opname,
			p.jumlah AS stok_akhir
		FROM
			transaction_incomes t
		JOIN products p ON t.kode_barang = p.kode
		JOIN contracts c ON t.no_kontrak = c.no_kontrak
		JOIN contract_products cp ON cp.no_kontrak = c.no_kontrak AND cp.kode_barang = p.kode
		JOIN bc_documents b ON b.no_document = t.no_document
		LEFT JOIN income_inventories i ON i.no_kontrak = c.no_kontrak
		LEFT JOIN income_inventories_products iip ON iip.id_inventories = i.id AND iip.kode_barang = p.kode
		WHERE t.deleted_at IS NULL
		`
	)

	if req.Q != "" {
		query += ` AND (
			t.no_kontrak ILIKE '%' || ? || '%' OR
			t.no_document ILIKE '%' || ? || '%' OR
			p.kode ILIKE '%' || ? || '%' OR
			p.nama ILIKE '%' || ? || '%'
		)`
		args = append(args, req.Q, req.Q, req.Q, req.Q)
	}

	query += " ORDER BY b.tanggal DESC"

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetLaporanMutasi - failed to query")
		return nil, err
	}

	return &entity.GetLaporanMutasiResp{
		Items: data,
		Meta:  types.Meta{TotalData: len(data)}, // Optional: Hitung manual kalau tidak pakai COUNT()
	}, nil
}
