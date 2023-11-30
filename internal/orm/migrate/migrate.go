package migrate

import (
	"embed"
	"os"
	"strconv"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// MigrationsDir migration dir
const MigrationsDir string = "migrations"

func init() {
	goose.SetBaseFS(embedMigrations)
	goose.SetSequential(true)
	goose.SetTableName("chain_monitor_v2_migrations")

	verbose, _ := strconv.ParseBool(os.Getenv("LOG_SQL_MIGRATIONS"))
	goose.SetVerbose(verbose)
}

// Migrate migrate db
func Migrate(db *gorm.DB) error {
	tx, err := db.DB()
	if err != nil {
		return err
	}
	return goose.Up(tx, MigrationsDir, goose.WithAllowMissing())
}

// ResetDB clean and migrate db.
func ResetDB(db *gorm.DB) error {
	if err := Rollback(db, 0); err != nil {
		return err
	}
	return Migrate(db)
}

// Rollback to the given version
func Rollback(db *gorm.DB, version int64) error {
	tx, err := db.DB()
	if err != nil {
		return err
	}
	return goose.DownTo(tx, MigrationsDir, version)
}
