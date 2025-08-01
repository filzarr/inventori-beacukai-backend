package entity

import (
	"inventori-beacukai-backend/pkg/types"
)

type GetContractsReq struct {
	Q        string `query:"q" validate:"omitempty,min=3"`
	Document bool   `query:"document"`
	types.MetaQuery
}

func (r *GetContractsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Contract struct {
	Common
	Kategori          string  `json:"kategori" db:"kategori_kontrak"`
	NamaPemasok       string  `json:"nama_pemasok" db:"nama_pemasok"`
	KodeDocumentBC    *string `json:"kode_document_bc" db:"kode_document_bc"`
	TanggalDocumentBc *string `json:"tanggal_document_bc" db:"tanggal_document_bc"`
	AlamatPemasok     string  `json:"alamat_pemasok" db:"alamat_pemasok"`
	NoKontrak         string  `json:"no_kontrak" db:"no_kontrak"`
	Tanggal           string  `json:"tanggal" db:"tanggal"`
}

type GetContractsResp struct {
	Items []Contract `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetContractReq struct {
	Id string `json:"id" validate:"required"`
}

type GetContractResp struct {
	Contract
}

type CreateContractReq struct {
	NoKontrak      string  `json:"no_kontrak" validate:"required,min=3"`
	SupliersId     string  `json:"supliers" validate:"required"`
	KodeDocumentBC *string `json:"kode_document_bc"`
	Kategori       string  `json:"kategori" validate:"required,min=3"`
	Tanggal        string  `json:"tanggal" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type CreateContractResp struct {
	Id string `json:"id"`
}

type UpdateContractReq struct {
	Id         string `params:"id" validate:"required"`
	SupliersId string `json:"supliers" validate:"required"`
	NoKontrak  string `json:"no_kontrak" validate:"required,min=3"`
	Kategori   string `json:"kategori" validate:"required,min=3"`
	Tanggal    string `json:"tanggal" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type DeleteContractReq struct {
	Id string `json:"id" validate:"required"`
}

type UpdateContractDocumentReq struct {
	NoKontrak       string `json:"no_kontrak" validate:"required"`
	NoDocumentBc    string `json:"no_document" validate:"required"`
	TanggalDocument string `json:"tanggal_document" db:"tanggal_document"`
}

type Transaction struct {
	Common
	Kategori        string `json:"kategori" db:"kategori"`
	KodeDocument    string `json:"kode_document" db:"kode_document"`
	TanggalDocument string `json:"tanggal_document" db:"tanggal_document"`
	NoKontrak       string `json:"no_kontrak" db:"no_kontrak"`
}

type GetTransactionsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

type GetTransactionsResp struct {
	Items []Transaction `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type GetContractNotRequiredReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

type GetContractNotRequiredResp struct {
	Items []ContractNotRequired `json:"items"`
	Meta  types.Meta            `json:"meta"`
}

type ContractNotRequired struct {
	Common
	NoKontrak string `json:"no_kontrak" db:"no_kontrak"`
}

func (r *GetContractNotRequiredReq) SetDefault() {
	r.MetaQuery.SetDefault()
}
