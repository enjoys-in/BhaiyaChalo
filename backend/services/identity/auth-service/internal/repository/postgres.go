package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.AuthRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) FindUserByPhone(ctx context.Context, phone string, role model.Role) (*model.User, error) {
	query := `SELECT id, phone, email, role, status, created_at, updated_at FROM users WHERE phone = $1 AND role = $2`
	row := r.db.QueryRowContext(ctx, query, phone, role)

	var u model.User
	err := row.Scan(&u.ID, &u.Phone, &u.Email, &u.Role, &u.Status, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, phone, email, role, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Phone, user.Email, user.Role, user.Status, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *postgresRepo) StoreRefreshToken(ctx context.Context, token *model.Token) error {
	query := `INSERT INTO refresh_tokens (id, user_id, refresh_token, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		token.ID, token.UserID, token.RefreshToken, token.ExpiresAt, token.CreatedAt)
	return err
}

func (r *postgresRepo) FindRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error) {
	query := `SELECT id, user_id, refresh_token, expires_at, created_at FROM refresh_tokens
		WHERE refresh_token = $1 AND expires_at > $2`
	row := r.db.QueryRowContext(ctx, query, refreshToken, time.Now())

	var t model.Token
	err := row.Scan(&t.ID, &t.UserID, &t.RefreshToken, &t.ExpiresAt, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *postgresRepo) RevokeRefreshToken(ctx context.Context, userID string) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
