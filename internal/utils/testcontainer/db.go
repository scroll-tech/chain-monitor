package testcontainer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/orm/migrate"
	"github.com/scroll-tech/chain-monitor/internal/utils/database"
)

// SetupDB create postgres test container
func SetupDB(ctx context.Context, t *testing.T) *gorm.DB {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	assert.NoError(t, err)

	connStr, connErr := pgContainer.ConnectionString(ctx, "sslmode=disable")
	conf := &database.Config{
		DSN:        connStr,
		DriverName: "postgres",
	}
	assert.NoError(t, connErr)

	db, initDbErr := database.InitDB(conf)
	assert.NoError(t, initDbErr)
	assert.NotNil(t, db)

	t.Cleanup(func() {
		assert.NoError(t, database.CloseDB(db))
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	_, pingErr := database.Ping(db)
	assert.NoError(t, pingErr)

	assert.NoError(t, migrate.Migrate(db))
	return db
}
