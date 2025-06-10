package entity

import "inventori-beacukai-backend/pkg/types"

type GetCurrenciesReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetCurrenciesReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Currency struct {
	Common
	Kode     string `json:"kode" db:"kode"`
	MataUang string `json:"mata_uang" db:"mata_uang"`
}

type GetCurrenciesResp struct {
	Items []Currency `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetCurrencyReq struct {
	Id string `json:"id" validate:"required"`
}

type GetCurrencyResp struct {
	Currency
}

type CreateCurrencyReq struct {
	Kode     string `json:"kode" validate:"required,min=2"`
	MataUang string `json:"mata_uang" validate:"required,min=3"`
}

type CreateCurrencyResp struct {
	Id string `json:"id"`
}

type UpdateCurrencyReq struct {
	Id       string `params:"id" validate:"required"`
	Kode     string `json:"kode" validate:"required,min=2"`
	MataUang string `json:"mata_uang" validate:"required,min=3"`
}

type DeleteCurrencyReq struct {
	Id string `json:"id" validate:"required"`
}
