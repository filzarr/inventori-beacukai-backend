package entity

import "inventori-beacukai-backend/pkg/types"

type GetTransactionIncomesReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetTransactionIncomesReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type TransactionIncome struct {
	Common
	KodeDocumentBc string `json:"kode_document_bc" db:"kode_document_bc"`
	NoDocumentBc   string `json:"no_document" db:"no_document"`
	TglDocumentBc  string `json:"tgl_document_bc" db:"tgl_document_bc"`
	NoKontrak      string `json:"no_kontrak" db:"no_kontrak"`
	Kategori       string `json:"kategori_barang" db:"kategori_barang"`
	KodeBarang     string `json:"kode_barang" db:"kode_barang"`
	NamaBarang     string `json:"nama_barang" db:"nama_barang"`
	Jumlah         int    `json:"jumlah" db:"jumlah"`
}

type GetTransactionIncomesResp struct {
	Items []TransactionIncome `json:"items"`
	Meta  types.Meta          `json:"meta"`
}

type GetTransactionIncomeReq struct {
	Id string `json:"id" validate:"required"`
}

type GetTransactionIncomeResp struct {
	TransactionIncome
}

type CreateTransactionIncomeReq struct {
	NoKontrak    string `json:"no_kontrak" validate:"required"`
	NoDocumentBc string `json:"no_document" db:"no_document"`
	KodeBarang   string `json:"kode_barang" validate:"required"`
	Jumlah       int    `json:"jumlah" validate:"min=0"`
}

type CreateTransactionIncomeResp struct {
	Id string `json:"id"`
}

type UpdateTransactionIncomeReq struct {
	Id         string `params:"id" validate:"required"`
	NoKontrak  string `json:"no_kontrak" validate:"required"`
	KodeBarang string `json:"kode_barang" validate:"required"`
	Jumlah     int    `json:"jumlah" validate:"min=0"`
}

type DeleteTransactionIncomeReq struct {
	Id string `json:"id" validate:"required"`
}
