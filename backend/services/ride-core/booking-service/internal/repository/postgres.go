package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.BookingRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, booking *model.Booking) error {
	query := `
		INSERT INTO bookings (id, user_id, city_id, pickup_lat, pickup_lng, pickup_address,
			drop_lat, drop_lng, drop_address, vehicle_type, estimated_fare, final_fare,
			promo_code, discount_amount, status, driver_id, payment_method, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)`

	_, err := r.db.ExecContext(ctx, query,
		booking.ID, booking.UserID, booking.CityID,
		booking.PickupLat, booking.PickupLng, booking.PickupAddress,
		booking.DropLat, booking.DropLng, booking.DropAddress,
		booking.VehicleType, booking.EstimatedFare, booking.FinalFare,
		booking.PromoCode, booking.DiscountAmount, booking.Status,
		booking.DriverID, booking.PaymentMethod,
		booking.CreatedAt, booking.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Booking, error) {
	query := `
		SELECT id, user_id, city_id, pickup_lat, pickup_lng, pickup_address,
			drop_lat, drop_lng, drop_address, vehicle_type, estimated_fare, final_fare,
			promo_code, discount_amount, status, driver_id, payment_method, created_at, updated_at
		FROM bookings WHERE id = $1`

	b := &model.Booking{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&b.ID, &b.UserID, &b.CityID,
		&b.PickupLat, &b.PickupLng, &b.PickupAddress,
		&b.DropLat, &b.DropLng, &b.DropAddress,
		&b.VehicleType, &b.EstimatedFare, &b.FinalFare,
		&b.PromoCode, &b.DiscountAmount, &b.Status,
		&b.DriverID, &b.PaymentMethod,
		&b.CreatedAt, &b.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *postgresRepository) FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Booking, error) {
	query := `
		SELECT id, user_id, city_id, pickup_lat, pickup_lng, pickup_address,
			drop_lat, drop_lng, drop_address, vehicle_type, estimated_fare, final_fare,
			promo_code, discount_amount, status, driver_id, payment_method, created_at, updated_at
		FROM bookings WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*model.Booking
	for rows.Next() {
		b := &model.Booking{}
		if err := rows.Scan(
			&b.ID, &b.UserID, &b.CityID,
			&b.PickupLat, &b.PickupLng, &b.PickupAddress,
			&b.DropLat, &b.DropLng, &b.DropAddress,
			&b.VehicleType, &b.EstimatedFare, &b.FinalFare,
			&b.PromoCode, &b.DiscountAmount, &b.Status,
			&b.DriverID, &b.PaymentMethod,
			&b.CreatedAt, &b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, rows.Err()
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id string, status model.BookingStatus) error {
	query := `UPDATE bookings SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, status, time.Now().UTC(), id)
	return err
}

func (r *postgresRepository) UpdateDriver(ctx context.Context, id string, driverID string) error {
	query := `UPDATE bookings SET driver_id = $1, status = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, driverID, model.StatusDriverAssigned, time.Now().UTC(), id)
	return err
}

func (r *postgresRepository) UpdateFare(ctx context.Context, id string, finalFare float64) error {
	query := `UPDATE bookings SET final_fare = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, finalFare, time.Now().UTC(), id)
	return err
}

func (r *postgresRepository) Cancel(ctx context.Context, id string) error {
	query := `UPDATE bookings SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, model.StatusCancelled, time.Now().UTC(), id)
	return err
}
