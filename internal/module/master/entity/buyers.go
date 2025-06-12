package entity

import "inventori-beacukai-backend/pkg/types"

type GetBuyersReq struct {
	Q string `query:"q" validate:"omitempty,min=3"`
	types.MetaQuery
}

func (r *GetBuyersReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Buyer struct {
	Common
	Name   string `json:"name" db:"name"`
	Alamat string `json:"alamat" db:"alamat"`
}

type GetBuyersResp struct {
	Items []Buyer    `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type GetBuyerReq struct {
	Id string `json:"id" validate:"required"`
}

type GetBuyerResp struct {
	Buyer
}

type CreateBuyerReq struct {
	Name   string `json:"name" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=3"`
}

type CreateBuyerResp struct {
	Id string `json:"id"`
}

type UpdateBuyerReq struct {
	Id     string `params:"id" validate:"required"`
	Name   string `json:"name" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=3"`
}

type DeleteBuyerReq struct {
	Id string `json:"id" validate:"required"`
}
