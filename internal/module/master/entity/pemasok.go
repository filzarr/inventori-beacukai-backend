package entity

import "inventori-beacukai-backend/pkg/types"

type GetPemasoksReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetPemasoksReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Pemasok struct {
	Common
	Nama   string `json:"nama" db:"nama"`
	Alamat string `json:"alamat" db:"alamat"`
	NoHp   string `json:"noHp" db:"noHp"`
}

type GetPemasoksResp struct {
	Items []Pemasok  `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetPemasokReq struct {
	Id string `json:"id" validate:"required"`
}

type GetPemasokResp struct {
	Pemasok
}

type CreatePemasokReq struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp   string `json:"noHp" validate:"required"`
}

type CreatePemasokResp struct {
	Id string `json:"id"`
}

type UpdatePemasokReq struct {
	Id     string `json:"id" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp   string `json:"noHp" validate:"required"`
}

type DeletePemasokReq struct {
	Id string `json:"id" validate:"required"`
}
