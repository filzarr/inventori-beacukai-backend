package entity

type LoginReq struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginReq) Log() map[string]interface{} {
	return map[string]interface{}{
		"email": r.Email,
	}
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}
