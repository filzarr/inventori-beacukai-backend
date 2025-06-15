package entity

import "inventori-beacukai-backend/pkg/types"

type GetTransfersProductsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetTransfersProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type TransferProduct struct {
	Common
	KodeBarang string `json:"kode_barang" db:"kode_barang"`
	Jumlah     int    `json:"jumlah" db:"jumlah"`
}

type GetTransfersProductsResp struct {
	Items []TransferProduct `json:"items"`
	Meta  types.Meta        `json:"meta"`
}

type GetTransferProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetTransferProductResp struct {
	TransferProduct
}

type CreateTransferProductReq struct {
	KodeBarang string `json:"kode_barang" validate:"required,min=3"`
	Jumlah     int    `json:"jumlah" validate:"required"`
}

type CreateTransferProductResp struct {
	Id string `json:"id"`
}

type UpdateTransferProductReq struct {
	Id         string `params:"id" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required,min=3"`
	Jumlah     int    `json:"jumlah" validate:"required"`
}

type DeleteTransferProductReq struct {
	Id string `json:"id" validate:"required"`
}
