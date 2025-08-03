package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// bcDocumentsSeed seeds the bc_documents table.
func (s *Seed) bcDocumentsSeed() {
	bcDocumentMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ11", "kategori": "BC 2.3", "kode_document": "BC23-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ12", "kategori": "BC 2.5", "kode_document": "BC25-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ13", "kategori": "BC 2.6.1", "kode_document": "BC261-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ14", "kategori": "BC 2.7", "kode_document": "BC27-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ15", "kategori": "BC 4.0", "kode_document": "BC40-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ16", "kategori": "BC 4.1", "kode_document": "BC41-001-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ17", "kategori": "BC 2.3", "kode_document": "BC23-002-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ18", "kategori": "BC 2.5", "kode_document": "BC25-002-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ19", "kategori": "BC 2.6.1", "kode_document": "BC261-002-2024"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ20", "kategori": "BC 2.7", "kode_document": "BC27-002-2024"},
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
		INSERT INTO bc_documents (id, kategori, kode_document)
		VALUES (:id, :kategori, :kode_document)
	`, bcDocumentMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating bc_documents")
		return
	}

	log.Info().Msg("bc_documents table seeded successfully")
}