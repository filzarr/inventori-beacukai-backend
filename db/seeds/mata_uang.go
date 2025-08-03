package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// mataUangSeed seeds the currencies table.
func (s *Seed) mataUangSeed() {
	currencyMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ01", "kode": "IDR", "mata_uang": "Indonesian Rupiah"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ02", "kode": "USD", "mata_uang": "United States Dollar"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ03", "kode": "EUR", "mata_uang": "Euro"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ04", "kode": "JPY", "mata_uang": "Japanese Yen"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ05", "kode": "SGD", "mata_uang": "Singapore Dollar"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ06", "kode": "CNY", "mata_uang": "Chinese Yuan"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ07", "kode": "MYR", "mata_uang": "Malaysian Ringgit"},
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
		INSERT INTO currencies (id, kode, mata_uang)
		VALUES (:id, :kode, :mata_uang)
	`, currencyMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating currencies")
		return
	}

	log.Info().Msg("currencies table seeded successfully")
}