package entity

import "inventori-beacukai-backend/pkg/types"

type GetLaporanMutasiReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
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
