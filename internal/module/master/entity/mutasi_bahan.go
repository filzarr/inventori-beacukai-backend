package entity

import (
	"inventori-beacukai-backend/pkg/types"
)

type GetMutasiBahansReq struct {
	Q      string `query:"q" validate:"omitempty,min=3"`
	Gudang string `query:"gudang" validate:"required"`
	types.MetaQuery
}

func (r *GetMutasiBahansReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type MutasiBahan struct {
	Common
	Gudang      string `json:"nama_gudang" db:"nama_gudang"`
	Jumlah      int    `json:"jumlah" db:"jumlah"`
	NamaBarang  string `json:"nama_barang" db:"nama_barang"`
	KodeBarang  string `json:"kode_barang" db:"kode_barang"`
	IdBarang    string `json:"id_barang" db:"id_barang"`
	GudangId    string `json:"gudang_id" db:"gudang_id"`
	Status      string `json:"status" db:"status"`
	SaldoAwal   string `json:"saldo_awal" db:"saldo_awal"`
	Pemasukan   string `json:"pemasukan" db:"pemasukan"`
	Pengeluaran string `json:"pengeluaran" db:"pengeluaran"`
	Penyesuaian string `json:"penyesuaian" db:"penyesuaian"`
}

type GetMutasiBahansResp struct {
	Items []MutasiBahan `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type GetMutasiBahanReq struct {
	Id string `json:"id" validate:"required"`
}

type GetMutasiBahanResp struct {
	MutasiBahan
}

type CreateMutasiBahanReq struct {
	Inventories string `json:"inventories" validate:"required"`
	Gudang      string `json:"gudang" validate:"required"`
	Jumlah      int    `json:"jumlah" validate:"required"`
}

type CreateMutasiBahanResp struct {
	Id string `json:"id"`
}

type UpdateMutasiBahanReq struct {
	Id          string `json:"id" validate:"required"`
	Inventories string `json:"inventories" validate:"required"`
	Gudang      string `json:"gudang" validate:"required"`
	Jumlah      int    `json:"jumlah" validate:"required"`
	Status      string `json:"status" validate:"required"`
	SaldoAwal   int    `json:"saldo_awal" validate:"required"`
	Pemasukan   int    `json:"pemasukan" validate:"required"`
	Pengeluaran int    `json:"pengeluaran" validate:"required"`
	Penyesuaian int    `json:"penyesuaian" validate:"required"`
}

type UpdateSaldoMutasiReq struct {
	Id          string `json:"id" validate:"required"`
	SaldoAwal   int    `json:"saldo_awal" validate:"required"`
	Pemasukan   int    `json:"pemasukan" validate:"required"`
	Pengeluaran int    `json:"pengeluaran" validate:"required"`
	Penyesuaian int    `json:"penyesuaian" validate:"required"`
}

type DeleteMutasiBahanReq struct {
	Id string `json:"id" validate:"required"`
}

type UpdateStatusMutasiBahanReq struct {
	Id        string `json:"id" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Jumlah    int    `json:"jumlah" validate:"required"`
	Inventory string `json:"inventory" validate:"required"`
}
type GetLaporanMutasiBahanReq struct {
	Q      string `query:"q" validate:"omitempty,min=3"`
	Gudang string `query:"gudang"`
	types.MetaQuery
}
type LaporanMutasiBahan struct {
	Common
	KodeBarang  string `json:"kode_barang" db:"kode_barang"`
	NamaBarang  string `json:"nama_barang" db:"nama_barang"`
	SaldoAwal   string `json:"saldo_awal" db:"saldo_awal"`
	Pemasukan   string `json:"pemasukan" db:"pemasukan"`
	Penyesuaian string `json:"penyesuaian" db:"penyesuaian"`
	Pengeluaran string `json:"pengeluaran" db:"pengeluaran"`
	Jumlah      int    `json:"jumlah" db:"jumlah"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

func (r *GetLaporanMutasiBahanReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type GetLaporanMutasiBahanResp struct {
	Items []LaporanMutasiBahan `json:"items"`
	Meta  types.Meta           `json:"meta"`
}
