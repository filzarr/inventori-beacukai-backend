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
	Kategori     string `json:"kategori" db:"kategori"`
	KodeDocument string `json:"kode_document" db:"kode_document"`
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
	Kategori     string `json:"kategori" validate:"required"`
	KodeDocument string `json:"kode_document" validate:"required,min=3"`
}

type CreateBcDocumentResp struct {
	Id string `json:"id"`
}

type UpdateBcDocumentReq struct {
	Id           string `params:"id" validate:"required"`
	Kategori     string `json:"kategori" validate:"required"`
	KodeDocument string `json:"kode_document" validate:"required,min=3"`
}

type DeleteBcDocumentReq struct {
	Id string `json:"id" validate:"required"`
}
