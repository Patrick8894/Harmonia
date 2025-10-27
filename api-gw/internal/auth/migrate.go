// internal/auth/migrate.go
package auth

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func RunMigrations(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		username      VARCHAR(64)  NOT NULL UNIQUE,
		password_hash VARBINARY(255) NOT NULL,
		created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		) ENGINE=InnoDB;`)
	return err
}

func SeedDevData(ctx context.Context, db *sql.DB) error {
	// idempotent upserts (MySQL)
	seed := []struct{ u, p string }{
		{"patrick", "patrick"},
		{"admin", "admin"},
	}
	for _, s := range seed {
		hash, _ := bcrypt.GenerateFromPassword([]byte(s.p), bcrypt.DefaultCost)
		_, err := db.ExecContext(ctx, `
			INSERT INTO users (username, password_hash)
			VALUES (?, ?)
			ON DUPLICATE KEY UPDATE
			password_hash=VALUES(password_hash);`, s.u, hash)
		if err != nil {
			return err
		}
	}
	return nil
}
