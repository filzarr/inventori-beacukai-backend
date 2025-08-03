package seeds

import (
	"context"
	"inventori-beacukai-backend/internal/adapter"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// Seed struct.
type Seed struct {
	db *sqlx.DB
}

// NewSeed return a Seed with a pool of connection to a dabase.
func newSeed(db *sqlx.DB) Seed {
	return Seed{
		db: db,
	}
}

func Execute(db *sqlx.DB, table string, total int) {
	seed := newSeed(db)
	seed.run(table, total)
}

// Run seeds.
func (s *Seed) run(table string, total int) {

	switch table {
	case "roles":
		s.rolesSeed()
	case "user":
		s.userSeed()
	case "mata_uang":
		s.mataUangSeed()
	case "bc_documents":
		s.bcDocumentsSeed()
	case "warehouses":
		s.warehousesSeed()
	case "buyers":
		s.buyersSeed()
	case "suppliers":
		s.suppliersSeed()
	case "products":
		s.productsSeed()
	case "saldo_awal":
		s.saldoAwalSeed()
	case "all":
		s.rolesSeed()
		s.userSeed()
		s.mataUangSeed()
		s.bcDocumentsSeed()
		s.warehousesSeed()
		s.buyersSeed()
		s.suppliersSeed()
		s.productsSeed()
		s.saldoAwalSeed()
	case "delete-all":
		s.deleteAll()
	default:
		log.Warn().Msg("No seed to run")
	}

	if table != "" {
		log.Info().Msg("Seed ran successfully")
		log.Info().Msg("Exiting ...")
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
		os.Exit(0)
	}
}

func (s *Seed) deleteAll() {
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

	// Delete in reverse order to respect foreign key constraints
	_, err = tx.Exec(`DELETE FROM saldo_awals`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting saldo_awals")
		return
	}
	log.Info().Msg("saldo_awals table deleted successfully")

	_, err = tx.Exec(`DELETE FROM products`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting products")
		return
	}
	log.Info().Msg("products table deleted successfully")

	_, err = tx.Exec(`DELETE FROM warehouses`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting warehouses")
		return
	}
	log.Info().Msg("warehouses table deleted successfully")

	_, err = tx.Exec(`DELETE FROM buyers`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting buyers")
		return
	}
	log.Info().Msg("buyers table deleted successfully")

	_, err = tx.Exec(`DELETE FROM supliers`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting suppliers")
		return
	}
	log.Info().Msg("suppliers table deleted successfully")

	_, err = tx.Exec(`DELETE FROM bc_documents`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting bc_documents")
		return
	}
	log.Info().Msg("bc_documents table deleted successfully")

	_, err = tx.Exec(`DELETE FROM currencies`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting currencies")
		return
	}
	log.Info().Msg("currencies table deleted successfully")

	_, err = tx.Exec(`DELETE FROM roles`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting roles")
		return
	}
	log.Info().Msg("roles table deleted successfully")

	log.Info().Msg("=== All tables deleted successfully ===")
}
