package entity

import "inventori-beacukai-backend/pkg/types"

type GetReadyProductsReq struct {
	Q string `query:"q" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetReadyProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type ReadyProduct struct {
	Common
	Kode   string `json:"kode" db:"kode"`
	Nama   string `json:"nama" db:"nama"`
	Satuan string `json:"satuan" db:"satuan"`
	Jumlah int    `json:"jumlah" db:"jumlah"`
}

type GetReadyProductsResp struct {
	Items []ReadyProduct `json:"items"`
	Meta  types.Meta     `json:"meta"`
}

type GetReadyProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetReadyProductResp struct {
	ReadyProduct
}

type CreateReadyProductReq struct {
	Kode   string `json:"kode" validate:"required,min=3"`
	Nama   string `json:"nama" validate:"required,min=3"`
	Satuan string `json:"satuan" validate:"required,min=3"`
	Jumlah int    `json:"jumlah" validate:"min=0"`
}

type CreateReadyProductResp struct {
	Id string `json:"id"`
}

type UpdateReadyProductReq struct {
	Id     string `params:"id" validate:"required"`
	Kode   string `json:"kode" validate:"required,min=3"`
	Nama   string `json:"nama" validate:"required,min=3"`
	Satuan string `json:"satuan" validate:"required,min=3"`
	Jumlah int    `json:"jumlah" validate:"min=0"`
}

type DeleteReadyProductReq struct {
	Id string `json:"id" validate:"required"`
}
