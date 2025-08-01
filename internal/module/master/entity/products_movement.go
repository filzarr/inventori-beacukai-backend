package entity

import (
	"inventori-beacukai-backend/pkg/types"
)

type GetProductsMovementReq struct {
	Q      string `query:"q" validate:"omitempty,min=3"`
	Status string `query:"status" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetProductsMovementReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type ProductsMovement struct {
	Common
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	Jumlah     int    `json:"jumlah" db:"jumlah"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Satuan     string `json:"satuan" db:"satuan"`
}

type GetProductsMovementResp struct {
	Items []ProductsMovement `json:"items"`
	Meta  types.Meta         `json:"meta"`
}

type GetProductsMovementReqID struct {
	Id string `params:"id" validate:"required"`
}

type GetProductsMovementRespID struct {
	ProductsMovement
}

type CreateProductsMovementReq struct {
	KodeBarang string `json:"kode_barang" validate:"required,min=1"`
	Jumlah     int    `json:"jumlah" validate:"required,gt=0"`
}

type CreateProductsMovementResp struct {
	Id string `json:"id"`
}

type UpdateProductsMovementReq struct {
	Id         string `params:"id" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required,min=1"`
	Jumlah     int    `json:"jumlah" validate:"required,gt=0"`
}

type DeleteProductsMovementReq struct {
	Id string `params:"id" validate:"required"`
}

type UpdateStatusProductsMoveMentReq struct {
	Id     string `params:"id" validate:"required"`
	Jumlah int    `json:"jumlah" validate:"required"`
}
