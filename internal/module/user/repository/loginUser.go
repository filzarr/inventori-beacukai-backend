package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/user/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/jwthandler"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (r *userRepo) Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error) {
	type user struct {
		Id       string `db:"id"`
		Email    string `db:"email"`
		Password string `db:"password"`
		Role     string `db:"role"`
	}
	var (
		res    = new(entity.LoginResp)
		result = new(user)
	)

	query := `
		SELECT
			u.id,
			u.email,
			u.password,
			r.name as role
		FROM
			users u
		JOIN
			roles r ON r.id = u.role_id
		WHERE
			u.email = ?
			AND u.deleted_at IS NULL
	`

	err := r.db.GetContext(ctx, result, r.db.Rebind(query), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Err(err).Any("req", req.Log()).Msg("repo::Login - User not found")
			return nil, errmsg.NewCustomErrors(400).SetMessage("Kredensial yang Anda masukkan salah")
		}
		log.Error().Err(err).Any("req", req.Log()).Msg("repo::Login - Failed to get user")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		log.Warn().Err(err).Any("req", req.Log()).Msg("repo::Login - Password not match")
		return nil, errmsg.NewCustomErrors(400).SetMessage("Kredensial yang Anda masukkan salah")
	}

	// set token
	tokenExp := time.Now().UTC().Add(time.Hour * 24)
	payload := jwthandler.CostumClaimsPayload{
		UserId:          result.Id,
		Role:            result.Role,
		TokenExpiration: tokenExp,
	}

	token, err := jwthandler.GenerateTokenString(payload)
	if err != nil {
		log.Error().Err(err).Any("req", req.Log()).Msg("repo::Login - Failed to generate token")
		return nil, errmsg.NewCustomErrors(500).SetMessage("Gagal membuat token")
	}

	res.AccessToken = token

	return res, nil
}
