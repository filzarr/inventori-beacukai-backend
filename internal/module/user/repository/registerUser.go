package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
	"inventori-beacukai-backend/pkg"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func (r *userRepo) RegisterUser(ctx context.Context, req *entity.RegisterReq) (*entity.RegisterResp, error) {
	query := `
		INSERT INTO users (
			id,
			name,
			email,
			password,
			role_id
		) VALUES (?, ?, ?, ?, ?)
	`

	var (
		Id   = ulid.Make().String()
		resp = new(entity.RegisterResp)
	)
	pass := pkg.GeneratePassword(10)
	hash, err := pkg.HashPassword(pass)
	if err != nil {
		log.Error().Err(err).Msg("repo::registeruser - Error hash password")
	}
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query),
		Id, req.Name, req.Email, hash, req.Role); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::CreateUser - failed to create User")
		return nil, err
	}

	resp.Id = Id
	resp.Email = req.Email
	resp.Name = req.Name
	resp.Password = pass
	return resp, nil
}
