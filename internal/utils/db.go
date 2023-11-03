package utils

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"chain-monitor/internal/config"
)

// InitDB init the db handler
func InitDB(config *config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.LogLevel)),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenNum)
	sqlDB.SetMaxIdleConns(config.MaxIdleNum)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDB close the db handler. notice the db handler only can close when then program exit.
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

// Ping check db status
func Ping(db *gorm.DB) (*sql.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	return sqlDB, nil
}
