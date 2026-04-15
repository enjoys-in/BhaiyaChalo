package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/ports"
)

type userService struct {
	repo      ports.UserRepository
	publisher ports.EventPublisher
}

func NewUserService(repo ports.UserRepository, publisher ports.EventPublisher) ports.UserService {
	return &userService{repo: repo, publisher: publisher}
}

func (s *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	now := time.Now().UTC()
	user := &model.User{
		ID:          uuid.New(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
		Email:       req.Email,
		AvatarURL:   req.AvatarURL,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		CityID:      req.CityID,
		Status:      model.UserStatusActive,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	_ = s.publisher.PublishUserCreated(ctx, user)
	return toUserResponse(user), nil
}

func (s *userService) GetUser(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func (s *userService) GetUserByPhone(ctx context.Context, phone string) (*dto.UserResponse, error) {
	user, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	applyUserUpdates(user, req)
	user.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	_ = s.publisher.PublishUserUpdated(ctx, user)
	return toUserResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	_ = s.publisher.PublishUserDeleted(ctx, user)
	return nil
}

func (s *userService) GetUserProfile(ctx context.Context, id uuid.UUID) (*dto.UserProfileResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	addrs, err := s.repo.ListAddresses(ctx, id)
	if err != nil {
		return nil, err
	}
	contacts, err := s.repo.ListEmergencyContacts(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.UserProfileResponse{
		User:              *toUserResponse(user),
		Addresses:         toAddressResponses(addrs),
		EmergencyContacts: toEmergencyContactResponses(contacts),
	}, nil
}

func (s *userService) AddAddress(ctx context.Context, userID uuid.UUID, req dto.AddAddressRequest) (*dto.AddressResponse, error) {
	addr := &model.Address{
		ID:               uuid.New(),
		UserID:           userID,
		Label:            model.AddressLabel(req.Label),
		Lat:              req.Lat,
		Lng:              req.Lng,
		FormattedAddress: req.FormattedAddress,
		IsDefault:        req.IsDefault,
		CreatedAt:        time.Now().UTC(),
	}
	if err := s.repo.AddAddress(ctx, addr); err != nil {
		return nil, err
	}
	return toAddressResponse(addr), nil
}

func (s *userService) UpdateAddress(ctx context.Context, addressID uuid.UUID, req dto.UpdateAddressRequest) (*dto.AddressResponse, error) {
	// Fetch existing address list to find the target — keeps repo interface minimal
	addr := &model.Address{ID: addressID}
	if req.Label != nil {
		addr.Label = model.AddressLabel(*req.Label)
	}
	if req.Lat != nil {
		addr.Lat = *req.Lat
	}
	if req.Lng != nil {
		addr.Lng = *req.Lng
	}
	if req.FormattedAddress != nil {
		addr.FormattedAddress = *req.FormattedAddress
	}
	if req.IsDefault != nil {
		addr.IsDefault = *req.IsDefault
	}
	if err := s.repo.UpdateAddress(ctx, addr); err != nil {
		return nil, err
	}
	return toAddressResponse(addr), nil
}

func (s *userService) DeleteAddress(ctx context.Context, addressID uuid.UUID) error {
	return s.repo.DeleteAddress(ctx, addressID)
}

func (s *userService) ListAddresses(ctx context.Context, userID uuid.UUID) ([]dto.AddressResponse, error) {
	addrs, err := s.repo.ListAddresses(ctx, userID)
	if err != nil {
		return nil, err
	}
	return toAddressResponses(addrs), nil
}

func (s *userService) AddEmergencyContact(ctx context.Context, userID uuid.UUID, req dto.AddEmergencyContactRequest) (*dto.EmergencyContactResponse, error) {
	contact := &model.EmergencyContact{
		ID:       uuid.New(),
		UserID:   userID,
		Name:     req.Name,
		Phone:    req.Phone,
		Relation: req.Relation,
	}
	if err := s.repo.AddEmergencyContact(ctx, contact); err != nil {
		return nil, err
	}
	return toEmergencyContactResponse(contact), nil
}

func (s *userService) ListEmergencyContacts(ctx context.Context, userID uuid.UUID) ([]dto.EmergencyContactResponse, error) {
	contacts, err := s.repo.ListEmergencyContacts(ctx, userID)
	if err != nil {
		return nil, err
	}
	return toEmergencyContactResponses(contacts), nil
}

// --- mappers ---

func applyUserUpdates(user *model.User, req dto.UpdateUserRequest) {
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.AvatarURL != nil {
		user.AvatarURL = *req.AvatarURL
	}
	if req.Gender != nil {
		user.Gender = *req.Gender
	}
	if req.DateOfBirth != nil {
		user.DateOfBirth = req.DateOfBirth
	}
	if req.CityID != nil {
		user.CityID = *req.CityID
	}
}

func toUserResponse(u *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Phone:       u.Phone,
		Email:       u.Email,
		AvatarURL:   u.AvatarURL,
		Gender:      u.Gender,
		DateOfBirth: u.DateOfBirth,
		CityID:      u.CityID,
		Status:      string(u.Status),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func toAddressResponse(a *model.Address) *dto.AddressResponse {
	return &dto.AddressResponse{
		ID:               a.ID,
		UserID:           a.UserID,
		Label:            string(a.Label),
		Lat:              a.Lat,
		Lng:              a.Lng,
		FormattedAddress: a.FormattedAddress,
		IsDefault:        a.IsDefault,
		CreatedAt:        a.CreatedAt,
	}
}

func toAddressResponses(addrs []model.Address) []dto.AddressResponse {
	out := make([]dto.AddressResponse, len(addrs))
	for i := range addrs {
		out[i] = *toAddressResponse(&addrs[i])
	}
	return out
}

func toEmergencyContactResponse(c *model.EmergencyContact) *dto.EmergencyContactResponse {
	return &dto.EmergencyContactResponse{
		ID:       c.ID,
		UserID:   c.UserID,
		Name:     c.Name,
		Phone:    c.Phone,
		Relation: c.Relation,
	}
}

func toEmergencyContactResponses(contacts []model.EmergencyContact) []dto.EmergencyContactResponse {
	out := make([]dto.EmergencyContactResponse, len(contacts))
	for i := range contacts {
		out[i] = *toEmergencyContactResponse(&contacts[i])
	}
	return out
}
