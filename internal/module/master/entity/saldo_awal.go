package entity

import "inventori-beacukai-backend/pkg/types"

type GetSaldoAwalsReq struct {
	Q string `query:"q" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetSaldoAwalsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type SaldoAwal struct {
	Common
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	SaldoAwal  string `json:"saldo_awal" db:"saldo_awal"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
}

type GetSaldoAwalsResp struct {
	Items []SaldoAwal `json:"items"`
	Meta  types.Meta  `json:"meta"`
}

type GetSaldoAwalReq struct {
	Id string `params:"id" validate:"required"`
}

type GetSaldoAwalResp struct {
	SaldoAwal
}

type CreateSaldoAwalReq struct {
	KodeBarang string `json:"kode_barang" validate:"required"`
	SaldoAwal  string `json:"saldo_awal" validate:"required"`
}

type CreateSaldoAwalResp struct {
	Id string `json:"id"`
}

type UpdateSaldoAwalReq struct {
	Id         string `params:"id" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	SaldoAwal  string `json:"saldo_awal" validate:"required"`
}

type DeleteSaldoAwalReq struct {
	Id string `params:"id" validate:"required"`
}
