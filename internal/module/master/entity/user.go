package entity

import "inventori-beacukai-backend/pkg/types"

type GetUsersReq struct {
	Q string `query:"q" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetUsersReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type User struct {
	Common
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Role  string `json:"role" db:"role"`
}

type GetUsersResp struct {
	Items []User     `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type DeleteLecturerReq struct {
	Id string `json:"id" validate:"required"`
}
