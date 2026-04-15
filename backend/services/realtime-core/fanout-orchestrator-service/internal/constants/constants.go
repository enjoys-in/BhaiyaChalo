package constants

const (
	ServiceName = "fanout-orchestrator-service"

	TopicMessageDelivered = "fanout.message.delivered"
	TopicMessageFailed    = "fanout.message.failed"
	TopicFanoutOutbound   = "fanout.outbound"

	RedisMessagePrefix  = "fanout_msg:"
	RedisDeliveryPrefix = "fanout_delivery:"
)
