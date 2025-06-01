package adapter

import (
	// "log"

	"inventori-beacukai-backend/internal/infrastructure/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func WithPostgres() Option {
	return func(a *Adapter) {
		dbUser := config.Envs.Postgres.Username
		dbPassword := config.Envs.Postgres.Password
		dbName := config.Envs.Postgres.Database
		dbHost := config.Envs.Postgres.Host
		dbSSLMode := config.Envs.Postgres.SslMode
		dbPort := config.Envs.Postgres.Port

		dbMaxPoolSize := config.Envs.DB.MaxOpenCons
		dbMaxIdleConns := config.Envs.DB.MaxIdleCons
		dbConnMaxLifetime := config.Envs.DB.ConnMaxLifetime

		connectionString := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=" + dbSSLMode + " TimeZone=UTC"
		db, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to Postgres")
		}

		db.SetMaxOpenConns(dbMaxPoolSize)
		db.SetMaxIdleConns(dbMaxIdleConns)
		db.SetConnMaxLifetime(time.Duration(dbConnMaxLifetime) * time.Second)

		// check connection
		err = db.Ping()
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to Digihub Postgres")
		}

		a.Postgres = db
		log.Info().Msg("Postgres connected")
	}
}
