package constants

const ServiceName = "matching-service"

// Event types
const (
	EventDriverMatched = "matching.driver_matched"
	EventMatchFailed   = "matching.match_failed"
)

// Match request statuses
const (
	MatchStatusSearching = "searching"
	MatchStatusMatched   = "matched"
	MatchStatusFailed    = "failed"
)

// Kafka topics
const (
	TopicMatchingEvents = "matching-events"
)
