package repository

import (
	"context"
	"database/sql"
	"inventori-beacukai-backend/internal/module/user/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"

	"github.com/rs/zerolog/log"
)

func (r *userRepo) GetProfile(ctx context.Context, req *entity.AuthListenReq) (*entity.AuthListenResp, error) {
	var (
		resp = new(entity.AuthListenResp)
	)

	query := `SELECT u.id, u.name, u.email, u.superadmin, r.name AS role , u.role_id FROM users u JOIN roles r ON u.role_id = r.id WHERE u.id = ? AND u.deleted_at IS NULL`

	if err := r.db.GetContext(ctx, resp, r.db.Rebind(query), req.Id); err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Any("req", req).Msg("repo::GetBcDocument - not found")
			return nil, errmsg.NewCustomErrors(404).SetMessage("Dokumen BC tidak ditemukan")
		}
		log.Error().Err(err).Any("req", req).Msg("repo::GetProfile - failed to get profile")
		return nil, err
	}

	return resp, nil
}

func (r *userRepo) UpdateProfile(ctx context.Context, req *entity.UpdateProfileReq) error {
	query := `UPDATE users SET name = ?, email = ?, role_id = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS null`

	if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Name, req.Email, req.Role, req.Id); err != nil {
		log.Error().Err(err).Any("req", req).Msg("repo::UpdateProfile - failed to update")
		return err
	}
	return nil
}
