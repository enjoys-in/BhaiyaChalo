package constants

const (
	ServiceName = "socket-session-service"

	TopicSessionConnected    = "session.connected"
	TopicSessionDisconnected = "session.disconnected"

	RedisSessionPrefix  = "session:"
	RedisUserSessions   = "user_sessions:"
	RedisServerSessions = "server_sessions:"
	RedisActiveCount    = "session:active_count"
)
