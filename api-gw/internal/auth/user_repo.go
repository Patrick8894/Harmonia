package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	mysql "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct{ db *sql.DB }

func NewUserRepo(db *sql.DB) *UserRepo { return &UserRepo{db: db} }

type User struct {
	ID           uint64
	Username     string
	PasswordHash []byte
}

func (r *UserRepo) GetByUsername(ctx context.Context, u string) (*User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, username, password_hash FROM users WHERE username=?`, u)
	var x User
	if err := row.Scan(&x.ID, &x.Username, &x.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &x, nil
}

var (
	ErrUserExists = errors.New("user already exists")
)

func HashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}
func CheckPassword(hash []byte, pw string) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(pw)) == nil
}

// Create inserts a new user; returns ErrUserExists on duplicate username.
func (r *UserRepo) Create(ctx context.Context, username, password string) (*User, error) {
	username = strings.TrimSpace(username)

	hash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	res, err := r.db.ExecContext(ctx,
		`INSERT INTO users (username, password_hash) VALUES (?, ?)`,
		username, hash,
	)
	if err != nil {
		// MySQL duplicate key
		var me *mysql.MySQLError
		if errors.As(err, &me) && me.Number == 1062 {
			return nil, ErrUserExists
		}
		return nil, err
	}

	id, _ := res.LastInsertId()
	return &User{
		ID:           uint64(id),
		Username:     username,
		PasswordHash: hash,
	}, nil
}
