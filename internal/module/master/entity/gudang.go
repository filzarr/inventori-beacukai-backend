package entity

import "inventori-beacukai-backend/pkg/types"

type GetGudangsReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetGudangsReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Gudang struct {
	Common
	Nama   string `json:"nama" db:"nama"`
	Lokasi string `json:"lokasi" db:"lokasi"`
	User   string `json:"user" db:"user"`
}

type GetGudangsResp struct {
	Items []Gudang   `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetGudangReq struct {
	Id string `json:"id" validate:"required"`
}

type GetGudangResp struct {
	Gudang
}

type CreateGudangReq struct {
	Nama   string `json:"nama" validate:"required"`
	Lokasi string `json:"lokasi" validate:"required"`
	User   string `json:"user" validate:"required"`
}

type CreateGudangResp struct {
	Id string `json:"id"`
}

type UpdateGudangReq struct {
	Id     string `json:"id" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Lokasi string `json:"lokasi" validate:"required"`
	User   string `json:"user" validate:"required"`
}

type DeleteGudangReq struct {
	Id string `json:"id" validate:"required"`
}
