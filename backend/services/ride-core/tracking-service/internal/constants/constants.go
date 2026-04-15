package constants

const (
	TopicLocationUpdated = "tracking.location.updated"
	TopicTrackingStarted = "tracking.session.started"
	TopicTrackingStopped = "tracking.session.stopped"

	RedisGeoKey         = "driver:locations"
	RedisLocationPrefix = "driver:loc:"

	DefaultLocationTTL      = 300
	DefaultSnapshotInterval = 5
)
