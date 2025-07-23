package repository

import (
	"context"
	"inventori-beacukai-backend/internal/module/user/entity"
	"inventori-beacukai-backend/pkg"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (r *userRepo) ChangePassword(ctx context.Context, req *entity.ChangePasswordReq) error {
	type user struct {
		Password string `db:"password"`
	}
	var fetchUser user

	query := `SELECT password FROM users WHERE id = ?`
	if err := r.db.GetContext(ctx, &fetchUser, r.db.Rebind(query), req.UserId); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::ChangePassword - failed to fetch user")
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(fetchUser.Password), []byte(req.OldPassword)); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("repo::ChangePassword - old password does not match")
		return errmsg.NewCustomErrors(400).SetMessage("Password lama salah")
	}

	newPassword, err := pkg.HashPassword(req.NewPassword)
	if err != nil {
		log.Error().Err(err).Msg("repo::ChangePassword - error hashing new password")
		return errmsg.NewCustomErrors(500).SetMessage("Gagal mengenkripsi password")
	}

	updateQuery := `UPDATE users SET password = ?, updated_at = NOW() WHERE id = ?`
	if _, err := r.db.ExecContext(ctx, r.db.Rebind(updateQuery), newPassword, req.UserId); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::ChangePassword - failed to update password")
		return errmsg.NewCustomErrors(500).SetMessage("Gagal mengubah password")
	}
	return nil
}
