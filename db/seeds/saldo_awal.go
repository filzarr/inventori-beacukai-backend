package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// saldoAwalSeed seeds the saldo_awals table.
func (s *Seed) saldoAwalSeed() {
	saldoAwalMaps := []map[string]any{
		// Saldo awal untuk beberapa produk
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ71", "kode_barang": "BB-001", "saldo_awal": "1000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ72", "kode_barang": "BB-002", "saldo_awal": "500"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ73", "kode_barang": "BB-003", "saldo_awal": "100"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ74", "kode_barang": "BB-004", "saldo_awal": "80"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ75", "kode_barang": "BB-005", "saldo_awal": "200"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ76", "kode_barang": "BP-001", "saldo_awal": "50"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ77", "kode_barang": "BP-002", "saldo_awal": "100"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ78", "kode_barang": "BP-003", "saldo_awal": "75"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ79", "kode_barang": "MS-001", "saldo_awal": "10"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ80", "kode_barang": "MS-002", "saldo_awal": "50"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ81", "kode_barang": "BJ-001", "saldo_awal": "200"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ82", "kode_barang": "BJ-002", "saldo_awal": "150"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ83", "kode_barang": "BJ-003", "saldo_awal": "5000"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ84", "kode_barang": "BJ-004", "saldo_awal": "100"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ85", "kode_barang": "BJ-005", "saldo_awal": "50"},
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
		INSERT INTO saldo_awals (id, kode_barang, saldo_awal)
		VALUES (:id, :kode_barang, :saldo_awal)
	`, saldoAwalMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating saldo_awals")
		return
	}

	log.Info().Msg("saldo_awals table seeded successfully")
}