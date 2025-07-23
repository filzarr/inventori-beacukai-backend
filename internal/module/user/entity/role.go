package entity

type GetRolesReq struct {
	UserId string `json:"user_id"`
}

type Role struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GetRolesResp struct {
	Items []Role `json:"items"`
}
