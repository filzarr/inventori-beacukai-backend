package entity

import "inventori-beacukai-backend/pkg/types"

type GetContractsBcReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetContractsBcReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type ContractsBc struct {
	Common
	NoKontrak         string `json:"no_kontrak" db:"no_kontrak"`
	KodeDocumentBc    string `json:"kode_document_bc" db:"kode_document_bc"`
	NomorDocumentBc   string `json:"nomor_document_bc" db:"nomor_document_bc"`
	TanggalDocumentBc string `json:"tanggal_document_bc" db:"tanggal_document_bc"`
}

type GetContractsBcResp struct {
	Items []ContractsBc `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type GetContractBcReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}

type GetContractBcResp struct {
	ContractsBc
}

type CreateContractBcReq struct {
	UserId string `validate:"ulid"`

	NoKontrak         string `json:"no_kontrak" validate:"required"`
	KodeDocumentBc    string `json:"kode_document_bc" validate:"required"`
	NomorDocumentBc   string `json:"nomor_document_bc" validate:"required"`
	TanggalDocumentBc string `json:"tanggal_document_bc" validate:"required,min=0"`
}

type CreateContractBcResp struct {
	Id string `json:"id"`
}

type UpdateContractBcReq struct {
	UserId string `validate:"ulid"`

	Id                string `params:"id" validate:"required"`
	NoKontrak         string `json:"no_kontrak" validate:"required"`
	KodeDocumentBc    string `json:"kode_document_bc" validate:"required"`
	NomorDocumentBc   string `json:"nomor_document_bc" validate:"required"`
	TanggalDocumentBc string `json:"tanggal_document_bc" validate:"required,min=0"`
}

type DeleteContractBcReq struct {
	UserId string `validate:"ulid"`

	Id string `json:"id" validate:"required"`
}
