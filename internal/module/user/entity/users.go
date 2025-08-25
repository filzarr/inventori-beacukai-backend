package entity

import "inventori-beacukai-backend/pkg/types"

type GetUsersReq struct {
	Q string `query:"q" validate:"omitempty"`
	types.MetaQuery
}

func (r *GetUsersReq) SetDefault() {
	r.MetaQuery.SetDefault()
}

type Users struct {
	Id    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Role  string `json:"role" db:"role"`
}

type GetUsersResp struct {
	Items []Users    `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type DeleteUserReq struct {
	UserId string `json:"user_id" validate:"required"`

	Id string `json:"id" validate:"required"`
}
