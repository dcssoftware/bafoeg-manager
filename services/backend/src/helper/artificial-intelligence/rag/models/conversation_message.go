package models

type ConversationMessage struct {
	Message string `db:"message"`
	Role    string `db:"role"`
}
