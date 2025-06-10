package entity

import "inventori-beacukai-backend/pkg/types"

type GetIncomeInventoriesReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetIncomeInventoriesReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type IncomeInventory struct {
	Common
	NoKontrak      string `json:"no_kontrak" db:"no_kontrak"`
	KategoriBarang string `json:"kategori_barang" db:"kategori_barang"`
}

type GetIncomeInventoriesResp struct {
	Items []IncomeInventory `json:"items"`
	Meta  types.Meta        `json:"meta"`
}

type GetIncomeInventoryReq struct {
	Id string `json:"id" validate:"required"`
}

type GetIncomeInventoryResp struct {
	IncomeInventory
}

type CreateIncomeInventoryReq struct {
	NoKontrak      string `json:"no_kontrak" validate:"required"`
	KategoriBarang string `json:"kategori_barang" validate:"required,oneof='Bahan Baku' 'Bahan Penolong' 'Mesin/Sparepart'"`
}

type CreateIncomeInventoryResp struct {
	Id string `json:"id"`
}

type UpdateIncomeInventoryReq struct {
	Id             string `params:"id" validate:"required"`
	NoKontrak      string `json:"no_kontrak" validate:"required"`
	KategoriBarang string `json:"kategori_barang" validate:"required,oneof='Bahan Baku' 'Bahan Penolong' 'Mesin/Sparepart'"`
}

type DeleteIncomeInventoryReq struct {
	Id string `json:"id" validate:"required"`
}
