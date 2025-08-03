package entity

import "inventori-beacukai-backend/pkg/types"

type GetWarehousesReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetWarehousesReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Warehouse struct {
	Common
	Kode       string `json:"kode" db:"kode"`
	Nama       string `json:"nama" db:"nama"`
	Kategori   string `json:"kategori" db:"kategori"`
	Keterangan string `json:"keterangan" db:"keterangan"`
}

type GetWarehousesResp struct {
	Items []Warehouse `json:"items"`
	Meta  types.Meta  `json:"meta"`
}

type GetWarehouseReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}

type GetWarehouseResp struct {
	Warehouse
}

type CreateWarehouseReq struct {
	UserId string `validate:"ulid"`

	Kode       string `json:"kode" validate:"required"`
	Nama       string `json:"nama" validate:"required"`
	Kategori   string `json:"kategori" validate:"required"`
	Keterangan string `json:"keterangan" validate:"omitempty"`
}

type CreateWarehouseResp struct {
	Id string `json:"id"`
}

type UpdateWarehouseReq struct {
	UserId string `validate:"ulid"`

	Id         string `params:"id" validate:"required"`
	Kode       string `json:"kode" validate:"required"`
	Nama       string `json:"nama" validate:"required"`
	Kategori   string `json:"kategori" validate:"required"`
	Keterangan string `json:"keterangan" validate:"omitempty"`
}

type DeleteWarehouseReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}
