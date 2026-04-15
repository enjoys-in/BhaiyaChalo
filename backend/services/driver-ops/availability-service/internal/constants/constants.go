package constants

import "time"

const (
	ServiceName = "availability-service"

	EventDriverOnline  = "driver.online"
	EventDriverOffline = "driver.offline"
	EventDriverBusy    = "driver.busy"
	EventDriverFree    = "driver.free"

	TopicDriverAvailability = "driver-availability-events"

	AvailabilityTTL = 10 * time.Minute
)
