package entity

type ChangePasswordReq struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
