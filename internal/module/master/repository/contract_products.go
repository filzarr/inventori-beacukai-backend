package repository

import (
	"context"
	"database/sql"
	"strings"

	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *masterRepo) GetContractProducts(ctx context.Context, req *entity.GetContractProductsReq) (*entity.GetContractProductsResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ContractProduct
	}

	var (
		resp    = new(entity.GetContractProductsResp)
		data    = make([]dao, 0)
		args    = make([]any, 0, 5)
		filters = []string{"cp.deleted_at IS NULL"}
	)

	resp.Items = make([]entity.ContractProduct, 0)

	// Tambah filter NoKontrak jika ada
	if req.NoKontrak != "" {
		filters = append(filters, "cp.no_kontrak = ?")
		args = append(args, req.NoKontrak)
	}

	// Tambah filter pencarian jika ada
	if req.Q != "" {
		filters = append(filters, "cp.no_kontrak ILIKE '%' || ? || '%'")
		args = append(args, req.Q)
	}

	// Bangun query dengan filter
	query := `
		SELECT
			COUNT(*) OVER() AS total_data,
			cp.id,
			cp.no_kontrak,
			cp.kode_barang,
			cp.jumlah,
			p.jumlah AS stok,
			p.nama AS nama_barang,
			s.name AS nama_pemasok,
			s.alamat,
			cp.satuan,
			cp.harga_satuan,
			cp.kode_mata_uang,
			cp.nilai_barang_fog,
			cp.nilai_barang_rp
		FROM
			contract_products cp
		JOIN products p ON cp.kode_barang = p.kode
		JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		JOIN supliers s ON c.supliers_id = s.id
	`

	if len(filters) > 0 {
		query += " WHERE " + joinFilters(filters)
	}

	query += " ORDER BY cp.created_at DESC LIMIT ? OFFSET ?"
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetContractProducts - failed to query contract products")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.ContractProduct)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *masterRepo) GetContractProduct(ctx context.Context, req *entity.GetContractProductReq) (*entity.GetContractProductResp, error) {
	var (
		resp = new(entity.GetContractProductResp)
		data = new(entity.ContractProduct)
	)

	query := `
		SELECT
			cp.id,
			cp.no_kontrak,
			cp.kode_barang,
			cp.jumlah,
			p.nama AS nama_barang,
			s.name AS nama_pemasok,
			s.alamat AS alamat,
			cp.satuan,
			cp.harga_satuan,
			cp.kode_mata_uang,
			cp.nilai_barang_fog,
			cp.nilai_barang_rp
		FROM
			contract_products cp
		JOIN products p ON cp.kode_barang = p.id
		JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		JOIN supliers s ON c.supliers_id = s.id
		WHERE
			id = ?
			AND cp.deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, data, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetContractProduct - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Data tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetContractProduct - failed to get")
		return nil, err
	}

	resp.ContractProduct = *data
	return resp, nil
}

func (r *masterRepo) CreateContractProduct(ctx context.Context, req *entity.CreateContractProductReq) (*entity.CreateContractProductResp, error) {
	query := `
		INSERT INTO contract_products (
			id, no_kontrak, kode_barang, jumlah, satuan, harga_satuan, kode_mata_uang, nilai_barang_fog, nilai_barang_rp
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.CreateContractProductResp)
	)

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), Id, req.NoKontrak, req.KodeBarang, req.Jumlah, req.Satuan, req.HargaSatuan, req.KodeMataUang, req.NilaiBarangFog, req.NilaiBarangRp); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateContractProduct - failed to insert")
		return nil, err
	}

	resp.Id = Id
	return resp, nil
}

func (r *masterRepo) UpdateContractProduct(ctx context.Context, req *entity.UpdateContractProductReq) error {
	query := `
		UPDATE contract_products
		SET
			no_kontrak = ?,
			kode_barang = ?,
			jumlah = ?,
			satuan = ?,
			harga_satuan = ?,
			kode_mata_uang = ?,
			nilai_barang_fog = ?,
			nilai_barang_rp = ?,
			updated_at = NOW()
		WHERE
			id = ?
			AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.NoKontrak, req.KodeBarang, req.Jumlah, req.Satuan, req.HargaSatuan, req.KodeMataUang, req.NilaiBarangFog, req.NilaiBarangRp, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateContractProduct - failed to update")
		return err
	}

	return nil
}

func (r *masterRepo) DeleteContractProduct(ctx context.Context, req *entity.DeleteContractProductReq) error {
	query := `
		UPDATE contract_products
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::DeleteContractProduct - failed to delete")
		return err
	}

	return nil
}

func joinFilters(filters []string) string {
	return strings.Join(filters, " AND ")
}
