package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/ports"
)

type mongoRepo struct {
	trips    *mongo.Collection
	timeline *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) ports.TripRepository {
	return &mongoRepo{
		trips:    db.Collection("trips"),
		timeline: db.Collection("trip_timeline"),
	}
}

func (r *mongoRepo) Create(ctx context.Context, trip *model.Trip) error {
	_, err := r.trips.InsertOne(ctx, trip)
	return err
}

func (r *mongoRepo) FindByID(ctx context.Context, id string) (*model.Trip, error) {
	var trip model.Trip
	err := r.trips.FindOne(ctx, bson.M{"_id": id}).Decode(&trip)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("trip not found")
		}
		return nil, err
	}
	return &trip, nil
}

func (r *mongoRepo) FindByUserID(ctx context.Context, userID string, limit, offset int) ([]model.Trip, error) {
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.trips.Find(ctx, bson.M{"user_id": userID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trips []model.Trip
	if err := cursor.All(ctx, &trips); err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *mongoRepo) FindByDriverID(ctx context.Context, driverID string, limit, offset int) ([]model.Trip, error) {
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.trips.Find(ctx, bson.M{"driver_id": driverID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trips []model.Trip
	if err := cursor.All(ctx, &trips); err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *mongoRepo) UpdateStatus(ctx context.Context, tripID string, status model.TripStatus) error {
	now := time.Now().UTC()
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": now,
		},
	}

	switch status {
	case model.TripStatusStarted:
		update["$set"].(bson.M)["actual_pickup_at"] = now
	case model.TripStatusCompleted:
		update["$set"].(bson.M)["actual_drop_at"] = now
	}

	result, err := r.trips.UpdateOne(ctx, bson.M{"_id": tripID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("trip not found")
	}
	return nil
}

func (r *mongoRepo) AddTimelineEvent(ctx context.Context, event *model.TripTimeline) error {
	_, err := r.timeline.InsertOne(ctx, event)
	return err
}

func (r *mongoRepo) GetTimeline(ctx context.Context, tripID string) ([]model.TripTimeline, error) {
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: 1}})
	cursor, err := r.timeline.Find(ctx, bson.M{"trip_id": tripID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []model.TripTimeline
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}
