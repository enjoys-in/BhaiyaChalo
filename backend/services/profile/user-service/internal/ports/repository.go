package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	FindByPhone(ctx context.Context, phone string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error

	AddAddress(ctx context.Context, address *model.Address) error
	UpdateAddress(ctx context.Context, address *model.Address) error
	DeleteAddress(ctx context.Context, id uuid.UUID) error
	ListAddresses(ctx context.Context, userID uuid.UUID) ([]model.Address, error)

	AddEmergencyContact(ctx context.Context, contact *model.EmergencyContact) error
	ListEmergencyContacts(ctx context.Context, userID uuid.UUID) ([]model.EmergencyContact, error)
}
