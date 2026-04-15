package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/ports"
)

var ErrNotFound = errors.New("record not found")

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.UserRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, phone, email, avatar_url, gender, date_of_birth, city_id, status, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.FirstName, user.LastName, user.Phone, user.Email,
		user.AvatarURL, user.Gender, user.DateOfBirth, user.CityID,
		user.Status, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

func (r *postgresRepo) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, avatar_url, gender, date_of_birth, city_id, status, created_at, updated_at, deleted_at
		FROM users WHERE id = $1 AND deleted_at IS NULL`
	return r.scanUser(r.db.QueryRowContext(ctx, query, id))
}

func (r *postgresRepo) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, avatar_url, gender, date_of_birth, city_id, status, created_at, updated_at, deleted_at
		FROM users WHERE phone = $1 AND deleted_at IS NULL`
	return r.scanUser(r.db.QueryRowContext(ctx, query, phone))
}

func (r *postgresRepo) Update(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users SET first_name=$1, last_name=$2, email=$3, avatar_url=$4, gender=$5, date_of_birth=$6, city_id=$7, status=$8, updated_at=$9
		WHERE id = $10 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query,
		user.FirstName, user.LastName, user.Email, user.AvatarURL,
		user.Gender, user.DateOfBirth, user.CityID, user.Status,
		user.UpdatedAt, user.ID,
	)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(res)
}

func (r *postgresRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET deleted_at = $1, updated_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	now := time.Now().UTC()
	res, err := r.db.ExecContext(ctx, query, now, id)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(res)
}

func (r *postgresRepo) AddAddress(ctx context.Context, addr *model.Address) error {
	query := `
		INSERT INTO addresses (id, user_id, label, lat, lng, formatted_address, is_default, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.ExecContext(ctx, query,
		addr.ID, addr.UserID, addr.Label, addr.Lat, addr.Lng,
		addr.FormattedAddress, addr.IsDefault, addr.CreatedAt,
	)
	return err
}

func (r *postgresRepo) UpdateAddress(ctx context.Context, addr *model.Address) error {
	query := `
		UPDATE addresses SET label=$1, lat=$2, lng=$3, formatted_address=$4, is_default=$5
		WHERE id = $6`
	res, err := r.db.ExecContext(ctx, query,
		addr.Label, addr.Lat, addr.Lng, addr.FormattedAddress, addr.IsDefault, addr.ID,
	)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(res)
}

func (r *postgresRepo) DeleteAddress(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM addresses WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(res)
}

func (r *postgresRepo) ListAddresses(ctx context.Context, userID uuid.UUID) ([]model.Address, error) {
	query := `SELECT id, user_id, label, lat, lng, formatted_address, is_default, created_at FROM addresses WHERE user_id = $1 ORDER BY created_at`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addrs []model.Address
	for rows.Next() {
		var a model.Address
		if err := rows.Scan(&a.ID, &a.UserID, &a.Label, &a.Lat, &a.Lng, &a.FormattedAddress, &a.IsDefault, &a.CreatedAt); err != nil {
			return nil, err
		}
		addrs = append(addrs, a)
	}
	return addrs, rows.Err()
}

func (r *postgresRepo) AddEmergencyContact(ctx context.Context, c *model.EmergencyContact) error {
	query := `INSERT INTO emergency_contacts (id, user_id, name, phone, relation) VALUES ($1,$2,$3,$4,$5)`
	_, err := r.db.ExecContext(ctx, query, c.ID, c.UserID, c.Name, c.Phone, c.Relation)
	return err
}

func (r *postgresRepo) ListEmergencyContacts(ctx context.Context, userID uuid.UUID) ([]model.EmergencyContact, error) {
	query := `SELECT id, user_id, name, phone, relation FROM emergency_contacts WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.EmergencyContact
	for rows.Next() {
		var c model.EmergencyContact
		if err := rows.Scan(&c.ID, &c.UserID, &c.Name, &c.Phone, &c.Relation); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, rows.Err()
}

func (r *postgresRepo) scanUser(row *sql.Row) (*model.User, error) {
	var u model.User
	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Phone, &u.Email,
		&u.AvatarURL, &u.Gender, &u.DateOfBirth, &u.CityID,
		&u.Status, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &u, err
}

func (r *postgresRepo) checkRowsAffected(res sql.Result) error {
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
