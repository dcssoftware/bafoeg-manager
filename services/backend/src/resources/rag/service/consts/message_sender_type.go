package consts

type MessageSenderType string

const (
	MessageSenderTypeUser   MessageSenderType = "USER"
	MessageSenderTypeSystem MessageSenderType = "SYSTEM"
)

func (m MessageSenderType) String() string {
	return string(m)
}
