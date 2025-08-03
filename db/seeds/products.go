package seeds

import (
	"context"

	"github.com/rs/zerolog/log"
)

// productsSeed seeds the products table.
func (s *Seed) productsSeed() {
	productMaps := []map[string]any{
		// Bahan Baku
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ51", "kode": "BB-001", "nama": "Besi Beton 12mm", "satuan": "Batang", "kategori": "Bahan Baku", "saldo_awal": 1000, "jumlah": 1000},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ52", "kode": "BB-002", "nama": "Semen Portland", "satuan": "Sak", "kategori": "Bahan Baku", "saldo_awal": 500, "jumlah": 500},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ53", "kode": "BB-003", "nama": "Pasir Halus", "satuan": "M3", "kategori": "Bahan Baku", "saldo_awal": 100, "jumlah": 100},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ54", "kode": "BB-004", "nama": "Kerikil", "satuan": "M3", "kategori": "Bahan Baku", "saldo_awal": 80, "jumlah": 80},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ55", "kode": "BB-005", "nama": "Kawat Bendrat", "satuan": "Kg", "kategori": "Bahan Baku", "saldo_awal": 200, "jumlah": 200},
		
		// Bahan Penolong
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ56", "kode": "BP-001", "nama": "Cat Primer", "satuan": "Kaleng", "kategori": "Bahan Penolong", "saldo_awal": 50, "jumlah": 50},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ57", "kode": "BP-002", "nama": "Thinner", "satuan": "Liter", "kategori": "Bahan Penolong", "saldo_awal": 100, "jumlah": 100},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ58", "kode": "BP-003", "nama": "Lem Epoxy", "satuan": "Tube", "kategori": "Bahan Penolong", "saldo_awal": 75, "jumlah": 75},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ59", "kode": "BP-004", "nama": "Silikon Sealant", "satuan": "Tube", "kategori": "Bahan Penolong", "saldo_awal": 60, "jumlah": 60},
		
		// Mesin/Sparepart
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ60", "kode": "MS-001", "nama": "Motor Listrik 5HP", "satuan": "Unit", "kategori": "Mesin/Sparepart", "saldo_awal": 10, "jumlah": 10},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ61", "kode": "MS-002", "nama": "Bearing 6205", "satuan": "Pcs", "kategori": "Mesin/Sparepart", "saldo_awal": 50, "jumlah": 50},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ62", "kode": "MS-003", "nama": "V-Belt A-50", "satuan": "Pcs", "kategori": "Mesin/Sparepart", "saldo_awal": 25, "jumlah": 25},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ63", "kode": "MS-004", "nama": "Gear Box 1:20", "satuan": "Unit", "kategori": "Mesin/Sparepart", "saldo_awal": 5, "jumlah": 5},
		
		// Barang Jadi
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ64", "kode": "BJ-001", "nama": "Panel Beton Precast", "satuan": "Panel", "kategori": "Barang Jadi", "saldo_awal": 200, "jumlah": 200},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ65", "kode": "BJ-002", "nama": "Tiang Pancang 30cm", "satuan": "Batang", "kategori": "Barang Jadi", "saldo_awal": 150, "jumlah": 150},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ66", "kode": "BJ-003", "nama": "Paving Block 20x20", "satuan": "Pcs", "kategori": "Barang Jadi", "saldo_awal": 5000, "jumlah": 5000},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ67", "kode": "BJ-004", "nama": "U-Ditch 40x40x120", "satuan": "Pcs", "kategori": "Barang Jadi", "saldo_awal": 100, "jumlah": 100},
		{"id": "01J3X0H3TK5MWPXK9D8GDRFZ68", "kode": "BJ-005", "nama": "Box Culvert 100x100", "satuan": "Pcs", "kategori": "Barang Jadi", "saldo_awal": 50, "jumlah": 50},
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
		INSERT INTO products (id, kode, nama, satuan, kategori, saldo_awal, jumlah)
		VALUES (:id, :kode, :nama, :satuan, :kategori, :saldo_awal, :jumlah)
	`, productMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating products")
		return
	}

	log.Info().Msg("products table seeded successfully")
}