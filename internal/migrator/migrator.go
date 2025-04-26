package migrator

import (
	"errors"
	"log/slog"
	"person-enrichment-api/internal/utils/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(log *logger.Logger, dbURL, migrationsPath string) error {
	m, err := migrate.New("file://"+migrationsPath, dbURL)
	if err != nil {
		log.Error("Failed to create migrator", slog.String("error", err.Error()))
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("No new migrations to apply")
			return nil
		}
		log.Error("Failed to apply migrations", slog.String("error", err.Error()))
		return err
	}

	log.Info("Migrations applied successfully")
	return nil
}
