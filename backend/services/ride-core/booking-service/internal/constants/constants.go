package constants

const ServiceName = "booking-service"

// Event types
const (
	EventBookingCreated   = "booking.created"
	EventBookingConfirmed = "booking.confirmed"
	EventBookingCancelled = "booking.cancelled"
	EventDriverAssigned   = "booking.driver_assigned"
)

// Kafka topics
const (
	TopicBookingEvents = "booking-events"
)
