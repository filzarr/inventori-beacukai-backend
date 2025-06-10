package seeds

import (
	"context"
	"inventori-beacukai-backend/pkg"

	"github.com/rs/zerolog/log"
)

func (s *Seed) userSeed() {
	hash, err := pkg.HashPassword("superadmin")
	if err != nil {
		log.Error().Err(err).Msg("Error hash password")
	}
	userMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ10", "name": "Super Admin", "email": "superadmin@gmail.com", "password": hash, "superadmin": "TRUE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZY9"},
	}

	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	_, err = tx.NamedExec(`
		INSERT INTO users (id, name, email, password, superadmin, role_id)
		VALUES (:id, :name, :email, :password,:superadmin, :role_id)
	`, userMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating roles")
		return
	}

	log.Info().Msg("roles table seeded successfully")
}
