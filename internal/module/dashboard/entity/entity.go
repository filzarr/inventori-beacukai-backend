package entity

import "inventori-beacukai-backend/pkg/types"

type GetPenjualanChartReq struct {
	Q      string `query:"q" validate:"omitempty"`
	UserId string `json:"user_id"`
	types.MetaQuery
}

func (r *GetPenjualanChartReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type PenjualanChart struct {
	Id           string `json:"id" db:"id"`
	NomorKontrak string `json:"nomor_kontrak" db:"nomor_kontrak"`
	KodeBarang   string `json:"kode_barang" db:"kode_barang"`
	NamaBarang   string `json:"nama_barang" db:"nama_barang"`
	Jumlah       string `json:"jumlah" db:"jumlah"`
	Tanggal      string `json:"tanggal" db:"tanggal"`
}

type GetPenjualanChartResp struct {
	Items []PenjualanChart `json:"items"`
	Meta  types.Meta       `json:"meta"`
}

type GetTotalPenjualanReq struct {
	UserId string `json:"user_id"`
}

type TotalPenjualan struct {
	Total int `json:"total" db:"total"`
}

type GetTotalPenjualanResp struct {
	TotalPenjualan
}

type GetTotalPembelianReq struct {
	UserId string `json:"user_id"`
}

type TotalPembelian struct {
	Total int `json:"total" db:"total"`
}

type GetTotalPembelianResp struct {
	TotalPembelian
}

type GetTotalWipTodayReq struct {
	UserId string `json:"user_id"`
}

type TotalWipToday struct {
	Total int `json:"total" db:"total"`
}

type GetTotalWipTodayResp struct {
	TotalWipToday
}

type GetTotalProductMovementNotProcessReq struct {
	UserId string `json:"user_id"`
}

type TotalProductMovementNotProcess struct {
	Total int `json:"total" db:"total"`
}
type GetTotalProductMovementNotProcessResp struct {
	TotalProductMovementNotProcess
}

type GetTotalStockMiminumReq struct {
	UserId string `json:"user_id"`
	Q      string `query:"q" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetTotalStockMiminumReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type TotalStockMiminum struct {
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Total      int    `json:"total" db:"total"`
}
type GetTotalStockMiminumResp struct {
	Items []TotalStockMiminum `json:"items"`
	Meta  types.Meta          `json:"meta"`
}
