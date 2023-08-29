package orm

import "gorm.io/gorm"

var (
	tables = []interface{}{
		&L1Block{},
		&L2Block{},

		&L1ETHEvent{},
		&L2ETHEvent{},
		&L1ERC20Event{},
		&L2ERC20Event{},
		&L1ERC721Event{},
		&L2ERC721Event{},
		&L1ERC1155Event{},
		&L2ERC1155Event{},
		//&L1MessengerEvent{},
		//&L2MessengerEvent{},
		&ScrollChainEvent{},

		&BatchConfirm{},
	}
)

func CreateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(tables...); err != nil {
		return err
	}
	return nil
}

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
