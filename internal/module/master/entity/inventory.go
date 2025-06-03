package entity

import "inventori-beacukai-backend/pkg/types"

type GetInventoriesReq struct {
	Q                 string `query:"q" validate:"omitempty,min=3"`
	Kategori          string `query:"kategori" validate:"omitempty,min=3"`
	DocumentStatus    string `query:"document-status" validate:"omitempty,min=3"`
	DocumentType      string `query:"document-type" validate:"omitempty,min=3"`
	InventoriesStatus string `query:"inventories-status" validate:"omitempty,min=3"`
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
	Pemasok           string `json:"pemasok" db:"pemasok"`
	Pembeli           string `json:"pembeli" db:"pembeli"`
	SaldoAwal         string `json:"saldo_awal" db:"saldo_awal"`
	Satuan            string `json:"satuan" db:"satuan"`
	StokOpname        string `json:"stok_opname" db:"stok_opname"`
	MataUang          string `json:"mata_uang" db:"mata_uang"`
	NegaraAsal        string `json:"negara_asal" db:"negara_asal"`
	DocumentType      string `json:"document_type" db:"document_type"`
	Keterangan        string `json:"keterangan" db:"keterangan"`
	DocumentStatus    string `json:"document_status" db:"document_status"`
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
	Pemasok           string `json:"pemasok" validate:"required"`
	Pembeli           string `json:"pembeli" validate:"required"`
	SaldoAwal         string `json:"saldo_awal" validate:"required"`
	Satuan            string `json:"satuan" validate:"required"`
	StokOpname        string `json:"stok_opname" validate:"required"`
	MataUang          string `json:"mata_uang" validate:"required"`
	NegaraAsal        string `json:"negara_asal" validate:"required"`
	DocumentType      string `json:"document_type" validate:"required"`
	Keterangan        string `json:"keterangan" validate:"required"`
	DocumentStatus    string `json:"document_status" validate:"required"`
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
	Pemasok           string `json:"pemasok" validate:"required"`
	Pembeli           string `json:"pembeli" validate:"required"`
	SaldoAwal         string `json:"saldo_awal" validate:"required"`
	Satuan            string `json:"satuan" validate:"required"`
	StokOpname        string `json:"stok_opname" validate:"required"`
	MataUang          string `json:"mata_uang" validate:"required"`
	NegaraAsal        string `json:"negara_asal" validate:"required"`
	DocumentType      string `json:"document_type" validate:"required"`
	Keterangan        string `json:"keterangan" validate:"required"`
	DocumentStatus    string `json:"document_status" validate:"required"`
}

type DeleteInventoryReq struct {
	Id string `json:"id" validate:"required"`
}
