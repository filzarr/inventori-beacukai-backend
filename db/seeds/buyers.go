package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// buyersSeed seeds the buyers table.
func (s *Seed) buyersSeed() {
	buyerMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ31", "name": "PT Maju Bersama", "alamat": "Jl. Sudirman No. 123, Jakarta Pusat, DKI Jakarta 10220", "npwp": "01.234.567.8-901.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ32", "name": "CV Sukses Mandiri", "alamat": "Jl. Gatot Subroto No. 456, Bandung, Jawa Barat 40123", "npwp": "02.345.678.9-012.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ33", "name": "PT Global Trading", "alamat": "Jl. Ahmad Yani No. 789, Surabaya, Jawa Timur 60234", "npwp": "03.456.789.0-123.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ34", "name": "UD Berkah Jaya", "alamat": "Jl. Diponegoro No. 321, Yogyakarta, DI Yogyakarta 55234", "npwp": "04.567.890.1-234.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ35", "name": "PT Nusantara Export", "alamat": "Jl. Imam Bonjol No. 654, Medan, Sumatera Utara 20114", "npwp": "05.678.901.2-345.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ36", "name": "CV Harapan Baru", "alamat": "Jl. Veteran No. 987, Semarang, Jawa Tengah 50231", "npwp": "06.789.012.3-456.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ37", "name": "PT Indo Pacific", "alamat": "Jl. Thamrin No. 147, Jakarta Selatan, DKI Jakarta 12930", "npwp": "07.890.123.4-567.000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ38", "name": "UD Sinar Terang", "alamat": "Jl. Pahlawan No. 258, Malang, Jawa Timur 65145", "npwp": "08.901.234.5-678.000"},
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
		INSERT INTO supliers (id, name, kategori_supliers, alamat, npwp)
		VALUES (:id, :name, 'Pembeli', :alamat, :npwp)
	`, buyerMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating buyers")
		return
	}

	log.Info().Msg("buyers table seeded successfully")
}
