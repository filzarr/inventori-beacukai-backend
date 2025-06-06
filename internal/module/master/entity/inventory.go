package entity

import "inventori-beacukai-backend/pkg/types"

type GetInventoriesReq struct {
	Q        string `query:"q" validate:"omitempty,min=3"`
	Kategori string `query:"kategori" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetInventoriesReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Inventory struct {
	Common
	InventoriesStatus string `json:"inventories_status" db:"inventories_status"`
	KodeBarang        string `json:"kode_barang" db:"kode_barang"`
	NamaBarang        string `json:"nama_barang" db:"nama_barang"`
	Kategori          string `json:"kategori" db:"kategori"`
	Jumlah            int    `json:"jumlah" db:"jumlah"`
}

type GetInventoriesResp struct {
	Items []Inventory `json:"items"`
	Meta  types.Meta  `json:"meta"`
}

type GetInventoryReq struct {
	Id string `json:"id" validate:"required"`
}

type GetInventoryResp struct {
	Inventory
}

type CreateInventoryReq struct {
	InventoriesStatus string `json:"inventories_status" validate:"required"`
	KodeBarang        string `json:"kode_barang" validate:"required"`
	NamaBarang        string `json:"nama_barang" validate:"required"`
	Kategori          string `json:"kategori" validate:"required"`
	Jumlah            string `json:"jumlah" db:"jumlah"`
}

type CreateInventoryResp struct {
	Id string `json:"id"`
}

type UpdateInventoryReq struct {
	Id                string `json:"id" validate:"required"`
	InventoriesStatus string `json:"inventories_status" validate:"required"`
	KodeBarang        string `json:"kode_barang" validate:"required"`
	NamaBarang        string `json:"nama_barang" validate:"required"`
	Kategori          string `json:"kategori" validate:"required"`
	Jumlah            int    `json:"jumlah" db:"jumlah"`
}

type DeleteInventoryReq struct {
	Id string `json:"id" validate:"required"`
}

type GetInventoriesBahanBakuReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetInventoriesBahanBakuReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type GetInventoriesBahanBakuResp struct {
	Items []Inventory `json:"items"`
	Meta  types.Meta  `json:"meta"`
}
