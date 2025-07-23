package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
)

type UserRepository interface {
	Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error)
	RegisterUser(ctx context.Context, req *entity.RegisterReq) (*entity.RegisterResp, error)
	GetProfile(ctx context.Context, req *entity.AuthListenReq) (*entity.AuthListenResp, error)
	GetRole(ctx context.Context, req *entity.GetRolesReq) (*entity.GetRolesResp, error)
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	ChangePassword(ctx context.Context, req *entity.ChangePasswordReq) error
	UpdateProfile(ctx context.Context, req *entity.UpdateProfileReq) error
	DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error
}

type UserService interface {
	Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error)
	RegisterUser(ctx context.Context, req *entity.RegisterReq) (*entity.RegisterResp, error)
	GetProfile(ctx context.Context, req *entity.AuthListenReq) (*entity.AuthListenResp, error)
	ChangePassword(ctx context.Context, req *entity.ChangePasswordReq) error
	GetUsers(ctx context.Context, req *entity.GetUsersReq) (*entity.GetUsersResp, error)
	GetRole(ctx context.Context, req *entity.GetRolesReq) (*entity.GetRolesResp, error)
	UpdateProfile(ctx context.Context, req *entity.UpdateProfileReq) error
	DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error
}
