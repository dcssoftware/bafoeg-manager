package sessionlocals

type SessionLocals string

const (
	UserUUID     SessionLocals = "user-uuid"
	Permissions  SessionLocals = "permissions"
	JWTtoken     SessionLocals = "JWTtoken"
	SessionUUID  SessionLocals = "session-uuid"
	AuthProvider SessionLocals = "auth-provider"
)
