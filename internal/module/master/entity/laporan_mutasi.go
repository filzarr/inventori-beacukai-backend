package entity

import "inventori-beacukai-backend/pkg/types"

type GetLaporanMutasiReq struct {
	Q        string `query:"q" validate:"omitempty,min=3"`
	Kategori string `query:"kategori" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetLaporanMutasiReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type LaporanMutasi struct {
	Id         string `json:"id" db:"id"`
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Satuan     string `json:"satuan" db:"satuan"`
	SaldoAwal  int    `json:"saldo_awal" db:"saldo_awal"`
	Pemasukan  int    `json:"pemasukan" db:"pemasukan"`
	StokOpname int    `json:"stok_opname" db:"stok_opname"`
}

type GetLaporanMutasiResp struct {
	Items []LaporanMutasi `json:"items"`
	Meta  types.Meta      `json:"meta"`
}

type GetLaporanMutasiPemasukanReq struct {
	Q          string `query:"q" validate:"omitempty,min=3"`
	KodeBarang string `query:"kode_barang"`
	types.MetaQuery
}

func (r *GetLaporanMutasiPemasukanReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type LaporanMutasiPemasukan struct {
	Id          string `json:"id" db:"id"`
	KodeDokumen string `json:"kode_document" db:"kode_document"`
	Kategori    string `json:"kategori" db:"kategori"`
	NoKontrak   string `json:"no_kontrak" db:"no_kontrak"`
	Tanggal     string `json:"tanggal" db:"tanggal"`
	Jumlah      int    `json:"jumlah" db:"jumlah"`
	Satuan      string `json:"satuan" db:"satuan"`
	KodeBarang  string `json:"kode_barang" db:"kode_barang"`
	NamaBarang  string `json:"nama_barang" db:"nama_barang"`
}

type GetLaporanMutasiPemasukanResp struct {
	Items []LaporanMutasiPemasukan `json:"items"`
	Meta  types.Meta               `json:"meta"`
}
