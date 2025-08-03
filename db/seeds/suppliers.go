package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// suppliersSeed seeds the supliers table.
func (s *Seed) suppliersSeed() {
	supplierMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ41", "name": "PT Bahan Baku Prima", "alamat": "Jl. Industri No. 100, Tangerang, Banten 15134", "npwp": "11.234.567.8-901.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ42", "name": "CV Supplier Utama", "alamat": "Jl. Raya Bekasi No. 250, Bekasi, Jawa Barat 17141", "npwp": "12.345.678.9-012.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ43", "name": "PT Mesin Industri", "alamat": "Jl. Teknik No. 75, Karawang, Jawa Barat 41361", "npwp": "13.456.789.0-123.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ44", "name": "UD Kimia Sejahtera", "alamat": "Jl. Kimia Raya No. 45, Cilegon, Banten 42435", "npwp": "14.567.890.1-234.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ45", "name": "PT Logam Mulia", "alamat": "Jl. Logam No. 88, Sidoarjo, Jawa Timur 61215", "npwp": "15.678.901.2-345.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ46", "name": "CV Plastik Nusantara", "alamat": "Jl. Plastik Indah No. 123, Gresik, Jawa Timur 61111", "npwp": "16.789.012.3-456.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ47", "name": "PT Elektronik Global", "alamat": "Jl. Elektronik No. 200, Batam, Kepulauan Riau 29444", "npwp": "17.890.123.4-567.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ48", "name": "UD Tekstil Mandiri", "alamat": "Jl. Tekstil No. 300, Solo, Jawa Tengah 57126", "npwp": "18.901.234.5-678.000"},
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
		INSERT INTO supliers (id, name, alamat, npwp)
		VALUES (:id, :name, :alamat, :npwp)
	`, supplierMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating suppliers")
		return
	}

	log.Info().Msg("suppliers table seeded successfully")
}