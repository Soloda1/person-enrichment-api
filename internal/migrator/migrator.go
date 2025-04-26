package migrator

import (
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(log *slog.Logger, dbURL, migrationsPath string) {
	m, err := migrate.New("file://"+migrationsPath, dbURL)
	if err != nil {
		log.Error("Failed to create migrator", slog.String("error", err.Error()))
		panic("cannot create migrator: " + err.Error())
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("No new migrations to apply")
		} else {
			log.Error("Failed to apply migrations", slog.String("error", err.Error()))
			panic("cannot apply migrations: " + err.Error())
		}
	} else {
		log.Info("Migrations applied successfully")
	}
}
