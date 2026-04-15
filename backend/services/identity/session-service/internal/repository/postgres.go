package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.SessionRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, session *model.Session) error {
	query := `
		INSERT INTO sessions (id, user_id, role, device_id, ip, user_agent, active, expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.ExecContext(ctx, query,
		session.ID,
		session.UserID,
		session.Role,
		session.DeviceID,
		session.IP,
		session.UserAgent,
		session.Active,
		session.ExpiresAt,
		session.CreatedAt,
		session.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Session, error) {
	query := `
		SELECT id, user_id, role, device_id, ip, user_agent, active, expires_at, created_at, updated_at
		FROM sessions WHERE id = $1`

	s := &model.Session{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&s.ID, &s.UserID, &s.Role, &s.DeviceID,
		&s.IP, &s.UserAgent, &s.Active,
		&s.ExpiresAt, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *postgresRepository) FindByUserID(ctx context.Context, userID string) ([]*model.Session, error) {
	query := `
		SELECT id, user_id, role, device_id, ip, user_agent, active, expires_at, created_at, updated_at
		FROM sessions WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*model.Session
	for rows.Next() {
		s := &model.Session{}
		if err := rows.Scan(
			&s.ID, &s.UserID, &s.Role, &s.DeviceID,
			&s.IP, &s.UserAgent, &s.Active,
			&s.ExpiresAt, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	return sessions, rows.Err()
}

func (r *postgresRepository) Invalidate(ctx context.Context, id string) error {
	query := `UPDATE sessions SET active = false, updated_at = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), id)
	return err
}

func (r *postgresRepository) InvalidateAllByUser(ctx context.Context, userID string) error {
	query := `UPDATE sessions SET active = false, updated_at = $1 WHERE user_id = $2 AND active = true`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), userID)
	return err
}
