package orm

import (
	"sync"

	"gorm.io/gorm"
)

var (
	// L1Tables tables for L1 events.
	L1Tables = []interface{}{
		&L1Block{},
		&L1ETHEvent{},
		&L1ERC20Event{},
		&L1ERC721Event{},
		&L1ERC1155Event{},
		&L1MessengerEvent{},
		&L1ScrollChainEvent{},
	}
	// L2Tables tables for L2 events.
	L2Tables = []interface{}{
		&L2Block{},
		&L2ETHEvent{},
		&L2ERC20Event{},
		&L2ERC721Event{},
		&L2ERC1155Event{},
		&L2MessengerEvent{},
	}
	tables = []interface{}{
		&ChainConfirm{},
	}
	once sync.Once
)

func init() {
	once.Do(func() {
		tables = append(tables, L1Tables...)
		tables = append(tables, L2Tables...)
	})
}

// CreateTables auto migrates the given tables in the database.
func CreateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(tables...); err != nil {
		return err
	}
	return nil
}

// DropTables drops the given tables from the database if they exist.
func DropTables(db *gorm.DB) error {
	migrator := db.Migrator()
	for _, tb := range tables {
		if migrator.HasTable(tb) {
			if err := migrator.DropTable(tb); err != nil {
				return err
			}
		}
	}
	return nil
}
