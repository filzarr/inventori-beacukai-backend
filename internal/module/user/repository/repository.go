package repository

import (
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/user/ports"

	"github.com/jmoiron/sqlx"
)

var _ ports.UserRepository = &userRepo{}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository() *userRepo {
	return &userRepo{
		db: adapter.Adapters.Postgres,
	}
}
