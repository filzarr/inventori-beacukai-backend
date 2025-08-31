package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/dashboard/entity"
	"inventori-beacukai-backend/internal/module/dashboard/ports"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.DashboardRepository = &dashboardRepo{}

type dashboardRepo struct {
	db *sqlx.DB
}

func NewDashboardRepository() *dashboardRepo {
	return &dashboardRepo{
		db: adapter.Adapters.Postgres,
	}
}

func (r *dashboardRepo) GetDashboardChart(ctx context.Context, req *entity.GetPenjualanChartReq) (*entity.GetPenjualanChartResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.PenjualanChart
	}

	var (
		resp  = new(entity.GetPenjualanChartResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
			SELECT COUNT(*) OVER() AS total_data, cp.id, cp.kode_barang, p.nama AS nama_barang, cp.jumlah, c.tanggal AS tanggal
			FROM contract_products cp
			LEFT JOIN contracts c ON cp.no_kontrak = c.no_kontrak
			LEFT JOIN products p ON p.kode = cp.kode_barang
			WHERE EXISTS (
				SELECT 1
				FROM contracts_bc cb
				WHERE cb.no_kontrak = cp.no_kontrak
			)
			ORDER BY c.tanggal DESC`
	)
	resp.Items = make([]entity.PenjualanChart, 0)

	query += ` LIMIT ? OFFSET ?`
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetPenjualanChart - failed to query")
		return nil, err
	}

	for _, d := range data {
		resp.Meta.TotalData = d.TotalData
		resp.Items = append(resp.Items, d.PenjualanChart)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil
}

func (r *dashboardRepo) GetTotalPenjualan(ctx context.Context, req *entity.GetTotalPenjualanReq) (*entity.GetTotalPenjualanResp, error) {
	var (
		resp = new(entity.GetTotalPenjualanResp)
		data = new(entity.TotalPenjualan)
	)
	query := `
		SELECT 
			COALESCE(SUM(cp.jumlah * cp.nilai_barang_rp),0) AS total
		FROM contract_products cp
		LEFT JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		WHERE EXISTS (
				SELECT 1
				FROM contracts_bc cb
				WHERE cb.no_kontrak = cp.no_kontrak
			)
		AND c.kategori_kontrak = 'Penjualan'
		AND EXTRACT(YEAR FROM c.tanggal) = EXTRACT(YEAR FROM CURRENT_DATE) AND cp.deleted_at IS NULL
	`
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query)); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTotalPenjualan - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Dokumen BC tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTotalPenjualan - failed to get")
		return nil, err
	}

	resp.TotalPenjualan = *data
	return resp, nil
}
func (r *dashboardRepo) GetTotalPembelian(ctx context.Context, req *entity.GetTotalPembelianReq) (*entity.GetTotalPembelianResp, error) {
	var (
		resp = new(entity.GetTotalPembelianResp)
		data = new(entity.TotalPembelian)
	)
	query := `
		SELECT 
			COALESCE(SUM(cp.jumlah * cp.nilai_barang_rp),0) AS total
		FROM contract_products cp
		LEFT JOIN contracts c ON cp.no_kontrak = c.no_kontrak
		WHERE EXISTS (
				SELECT 1
				FROM contracts_bc cb
				WHERE cb.no_kontrak = cp.no_kontrak
			)
		AND c.kategori_kontrak = 'Pembelian'
		AND EXTRACT(YEAR FROM c.tanggal) = EXTRACT(YEAR FROM CURRENT_DATE) AND cp.deleted_at IS NULL
	`
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query)); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTotalPembelian - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTotalPembelian - failed to get")
		return nil, err
	}

	resp.TotalPembelian = *data
	return resp, nil
}

func (r *dashboardRepo) GetTotalWipToday(ctx context.Context, req *entity.GetTotalWipTodayReq) (*entity.GetTotalWipTodayResp, error) {
	var (
		resp = new(entity.GetTotalWipTodayResp)
		data = new(entity.TotalWipToday)
	)
	query := `
		SELECT 
			COALESCE(SUM(pm.jumlah), 0) AS total
		FROM products_movement pm
		LEFT JOIN warehouses w ON pm.warehouses_to = w.kode
		WHERE w.kategori = 'Produksi'
		AND pm.status_perpindahan = 'Diterima'
		AND DATE(pm.created_at) = CURRENT_DATE
		AND pm.deleted_at IS NULL
	`
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query)); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTotalWipToday - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTotalPembelian - failed to get")
		return nil, err
	}

	resp.TotalWipToday = *data
	return resp, nil
}

func (r *dashboardRepo) GetTotalProductMovementNotProcess(ctx context.Context, req *entity.GetTotalProductMovementNotProcessReq) (*entity.GetTotalProductMovementNotProcessResp, error) {
	var (
		resp = new(entity.GetTotalProductMovementNotProcessResp)
		data = new(entity.TotalProductMovementNotProcess)
	)
	query := `
		SELECT 
			COALESCE(COUNT(pm.*), 0) AS total
		FROM products_movement pm
		LEFT JOIN warehouses w ON pm.warehouses_to = w.kode
		WHERE pm.status_perpindahan = 'Diminta'
		AND DATE(pm.created_at) = CURRENT_DATE
		AND pm.deleted_at IS NULL
	`
	if err := r.db.GetContext(ctx, data, r.db.Rebind(query)); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetTotalProductMovementNotProcess - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetTotalProductMovementNotProcess - failed to get")
		return nil, err
	}

	resp.TotalProductMovementNotProcess = *data
	return resp, nil
}

func (r *dashboardRepo) GetTotalStockMiminum(ctx context.Context, req *entity.GetTotalStockMiminumReq) (*entity.GetTotalStockMiminumResp, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.TotalStockMiminum
	}

	var (
		resp  = new(entity.GetTotalStockMiminumResp)
		data  = make([]dao, 0)
		args  = make([]any, 0, 3)
		query = `
		SELECT 
			COUNT(*) OVER() AS total_data,
			p.kode AS kode_barang,
			p.nama AS nama_barang, 
			COALESCE(SUM(ws.jumlah), 0) AS total
		FROM products p
		LEFT JOIN warehouses_stocks ws 
			ON p.kode = ws.kode_barang 
			AND ws.deleted_at IS NULL
		LEFT JOIN warehouses w
			ON ws.warehouse_kode = w.kode
			AND w.kategori NOT IN ('Produksi', 'Penjualan')
			AND w.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
		GROUP BY p.kode, p.nama, p.satuan, p.kategori
		HAVING COALESCE(SUM(ws.jumlah), 0) < 50
		ORDER BY p.nama ASC
		LIMIT ? OFFSET ?`
	)

	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	if err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::GetTotalStockMinimum - failed to query")
		return nil, err
	}

	resp.Items = make([]entity.TotalStockMiminum, 0, len(data))

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		resp.Items = append(resp.Items, d.TotalStockMiminum)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)
	return resp, nil

}
