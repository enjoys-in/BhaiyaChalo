package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUser(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error)
	GetUserByPhone(ctx context.Context, phone string) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUserProfile(ctx context.Context, id uuid.UUID) (*dto.UserProfileResponse, error)

	AddAddress(ctx context.Context, userID uuid.UUID, req dto.AddAddressRequest) (*dto.AddressResponse, error)
	UpdateAddress(ctx context.Context, addressID uuid.UUID, req dto.UpdateAddressRequest) (*dto.AddressResponse, error)
	DeleteAddress(ctx context.Context, addressID uuid.UUID) error
	ListAddresses(ctx context.Context, userID uuid.UUID) ([]dto.AddressResponse, error)

	AddEmergencyContact(ctx context.Context, userID uuid.UUID, req dto.AddEmergencyContactRequest) (*dto.EmergencyContactResponse, error)
	ListEmergencyContacts(ctx context.Context, userID uuid.UUID) ([]dto.EmergencyContactResponse, error)
}
