package constants

const (
	ServiceName = "driver-service"

	StatusPending   = "pending"
	StatusApproved  = "approved"
	StatusActive    = "active"
	StatusInactive  = "inactive"
	StatusSuspended = "suspended"

	EventDriverCreated       = "driver.created"
	EventDriverUpdated       = "driver.updated"
	EventDriverStatusChanged = "driver.status_changed"

	TopicDriverEvents = "driver-events"
)
