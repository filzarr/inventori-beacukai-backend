package entity

import (
	"inventori-beacukai-backend/pkg/types"
)

type GetProductionsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetProductionsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Production struct {
	Common
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Satuan     string `json:"satuan" db:"satuan"`
	Jumlah     int    `json:"jumlah" db:"jumlah"`
}

type GetProductionsResp struct {
	Items []Production `json:"items"`
	Meta  types.Meta   `json:"meta"`
}

type GetProductionReq struct {
	Id string `params:"id" validate:"required"`
}

type GetProductionResp struct {
	Production
}

type CreateProductionReq struct {
	KodeBarang string `json:"kode_barang" validate:"required,min=1"`
	Jumlah     int    `json:"jumlah" validate:"required,gt=0"`
}

type CreateProductionResp struct {
	Id string `json:"id"`
}

type UpdateProductionReq struct {
	Id         string `params:"id" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required,min=1"`
	Jumlah     int    `json:"jumlah" validate:"required,gt=0"`
}

type DeleteProductionReq struct {
	Id string `json:"id" validate:"required"`
}
