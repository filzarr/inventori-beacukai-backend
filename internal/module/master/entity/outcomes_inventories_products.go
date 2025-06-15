package entity

import "inventori-beacukai-backend/pkg/types"

type GetOutcomesInventoriesProductsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetOutcomesInventoriesProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type OutcomesInventoriesProduct struct {
	Common
	KodeBarang    string `json:"kode_barang" db:"kode_barang"`
	NamaBarang    string `json:"nama_barang" db:"nama_barang"`
	NoKontrak     string `json:"no_kontrak" db:"no_kontrak"`
	SaldoAwal     string `json:"saldo_awal" db:"saldo_awal"`
	JumlahKontrak int    `json:"jumlah_kontrak" db:"jumlah_kontrak"`
	Jumlah        int    `json:"jumlah_masuk" db:"jumlah_masuk"`
}

type GetOutcomesInventoriesProductsResp struct {
	Items []OutcomesInventoriesProduct `json:"items"`
	Meta  types.Meta                   `json:"meta"`
}

type GetOutcomesInventoriesProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetOutcomesInventoriesProductResp struct {
	OutcomesInventoriesProduct
}

type CreateOutcomesInventoriesProductReq struct {
	NoKontrak  string `json:"no_kontrak" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	SaldoAwal  int    `json:"saldo_awal" validate:"required"`
	Jumlah     int    `json:"jumlah" validate:"min=0"`
}

type CreateOutcomesInventoriesProductResp struct {
	Id string `json:"id"`
}

type UpdateOutcomesInventoriesProductReq struct {
	Id         string `params:"id" validate:"required"`
	NoKontrak  string `json:"no_kontrak" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	SaldoAwal  string `json:"saldo_awal" db:"saldo_awal"`
	Jumlah     int    `json:"jumlah" validate:"min=0"`
}

type DeleteOutcomesInventoriesProductReq struct {
	Id string `json:"id" validate:"required"`
}
