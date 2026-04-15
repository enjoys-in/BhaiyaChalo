package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.DispatchRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) CreateOffer(ctx context.Context, offer *model.DispatchOffer) error {
	query := `
		INSERT INTO dispatch_offers (id, booking_id, driver_id, city_id, status, offer_expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query,
		offer.ID, offer.BookingID, offer.DriverID, offer.CityID,
		offer.Status, offer.OfferExpiresAt, offer.CreatedAt,
	)
	return err
}

func (r *postgresRepo) FindOfferByID(ctx context.Context, offerID string) (*model.DispatchOffer, error) {
	query := `
		SELECT id, booking_id, driver_id, city_id, status, offer_expires_at, responded_at, created_at
		FROM dispatch_offers WHERE id = $1`
	offer := &model.DispatchOffer{}
	err := r.db.QueryRowContext(ctx, query, offerID).Scan(
		&offer.ID, &offer.BookingID, &offer.DriverID, &offer.CityID,
		&offer.Status, &offer.OfferExpiresAt, &offer.RespondedAt, &offer.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("offer not found: %s", offerID)
	}
	return offer, err
}

func (r *postgresRepo) UpdateOfferStatus(ctx context.Context, offerID string, status model.OfferStatus) error {
	query := `UPDATE dispatch_offers SET status = $1, responded_at = NOW() WHERE id = $2`
	res, err := r.db.ExecContext(ctx, query, status, offerID)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("offer not found: %s", offerID)
	}
	return nil
}

func (r *postgresRepo) FindPendingByBooking(ctx context.Context, bookingID string) ([]*model.DispatchOffer, error) {
	query := `
		SELECT id, booking_id, driver_id, city_id, status, offer_expires_at, responded_at, created_at
		FROM dispatch_offers WHERE booking_id = $1 AND status = $2
		ORDER BY created_at ASC`
	rows, err := r.db.QueryContext(ctx, query, bookingID, model.OfferStatusPending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*model.DispatchOffer
	for rows.Next() {
		o := &model.DispatchOffer{}
		if err := rows.Scan(
			&o.ID, &o.BookingID, &o.DriverID, &o.CityID,
			&o.Status, &o.OfferExpiresAt, &o.RespondedAt, &o.CreatedAt,
		); err != nil {
			return nil, err
		}
		offers = append(offers, o)
	}
	return offers, rows.Err()
}

func (r *postgresRepo) CreateRound(ctx context.Context, round *model.DispatchRound) error {
	driverIDsJSON, err := json.Marshal(round.CandidateDriverIDs)
	if err != nil {
		return fmt.Errorf("marshal driver IDs: %w", err)
	}
	query := `
		INSERT INTO dispatch_rounds (id, booking_id, round_number, candidate_driver_ids, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.db.ExecContext(ctx, query,
		round.ID, round.BookingID, round.RoundNumber,
		driverIDsJSON, round.Status, round.CreatedAt,
	)
	return err
}

func (r *postgresRepo) FindRoundsByBooking(ctx context.Context, bookingID string) ([]*model.DispatchRound, error) {
	query := `
		SELECT id, booking_id, round_number, candidate_driver_ids, status, created_at
		FROM dispatch_rounds WHERE booking_id = $1
		ORDER BY round_number ASC`
	rows, err := r.db.QueryContext(ctx, query, bookingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rounds []*model.DispatchRound
	for rows.Next() {
		rd := &model.DispatchRound{}
		var driverIDsJSON []byte
		if err := rows.Scan(
			&rd.ID, &rd.BookingID, &rd.RoundNumber,
			&driverIDsJSON, &rd.Status, &rd.CreatedAt,
		); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(driverIDsJSON, &rd.CandidateDriverIDs); err != nil {
			return nil, fmt.Errorf("unmarshal driver IDs: %w", err)
		}
		rounds = append(rounds, rd)
	}
	return rounds, rows.Err()
}
