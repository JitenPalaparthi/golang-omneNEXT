package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dsnFromEnv() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASSWORD", "postgres")
	db := getenv("DB_NAME", "usersdb")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, db)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" { return v }
	return def
}

func Connect() (*gorm.DB, error) {
	dsn := dsnFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { return nil, err }
	sqlDB, err := db.DB()
	if err != nil { return nil, err }
	setupPool(sqlDB)
	return db, nil
}

func setupPool(sqlDB *sql.DB) {
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
}
