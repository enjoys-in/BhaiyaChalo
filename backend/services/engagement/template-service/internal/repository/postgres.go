package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/ports"
)

type templateRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.TemplateRepository {
	return &templateRepository{db: db}
}

func (r *templateRepository) Create(ctx context.Context, entity *model.Template) error {
	// TODO: implement
	return nil
}

func (r *templateRepository) FindByID(ctx context.Context, id string) (*model.Template, error) {
	// TODO: implement
	return nil, nil
}

func (r *templateRepository) Update(ctx context.Context, entity *model.Template) error {
	// TODO: implement
	return nil
}

func (r *templateRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
