package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up000018, Down000018)
}

// Up000018 Change gateway and messenger event_type.
func Up000018(ctx context.Context, tx *sql.Tx) error {
	// Update l1_erc20_gateway
	_, err := tx.ExecContext(ctx, "UPDATE l1_erc20_events SET type = type + 10;")
	if err != nil {
		return err
	}
	// Update l2_erc20_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc20_events SET type = type + 10;")
	if err != nil {
		return err
	}

	// Update l1_erc721_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l1_erc721_events SET type = type + 20;")
	if err != nil {
		return err
	}
	// Update l2_erc721_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc721_events SET type = type + 20;")
	if err != nil {
		return err
	}

	// Update l1_erc1155_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l1_erc1155_events SET type = type + 30;")
	if err != nil {
		return err
	}
	// Update l2_erc1155_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc1155_events SET type = type + 30;")
	if err != nil {
		return err
	}

	// Update l1_messenger_events
	_, err = tx.ExecContext(ctx, "UPDATE l1_messenger_events SET type = type + 200;")
	if err != nil {
		return err
	}
	// Update l2_messenger_events
	_, err = tx.ExecContext(ctx, "UPDATE l2_messenger_events SET type = type + 200;")
	if err != nil {
		return err
	}

	return nil
}

// Down000018 Revert gateway and messenger event_type.
func Down000018(ctx context.Context, tx *sql.Tx) error {
	// Update l1_erc20_gateway
	_, err := tx.ExecContext(ctx, "UPDATE l1_erc20_events SET type = type - 10;")
	if err != nil {
		return err
	}
	// Update l2_erc20_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc20_events SET type = type - 10;")
	if err != nil {
		return err
	}

	// Update l1_erc721_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l1_erc721_events SET type = type - 20;")
	if err != nil {
		return err
	}
	// Update l2_erc721_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc721_events SET type = type - 20;")
	if err != nil {
		return err
	}

	// Update l1_erc1155_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l1_erc1155_events SET type = type - 30;")
	if err != nil {
		return err
	}
	// Update l2_erc1155_gateway
	_, err = tx.ExecContext(ctx, "UPDATE l2_erc1155_events SET type = type - 30;")
	if err != nil {
		return err
	}

	// Update l1_messenger_events
	_, err = tx.ExecContext(ctx, "UPDATE l1_messenger_events SET type = type - 200;")
	if err != nil {
		return err
	}
	// Update l2_messenger_events
	_, err = tx.ExecContext(ctx, "UPDATE l2_messenger_events SET type = type - 200;")
	if err != nil {
		return err
	}
	return nil
}
