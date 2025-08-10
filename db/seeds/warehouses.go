package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// warehousesSeed seeds the warehouses table.
func (s *Seed) warehousesSeed() {
	warehouseMaps := []map[string]any{
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ21", "kode": "WH-BB-001", "nama": "Gudang Bahan Baku Utama", "kategori": "Bahan Baku", "keterangan": "Gudang penyimpanan bahan baku utama produksi"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ22", "kode": "WH-BP-001", "nama": "Gudang Bahan Penolong", "kategori": "Bahan Penolong", "keterangan": "Gudang penyimpanan bahan penolong dan tambahan"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ23", "kode": "WH-MS-001", "nama": "Gudang Mesin dan Sparepart", "kategori": "Mesin/Sparepart", "keterangan": "Gudang penyimpanan mesin dan komponen sparepart"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ24", "kode": "WH-BJ-001", "nama": "Gudang Barang Jadi A", "kategori": "Barang Jadi", "keterangan": "Gudang penyimpanan barang jadi siap kirim area A"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ25", "kode": "WH-BJ-002", "nama": "Gudang Barang Jadi B", "kategori": "Barang Jadi", "keterangan": "Gudang penyimpanan barang jadi siap kirim area B"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ26", "kode": "WH-BB-002", "nama": "Gudang Bahan Baku Cadangan", "kategori": "Bahan Baku", "keterangan": "Gudang cadangan untuk bahan baku overflow"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ27", "kode": "WH-MS-002", "nama": "Gudang Tools dan Equipment", "kategori": "Mesin/Sparepart", "keterangan": "Gudang penyimpanan tools dan equipment produksi"},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ28", "kode": "WH-PD-001", "nama": "Gudang Produksi 1", "kategori": "Produksi", "keterangan": "Gudang produksi minyak"},
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
		INSERT INTO warehouses (id, kode, nama, kategori, keterangan)
		VALUES (:id, :kode, :nama, :kategori, :keterangan)
	`, warehouseMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating warehouses")
		return
	}

	log.Info().Msg("warehouses table seeded successfully")
}
