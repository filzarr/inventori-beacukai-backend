package entity

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginReq) Log() map[string]interface{} {
	return map[string]interface{}{
		"email": r.Email,
	}
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type RegisterReq struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"required"`
}

func (r *RegisterReq) Log() map[string]interface{} {
	return map[string]interface{}{
		"email": r.Email,
		"Name":  r.Name,
		"Role":  r.Role,
	}
}

type RegisterResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	Id string `json:"id"`
}

type AuthListenReq struct {
	Id string `json:"id"`
}
type AuthListenResp struct {
	Id         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Role       string `json:"role" db:"role"`
	SuperAdmin bool   `json:"superadmin" db:"superadmin"`
	RoleId     string `json:"role_id" db:"role_id"`
}
