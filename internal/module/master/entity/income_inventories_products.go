package entity

import "inventori-beacukai-backend/pkg/types"

type GetIncomeInventoryProductsReq struct {
	Q    string `query:"q" validate:"omitempty,min=3"`
	Full bool   `query:"full" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetIncomeInventoryProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type IncomeInventoryProduct struct {
	Common
	KodeBarang    string `json:"kode_barang" db:"kode_barang"`
	NamaBarang    string `json:"nama_barang" db:"nama_barang"`
	NoKontrak     string `json:"no_kontrak" db:"no_kontrak"`
	SaldoAwal     string `json:"saldo_awal" db:"saldo_awal"`
	JumlahKontrak int    `json:"jumlah_kontrak" db:"jumlah_kontrak"`
	Jumlah        int    `json:"jumlah_masuk" db:"jumlah_masuk"`
}

type GetIncomeInventoryProductsResp struct {
	Items []IncomeInventoryProduct `json:"items"`
	Meta  types.Meta               `json:"meta"`
}

type GetIncomeInventoryProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetIncomeInventoryProductResp struct {
	IncomeInventoryProduct
}

type CreateIncomeInventoryProductReq struct {
	NoKontrak  string `json:"no_kontrak" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	Lokasi     string `json:"lokasi" validate:"required"`
	SaldoAwal  int    `json:"saldo_awal" validate:"required"`
	Jumlah     int    `json:"jumlah" validate:"min=0"`
}

type CreateIncomeInventoryProductResp struct {
	Id string `json:"id"`
}

type UpdateIncomeInventoryProductReq struct {
	Id         string `params:"id" validate:"required"`
	NoKontrak  string `json:"no_kontrak" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	SaldoAwal  string `json:"saldo_awal" db:"saldo_awal"`
	Jumlah     int    `json:"jumlah" validate:"min=0"`
}

type DeleteIncomeInventoryProductReq struct {
	Id string `json:"id" validate:"required"`
}
