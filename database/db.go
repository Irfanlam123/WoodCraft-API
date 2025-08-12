// # PostgreSQL connection setuppackage database
package database

import (
	"WoodCraft-API/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	return db, nil
}
