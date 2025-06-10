package entity

import "inventori-beacukai-backend/pkg/types"

type GetProductsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetProductsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Product struct {
	Common
	Kode     string `json:"kode" db:"kode"`
	Nama     string `json:"nama" db:"nama"`
	Kategori string `json:"kategori" db:"kategori"`
	Jumlah   int    `json:"jumlah" db:"jumlah"`
}

type GetProductsResp struct {
	Items []Product  `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetProductReq struct {
	Id string `json:"id" validate:"required"`
}

type GetProductResp struct {
	Product
}

type CreateProductReq struct {
	Kode     string `json:"kode" validate:"required,min=3"`
	Nama     string `json:"nama" validate:"required,min=3"`
	Kategori string `json:"kategori" validate:"required,oneof='Bahan Baku' 'Bahan Penolong' 'Mesin/Sparepart'"`
	Jumlah   int    `json:"jumlah" validate:"min=0"`
}

type CreateProductResp struct {
	Id string `json:"id"`
}

type UpdateProductReq struct {
	Id       string `params:"id" validate:"required"`
	Kode     string `json:"kode" validate:"required,min=3"`
	Nama     string `json:"nama" validate:"required,min=3"`
	Kategori string `json:"kategori" validate:"required,oneof='Bahan Baku' 'Bahan Penolong' 'Mesin/Sparepart'"`
	Jumlah   int    `json:"jumlah" validate:"min=0"`
}

type DeleteProductReq struct {
	Id string `json:"id" validate:"required"`
}
