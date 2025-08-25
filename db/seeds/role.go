package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// rolesSeed seeds the roles table.
func (s *Seed) rolesSeed() {
	roleMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYA", "name": "Admin"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYB", "name": "Gudang Bahan Baku"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYK", "name": "Gudang Bahan Penolong"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYC", "name": "Gudang Mesin"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYD", "name": "Produksi"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYF", "name": "Guest"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYG", "name": "Exim"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYH", "name": "Accounting"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYI", "name": "Timbangan"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZYJ", "name": "Beacukai"},
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
		INSERT INTO roles (id, name)
		VALUES (:id, :name)
	`, roleMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating roles")
		return
	}

	log.Info().Msg("roles table seeded successfully")
}
