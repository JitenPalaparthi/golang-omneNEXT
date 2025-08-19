package database_test

import (
	"testing"
	"time"
	"users-service/database"
	"users-service/models"

	"github.com/stretchr/testify/assert"
	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewInMemoryGorm returns a *gorm.DB inside a transaction.
// The transaction is rolled back at test end, keeping schema but discarding data.
func NewInMemoryGorm(t *testing.T, autoMigrate bool) *gorm.DB {
	t.Helper()

	// Use a named in-memory DB + shared cache so the same DSN can be reused if needed.
	// Either of these work:
	//   "file::memory:?cache=shared"
	//   "file:testdb?mode=memory&cache=shared"
	dsn := "file::memory:?cache=shared"

	db, err := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{
			// faster tests; tune as you like
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		t.Fatalf("open in-memory sqlite: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("db.DB(): %v", err)
	}

	// IMPORTANT for sqlite in-memory: keep exactly one open connection,
	// otherwise the pool might open a new one that has an empty DB.
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(0)

	// Foreign keys are off by default in sqlite; enable them.
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		t.Fatalf("enable foreign_keys: %v", err)
	}

	// Migrate schema
	if autoMigrate {
		if err := db.AutoMigrate(&models.User{}, &models.Order{}); err != nil {
			t.Fatalf("automigrate: %v", err)
		}
	}

	// Run each test inside a transaction and roll back at the end.
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("begin tx: %v", tx.Error)
	}
	t.Cleanup(func() {
		_ = tx.Rollback().Error // ignore error if already rolled back
		_ = sqlDB.Close()
	})

	return tx
}

// func GetConnection() {
// 	dsn := "file::memory:?cache=shared"
// 	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	DB = db
// }

func TestCrateUser(t *testing.T) {
	db := NewInMemoryGorm(t, true)

	udb := database.NewUserDB(db)
	user := &models.User{Name: "Jiten", Email: "JitenP@outlook.com", Mobile: "9618558500", CommonModel: models.CommonModel{Status: "active", LastModified: time.Now().Unix()}}

	user, err := udb.Create(user)
	assert.Equal(t, err, nil)
	if user.ID != 1 {
		t.Fatalf("The user id should be 1 but it is %d", user.ID)
	}

}
