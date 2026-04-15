package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.DocumentRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, doc *model.Document) error {
	query := `
		INSERT INTO documents (id, owner_id, owner_type, doc_type, file_url, status,
			reviewer_id, review_note, expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.ExecContext(ctx, query,
		doc.ID, doc.OwnerID, doc.OwnerType, doc.DocType,
		doc.FileURL, doc.Status, doc.ReviewerID, doc.ReviewNote,
		doc.ExpiresAt, doc.CreatedAt, doc.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Document, error) {
	query := `
		SELECT id, owner_id, owner_type, doc_type, file_url, status,
			reviewer_id, review_note, expires_at, created_at, updated_at
		FROM documents WHERE id = $1`

	d := &model.Document{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&d.ID, &d.OwnerID, &d.OwnerType, &d.DocType,
		&d.FileURL, &d.Status, &d.ReviewerID, &d.ReviewNote,
		&d.ExpiresAt, &d.CreatedAt, &d.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *postgresRepository) FindByOwner(ctx context.Context, ownerID string, ownerType model.OwnerType) ([]*model.Document, error) {
	query := `
		SELECT id, owner_id, owner_type, doc_type, file_url, status,
			reviewer_id, review_note, expires_at, created_at, updated_at
		FROM documents WHERE owner_id = $1 AND owner_type = $2
		ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, ownerID, ownerType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDocuments(rows)
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id string, status model.DocumentStatus, reviewerID, reviewNote string) error {
	query := `
		UPDATE documents SET status = $1, reviewer_id = $2, review_note = $3, updated_at = $4
		WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query, status, reviewerID, reviewNote, time.Now().UTC(), id)
	return err
}

func (r *postgresRepository) ListPending(ctx context.Context, limit, offset int) ([]*model.Document, error) {
	query := `
		SELECT id, owner_id, owner_type, doc_type, file_url, status,
			reviewer_id, review_note, expires_at, created_at, updated_at
		FROM documents WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, model.DocumentStatusPending, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDocuments(rows)
}

func (r *postgresRepository) ListExpiring(ctx context.Context, before time.Time) ([]*model.Document, error) {
	query := `
		SELECT id, owner_id, owner_type, doc_type, file_url, status,
			reviewer_id, review_note, expires_at, created_at, updated_at
		FROM documents WHERE status = $1 AND expires_at IS NOT NULL AND expires_at <= $2
		ORDER BY expires_at ASC`

	rows, err := r.db.QueryContext(ctx, query, model.DocumentStatusApproved, before)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDocuments(rows)
}

func scanDocuments(rows *sql.Rows) ([]*model.Document, error) {
	var docs []*model.Document
	for rows.Next() {
		d := &model.Document{}
		if err := rows.Scan(
			&d.ID, &d.OwnerID, &d.OwnerType, &d.DocType,
			&d.FileURL, &d.Status, &d.ReviewerID, &d.ReviewNote,
			&d.ExpiresAt, &d.CreatedAt, &d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		docs = append(docs, d)
	}
	return docs, rows.Err()
}
