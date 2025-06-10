package entity

import "inventori-beacukai-backend/pkg/types"

type GetSupliersReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetSupliersReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Suplier struct {
	Common
	Name   string `json:"name" db:"name"`
	Alamat string `json:"alamat" db:"alamat"`
}

type GetSupliersResp struct {
	Items []Suplier  `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetSuplierReq struct {
	Id string `json:"id" validate:"required"`
}

type GetSuplierResp struct {
	Suplier
}

type CreateSuplierReq struct {
	Name   string `json:"name" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=5"`
}

type CreateSuplierResp struct {
	Id string `json:"id"`
}

type UpdateSuplierReq struct {
	Id     string `params:"id" validate:"required"`
	Name   string `json:"name" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=5"`
}

type DeleteSuplierReq struct {
	Id string `json:"id" validate:"required"`
}
