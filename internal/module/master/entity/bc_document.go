package entity

import "inventori-beacukai-backend/pkg/types"

type GetBcDocumentsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetBcDocumentsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type BcDocument struct {
	Common
	Kategori   string `json:"kategori" db:"kategori"`
	NoDocument string `json:"no_document" db:"no_document"`
	Tanggal    string `json:"tanggal" db:"tanggal"`
}

type GetBcDocumentsResp struct {
	Items []BcDocument `json:"items"`
	Meta  types.Meta   `json:"meta"`
}

type GetBcDocumentReq struct {
	Id string `json:"id" validate:"required"`
}

type GetBcDocumentResp struct {
	BcDocument
}

type CreateBcDocumentReq struct {
	Kategori   string `json:"kategori" validate:"required,oneof='BC 23' 'BC 27 In' 'BC 262' 'BC 40'"`
	NoDocument string `json:"no_document" validate:"required,min=3"`
	Tanggal    string `json:"tanggal" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type CreateBcDocumentResp struct {
	Id string `json:"id"`
}

type UpdateBcDocumentReq struct {
	Id         string `params:"id" validate:"required"`
	Kategori   string `json:"kategori" validate:"required,oneof='BC 23' 'BC 27 In' 'BC 262' 'BC 40'"`
	NoDocument string `json:"no_document" validate:"required,min=3"`
	Tanggal    string `json:"tanggal" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type DeleteBcDocumentReq struct {
	Id string `json:"id" validate:"required"`
}
