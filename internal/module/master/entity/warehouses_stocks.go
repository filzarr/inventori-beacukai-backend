package entity

import "inventori-beacukai-backend/pkg/types"

type GetWarehousesStocksReq struct {
	Q             string `query:"q" validate:"omitempty"`
	WarehouseKode string `query:"warehouse" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetWarehousesStocksReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type WarehousesStock struct {
	Common
	WarehouseKode string `json:"warehouse_kode" db:"warehouse_kode"`
	NamaBarang    string `json:"nama_barang" db:"nama_barang"`
	KodeBarang    string `json:"kode_barang" db:"kode_barang"`
	Satuan        string `json:"satuan" db:"satuan"`
	Kategori      string `json:"kategori" db:"kategori"`
	Jumlah        int    `json:"jumlah" db:"jumlah"`
}

type GetWarehousesStocksResp struct {
	Items []WarehousesStock `json:"items"`
	Meta  types.Meta        `json:"meta"`
}

type GetWarehousesStockReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}

type GetWarehousesStockResp struct {
	WarehousesStock
}

type CreateWarehousesStockReq struct {
	UserId string `validate:"ulid"`

	WarehouseKode string `json:"warehouse_kode" validate:"required"`
	KodeBarang    string `json:"kode_barang" validate:"required"`
	Jumlah        int    `json:"jumlah" validate:"required,min=0"`
}

type CreateWarehousesStockResp struct {
	Id string `json:"id"`
}

type UpdateWarehousesStockReq struct {
	UserId string `validate:"ulid"`

	Id            string `params:"id" validate:"required"`
	WarehouseKode string `json:"warehouse_kode" validate:"required"`
	KodeBarang    string `json:"kode_barang" validate:"required"`
	Jumlah        int    `json:"jumlah" validate:"required,min=0"`
}

type DeleteWarehousesStockReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}

type UpdateStockWarehousesReq struct {
	UserId string `validate:"ulid"`

	Id     string `params:"id" validate:"required"`
	Jumlah int    `json:"jumlah" validate:"required,min=0"`
}
