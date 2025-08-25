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
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ10", "name": "Super Admin", "email": "superadmin@gmail.com", "password": hash, "superadmin": "TRUE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYA"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ11", "name": "Admin Bahan Baku", "email": "bahanbaku@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYB"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ12", "name": "Admin Bahan Penolong", "email": "penolong@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYK"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ13", "name": "Admin Mesin", "email": "mesin@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYC"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ14", "name": "Staff Produksi", "email": "produksi@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYD"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ16", "name": "Staff Exim", "email": "exim@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYG"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ17", "name": "Staff Accounting", "email": "accounting@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYH"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ18", "name": "Staff Timbangan", "email": "timbangan@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYI"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ19", "name": "Staff Beacukai", "email": "beacukai@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYJ"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRVZ20", "name": "Guest User", "email": "guest@gmail.com", "password": hash, "superadmin": "FALSE", "role_id": "01J3X0H3TK5MWPXK9D8GDRFZYF"},
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
		log.Error().Err(err).Msg("Error creating users")
		return
	}

	log.Info().Msg("users table seeded successfully")
}
