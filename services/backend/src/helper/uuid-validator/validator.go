package uuidvalidator

import (
	"github.com/google/uuid"
)

func ValidateUUID(uuidStr string) bool {
	uuidErr := uuid.Validate(uuidStr)
	return uuidErr == nil
}
