package constants

// Kafka topic names — all prefixed by domain.
const (
	TopicAuthEvents         = "identity.auth.events"
	TopicSessionEvents      = "identity.session.events"
	TopicUserEvents         = "profile.user.events"
	TopicDriverEvents       = "profile.driver.events"
	TopicBookingEvents      = "ride.booking.events"
	TopicMatchingEvents     = "ride.matching.events"
	TopicDispatchEvents     = "ride.dispatch.events"
	TopicTrackingEvents     = "ride.tracking.events"
	TopicTripEvents         = "ride.trip.events"
	TopicFareEvents         = "ride.fare.events"
	TopicPaymentEvents      = "payments.payment.events"
	TopicWalletEvents       = "payments.wallet.events"
	TopicNotificationEvents = "engagement.notification.events"
	TopicRatingEvents       = "engagement.rating.events"
	TopicPromoEvents        = "growth.promo.events"
	TopicFraudEvents        = "trust.fraud.events"
	TopicSupportEvents      = "support.ticket.events"
	TopicAnalyticsEvents    = "analytics.ingestion.events"
	TopicAuditEvents        = "analytics.audit.events"
)

// Header keys.
const (
	HeaderRequestID   = "X-Request-ID"
	HeaderCityID      = "X-City-ID"
	HeaderUserID      = "X-User-ID"
	HeaderDriverID    = "X-Driver-ID"
	HeaderAuthToken   = "Authorization"
	HeaderContentType = "Content-Type"
)
