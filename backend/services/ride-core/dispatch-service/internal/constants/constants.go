package constants

const (
	TopicOfferSent      = "dispatch.offer.sent"
	TopicOfferAccepted  = "dispatch.offer.accepted"
	TopicOfferRejected  = "dispatch.offer.rejected"
	TopicOfferExpired   = "dispatch.offer.expired"
	TopicDispatchFailed = "dispatch.failed"

	HeaderBookingID = "X-Booking-ID"
	HeaderDriverID  = "X-Driver-ID"

	DefaultOfferTimeoutSeconds = 30
	DefaultMaxRetries          = 3
	MaxCandidatesPerRound      = 5
)
